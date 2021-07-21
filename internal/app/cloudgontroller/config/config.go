package config

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/creasty/defaults"
	"github.com/go-playground/validator"
	"github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v4"
	"github.com/spf13/viper"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/helpers"
	log "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/logging/enums"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

var once sync.Once
var config *GopilotConfig

type GopilotConfig struct {
	Port      int `env:"GOPILOT_PORT" validate:"required,gte=0,lte=65535" default:"8080"`
	DB        DBConfig
	Log       ZapConfig
	RateLimit RateLimitConf
}

type DBConfig struct {
	// For connection string documentation see
	// Postgres: https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-PARAMKEYWORDS
	// Mysql: https://github.com/go-sql-driver/mysql#dsn-data-source-name
	ConnectionString string `env:"GOPILOT_DB_CONNECTIONSTRING" validate:"required"`
	Type             string `env:"GOPILOT_DB_TYPE" validate:"required,oneof=postgres mysql"`
	Create           bool   `env:"GOPILOT_DB_CREATE" default:"true"`
	Migrate          bool   `env:"GOPILOT_DB_MIGRATE" default:"true"`
	Log              ZapConfig
}

type ZapConfig struct {
	Level      string `env:"GOPILOT_LOG_LEVEL" validate:"required,oneof=debug info warn error dpanic panic fatal" default:"info"`
	Production bool   `env:"GOPILOT_LOG_PRODUCTION" default:"true"`
}

type RateLimitConf struct {
	Enabled   bool          `env:"GOPILOT_RATELIMIT_ENABLED" default:"true"`
	Rate      int           `env:"GOPILOT_RATELIMIT_RATE" validate:"gte=0,lte=1000" default:"10"`
	Burst     int           `env:"GOPILOT_RATELIMIT_BURST" validate:"gte=0,lte=1000" default:"20"`
	ExpiresIn time.Duration `env:"GOPILOT_RATELIMIT_EXPIRESIN" default:"180"` // Default in Seconds
}

// Parses the config once, otherwise just returns it
// Priority from lowest to highest: Default Values, Config File, Environment Variables
// After setting the Config in this order it is validated according to struct tags
// If validation is not succefull we have a invalid config and thus throw a FATAL
func Get(params ...string) *GopilotConfig {
	once.Do(func() {
		tmp := &GopilotConfig{}
		// Default Values
		helpers.CheckErrFatal(defaults.Set(tmp))
		// If Config File is present, Parse it into a struct
		if len(params) == 1 {
			viper.SetConfigFile(params[0])
			err := viper.ReadInConfig()
			if err == nil {
				zap.L().Info(fmt.Sprintf("Using config file %s", viper.ConfigFileUsed()), zap.String(log.Tags.File, viper.ConfigFileUsed()))
			} else {
				helpers.CheckErrFatal(err)
			}
			err = viper.Unmarshal(tmp)
			helpers.CheckErrFatal(err)
		}
		// Overwrite values with env Variables
		helpers.CheckErrFatal(env.Parse(tmp))
		if zap.L().Core().Enabled(zap.DebugLevel) {
			usedConfig, err := yaml.Marshal(config)
			helpers.CheckErrFatal(err)
			zap.L().Debug("Using following config", zap.String("Config", string(usedConfig)))
		}
		// Validate the Result if it is valid
		helpers.CheckErrFatal(tmp.Validate())
		config = tmp
	})
	return config
}

// SetDefaults implements defaults.Setter interface
func (s *GopilotConfig) SetDefaults() {
	s.RateLimit.ExpiresIn = s.RateLimit.ExpiresIn * time.Second
}

// We need to extend the validation function because of https://github.com/go-playground/validator/issues/782
// Otherwise users would not get a clue why their config did not validate (meaningless error message)
func (s *GopilotConfig) Validate() error {
	// Check config with validator tags
	if err := validator.New().Struct(s); err != nil {
		fields := strings.Split(strings.Split(err.Error(), "'")[1], ".")
		currentField := reflect.TypeOf(s).Elem()
		var resultField reflect.StructField
		// Runtime is O(1) for this loop, as we use FieldByName
		for fieldNumber, fieldName := range fields {
			// The first field is the struct name, we must skip that
			if fieldNumber > 0 {
				f, _ := currentField.FieldByName(fieldName)
				if f.Type.Kind() == reflect.Struct {
					// If the next field is a struct, we itterate deeper
					currentField = f.Type
				} else {
					// If it is not we have a basic datatype so we save that field
					resultField = f
					break
				}
			}
		}
		return fmt.Errorf("error parsing config: %s. values must comply with following schema: \"%s\". Read more what the validate tags mean here: https://github.com/go-playground/validator ", err, resultField.Tag.Get("validate"))
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
