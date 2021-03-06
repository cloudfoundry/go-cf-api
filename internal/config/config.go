package config

import (
	_ "embed"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/cloudfoundry/go-cf-api/internal/logging/tags"
	"github.com/go-playground/validator"
	"github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v4"
	"github.com/mitchellh/mapstructure"
	promconfig "github.com/prometheus/common/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

//go:embed defaults.yml
var defaults []byte

type CfAPIConfig struct {
	Listen           string        `yaml:"listen"`
	ExternalDomain   string        `yaml:"external_domain" validate:"required"`
	ExternalProtocol string        `yaml:"external_protocol" validate:"required"`
	Info             InfoConfig    `yaml:"info"`
	DB               DBConfig      `yaml:"db"`
	Log              ZapConfig     `yaml:"log"`
	RateLimit        RateLimitConf `yaml:"rate_limiting"`
	Uaa              UaaConfig     `yaml:"uaa"`
	URLs             URLs          `yaml:"urls"`
	AppSSH           AppSSH        `yaml:"app_ssh"`
}

type InfoConfig struct {
	Name                     string `yaml:"name" validate:"required"`
	Build                    string `yaml:"build" validate:"required"`
	Version                  int    `yaml:"version" validate:"gte=0"`
	SupportAddress           string `yaml:"support_address" validate:"required"`
	Description              string `yaml:"description" validate:"required"`
	MinCliVersion            string `yaml:"min_cli_version"`
	MinRecommendedCliVersion string `yaml:"min_recommended_cli_version"`
}

type URLs struct {
	LogCache   string  `yaml:"log_cache" validate:"required,url"`
	LogStream  string  `yaml:"log_stream" validate:"required,url"`
	Doppler    string  `yaml:"doppler" validate:"required,url"`
	Login      string  `yaml:"login" validate:"required,url"`
	UAA        string  `yaml:"uaa" validate:"required,url"`
	RoutingAPI *string `yaml:"routing_api" validate:"omitempty,url"`
}

type AppSSH struct {
	Endpoint           string `yaml:"endpoint" validate:"required,url"`
	OAuthClient        string `yaml:"oauth_client" validate:"required"`
	HostKeyFingerprint string `yaml:"host_key_fingerprint"`
}

type UaaConfig struct {
	URL    string                      `yaml:"url"`
	Client promconfig.HTTPClientConfig `yaml:"client"`
}

type DBConfig struct {
	// For connection string documentation see
	// Postgres: https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-PARAMKEYWORDS
	// Mysql: https://github.com/go-sql-driver/mysql#dsn-data-source-name
	ConnectionString string    `yaml:"connection_string" validate:"required"`
	Type             string    `yaml:"type" validate:"required,oneof=postgres mysql"`
	Log              ZapConfig `yaml:"log"`
	MaxConnections   int       `yaml:"max_connections" validate:"gte=1,lte=1000"`
	MinConnections   int       `yaml:"min_connections" validate:"gte=1,lte=1000,ltefield=MaxConnections"`
}

type ZapConfig struct {
	Level      string `yaml:"level" validate:"required,oneof=debug info warn error dpanic panic fatal"`
	Production bool   `yaml:"production"`
}

type RateLimitConf struct {
	Enabled                        bool          `yaml:"enabled"`
	PerProcessGeneralLimit         int           `yaml:"per_process_general_limit"`
	GlobalGeneralLimit             int           `yaml:"global_general_limit"`
	PerProcessUnauthenticatedLimit int           `yaml:"per_process_unauthenticated_limit"`
	GlobalUnauthenticatedLimit     int           `yaml:"global_unauthenticated_limit"`
	ResetInterval                  time.Duration `yaml:"reset_interval"`
}

// Parses the config once, otherwise just returns it
// Priority from lowest to highest: Default Values, Config File, Environment Variables
// After setting the Config in this order it is validated according to struct tags
// If validation is not succefull we have a invalid config and thus throw a FATAL
func Get(configFile string) (*CfAPIConfig, error) {
	var config *CfAPIConfig
	err := yaml.Unmarshal(defaults, &config)
	if err != nil {
		return nil, err
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("CC")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// If Config File is present, tell viper about it
	if configFile != "" {
		viper.SetConfigFile(configFile)
	}

	var configFileNotFound viper.ConfigFileNotFoundError
	if err := viper.ReadInConfig(); err != nil && !errors.As(err, &configFileNotFound) {
		return nil, err
	}

	zap.L().Info(fmt.Sprintf("Using config file %s", viper.ConfigFileUsed()), zap.String(tags.File, viper.ConfigFileUsed()))

	err = viper.Unmarshal(config, func(config *mapstructure.DecoderConfig) {
		// The Prometheus HTTPClientConfig class we use has "yaml" tags not "mapstructure"
		config.TagName = "yaml"
	})
	if err != nil {
		return nil, err
	}

	if zap.L().Core().Enabled(zap.DebugLevel) {
		usedConfig, err := yaml.Marshal(config)
		if err != nil {
			return nil, err
		}
		zap.L().Debug("Using following config", zap.String("Config", string(usedConfig)))
	}

	// Validate the Result if it is valid
	return config, config.Validate()
}

// We need to extend the validation function because of https://github.com/go-playground/validator/issues/782
// Otherwise users would not get a clue why their config did not validate (meaningless error message).
func (s *CfAPIConfig) Validate() error {
	// Check config with validator tags
	if err := validator.New().Struct(s); err != nil {
		fields := strings.Split(strings.Split(err.Error(), "'")[1], ".")
		currentField := reflect.TypeOf(s).Elem()
		var resultField reflect.StructField
		// Runtime is O(1) for this loop, as we use FieldByName
		for fieldNumber, fieldName := range fields {
			// The first field is the struct name, we must skip that
			if fieldNumber > 0 {
				field, _ := currentField.FieldByName(fieldName)
				if field.Type.Kind() == reflect.Struct {
					// If the next field is a struct, we itterate deeper
					currentField = field.Type
				} else {
					// If it is not we have a basic datatype so we save that field
					resultField = field
					break
				}
			}
		}
		return fmt.Errorf("error parsing config: %s. values must comply with following schema: \"%s\"."+
			"Read more what the validate tags mean here: https://github.com/go-playground/validator ", err, resultField.Tag.Get("validate"))
	}

	// Check the DB Connection String
	if s.DB.Type == "postgres" {
		_, err := pgx.ParseConfig(s.DB.ConnectionString)
		if err != nil {
			return fmt.Errorf("invalid config.db.connectionstring: %v", err)
		}
	} else {
		_, err := mysql.ParseDSN(s.DB.ConnectionString)
		if err != nil {
			return fmt.Errorf("invalid config.db.connectionstring: %v", err)
		}
	}
	return nil
}
