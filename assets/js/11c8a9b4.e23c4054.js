"use strict";(self.webpackChunkmy_website=self.webpackChunkmy_website||[]).push([[8822],{3905:function(e,t,n){n.d(t,{Zo:function(){return s},kt:function(){return f}});var a=n(7294);function i(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function r(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?r(Object(n),!0).forEach((function(t){i(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):r(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,a,i=function(e,t){if(null==e)return{};var n,a,i={},r=Object.keys(e);for(a=0;a<r.length;a++)n=r[a],t.indexOf(n)>=0||(i[n]=e[n]);return i}(e,t);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);for(a=0;a<r.length;a++)n=r[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(i[n]=e[n])}return i}var c=a.createContext({}),p=function(e){var t=a.useContext(c),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},s=function(e){var t=p(e.components);return a.createElement(c.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},g=a.forwardRef((function(e,t){var n=e.components,i=e.mdxType,r=e.originalType,c=e.parentName,s=l(e,["components","mdxType","originalType","parentName"]),g=p(n),f=i,d=g["".concat(c,".").concat(f)]||g[f]||u[f]||r;return n?a.createElement(d,o(o({ref:t},s),{},{components:n})):a.createElement(d,o({ref:t},s))}));function f(e,t){var n=arguments,i=t&&t.mdxType;if("string"==typeof e||i){var r=n.length,o=new Array(r);o[0]=g;var l={};for(var c in t)hasOwnProperty.call(t,c)&&(l[c]=t[c]);l.originalType=e,l.mdxType="string"==typeof e?e:i,o[1]=l;for(var p=2;p<r;p++)o[p]=n[p];return a.createElement.apply(null,o)}return a.createElement.apply(null,n)}g.displayName="MDXCreateElement"},4941:function(e,t,n){n.r(t),n.d(t,{frontMatter:function(){return l},contentTitle:function(){return c},metadata:function(){return p},toc:function(){return s},default:function(){return g}});var a=n(7462),i=n(3366),r=(n(7294),n(3905)),o=["components"],l={},c=void 0,p={unversionedId:"Packages/internal/config",id:"Packages/internal/config",isDocsHomePage:!1,title:"config",description:"`go",source:"@site/godocs/Packages/internal/config.md",sourceDirName:"Packages/internal",slug:"/Packages/internal/config",permalink:"/cloudfoundry/go-cf-api/godocs/Packages/internal/config",editUrl:"https://github.com/cloudfoundry/go-cf-api/edit/main/docs/godocs/Packages/internal/config.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"timefilters",permalink:"/cloudfoundry/go-cf-api/godocs/Packages/internal/apicommon/v3/timefilters"},next:{title:"helpers",permalink:"/cloudfoundry/go-cf-api/godocs/Packages/internal/helpers"}},s=[{value:"Index",id:"index",children:[]},{value:"Variables",id:"variables",children:[]},{value:"type AppSSH",id:"type-appssh",children:[]},{value:"type CfAPIConfig",id:"type-cfapiconfig",children:[{value:"func Get",id:"func-get",children:[]},{value:"func (*CfAPIConfig) Validate",id:"func-cfapiconfig-validate",children:[]}]},{value:"type DBConfig",id:"type-dbconfig",children:[]},{value:"type InfoConfig",id:"type-infoconfig",children:[]},{value:"type RateLimitConf",id:"type-ratelimitconf",children:[]},{value:"type URLs",id:"type-urls",children:[]},{value:"type UaaConfig",id:"type-uaaconfig",children:[]},{value:"type ZapConfig",id:"type-zapconfig",children:[]}],u={toc:s};function g(e){var t=e.components,n=(0,i.Z)(e,o);return(0,r.kt)("wrapper",(0,a.Z)({},u,n,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("h1",{id:"config"},"config"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-go"},'import "github.com/cloudfoundry/go-cf-api/internal/config"\n')),(0,r.kt)("h2",{id:"index"},"Index"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("a",{parentName:"li",href:"#variables"},"Variables")),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("a",{parentName:"li",href:"#type-appssh"},"type AppSSH")),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("a",{parentName:"li",href:"#type-cfapiconfig"},"type CfAPIConfig"),(0,r.kt)("ul",{parentName:"li"},(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("a",{parentName:"li",href:"#func-get"},"func Get(configFile string) (*CfAPIConfig, error)")),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("a",{parentName:"li",href:"#func-cfapiconfig-validate"},"func (s *CfAPIConfig) Validate() error")))),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("a",{parentName:"li",href:"#type-dbconfig"},"type DBConfig")),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("a",{parentName:"li",href:"#type-infoconfig"},"type InfoConfig")),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("a",{parentName:"li",href:"#type-ratelimitconf"},"type RateLimitConf")),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("a",{parentName:"li",href:"#type-urls"},"type URLs")),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("a",{parentName:"li",href:"#type-uaaconfig"},"type UaaConfig")),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("a",{parentName:"li",href:"#type-zapconfig"},"type ZapConfig"))),(0,r.kt)("h2",{id:"variables"},"Variables"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-go"},"var defaults []byte\n")),(0,r.kt)("h2",{id:"type-appssh"},"type AppSSH"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-go"},'type AppSSH struct {\n    Endpoint           string `yaml:"endpoint" validate:"required,url"`\n    OAuthClient        string `yaml:"oauth_client" validate:"required"`\n    HostKeyFingerprint string `yaml:"host_key_fingerprint"`\n}\n')),(0,r.kt)("h2",{id:"type-cfapiconfig"},"type CfAPIConfig"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-go"},'type CfAPIConfig struct {\n    Listen           string        `yaml:"listen"`\n    ExternalDomain   string        `yaml:"external_domain" validate:"required"`\n    ExternalProtocol string        `yaml:"external_protocol" validate:"required"`\n    Info             InfoConfig    `yaml:"info"`\n    DB               DBConfig      `yaml:"db"`\n    Log              ZapConfig     `yaml:"log"`\n    RateLimit        RateLimitConf `yaml:"rate_limiting"`\n    Uaa              UaaConfig     `yaml:"uaa"`\n    URLs             URLs          `yaml:"urls"`\n    AppSSH           AppSSH        `yaml:"app_ssh"`\n}\n')),(0,r.kt)("h3",{id:"func-get"},"func Get"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-go"},"func Get(configFile string) (*CfAPIConfig, error)\n")),(0,r.kt)("p",null,"Parses the config once\\, otherwise just returns it Priority from lowest to highest: Default Values\\, Config File\\, Environment Variables After setting the Config in this order it is validated according to struct tags If validation is not succefull we have a invalid config and thus throw a FATAL"),(0,r.kt)("h3",{id:"func-cfapiconfig-validate"},"func ","(","*","CfAPIConfig",")"," Validate"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-go"},"func (s *CfAPIConfig) Validate() error\n")),(0,r.kt)("p",null,"We need to extend the validation function because of ",(0,r.kt)("a",{parentName:"p",href:"https://github.com/go-playground/validator/issues/782"},"https://github.com/go-playground/validator/issues/782")," Otherwise users would not get a clue why their config did not validate ","(","meaningless error message",")","."),(0,r.kt)("h2",{id:"type-dbconfig"},"type DBConfig"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-go"},'type DBConfig struct {\n    // For connection string documentation see\n    // Postgres: https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-PARAMKEYWORDS\n    // Mysql: https://github.com/go-sql-driver/mysql#dsn-data-source-name\n    ConnectionString string    `yaml:"connection_string" validate:"required"`\n    Type             string    `yaml:"type" validate:"required,oneof=postgres mysql"`\n    Log              ZapConfig `yaml:"log"`\n    MaxConnections   int       `yaml:"max_connections" validate:"gte=1,lte=1000"`\n    MinConnections   int       `yaml:"min_connections" validate:"gte=1,lte=1000,ltefield=MaxConnections"`\n}\n')),(0,r.kt)("h2",{id:"type-infoconfig"},"type InfoConfig"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-go"},'type InfoConfig struct {\n    Name                     string `yaml:"name" validate:"required"`\n    Build                    string `yaml:"build" validate:"required"`\n    Version                  int    `yaml:"version" validate:"gte=0"`\n    SupportAddress           string `yaml:"support_address" validate:"required"`\n    Description              string `yaml:"description" validate:"required"`\n    MinCliVersion            string `yaml:"min_cli_version"`\n    MinRecommendedCliVersion string `yaml:"min_recommended_cli_version"`\n}\n')),(0,r.kt)("h2",{id:"type-ratelimitconf"},"type RateLimitConf"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-go"},'type RateLimitConf struct {\n    Enabled                        bool          `yaml:"enabled"`\n    PerProcessGeneralLimit         int           `yaml:"per_process_general_limit"`\n    GlobalGeneralLimit             int           `yaml:"global_general_limit"`\n    PerProcessUnauthenticatedLimit int           `yaml:"per_process_unauthenticated_limit"`\n    GlobalUnauthenticatedLimit     int           `yaml:"global_unauthenticated_limit"`\n    ResetInterval                  time.Duration `yaml:"reset_interval"`\n}\n')),(0,r.kt)("h2",{id:"type-urls"},"type URLs"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-go"},'type URLs struct {\n    LogCache   string  `yaml:"log_cache" validate:"required,url"`\n    LogStream  string  `yaml:"log_stream" validate:"required,url"`\n    Doppler    string  `yaml:"doppler" validate:"required,url"`\n    Login      string  `yaml:"login" validate:"required,url"`\n    UAA        string  `yaml:"uaa" validate:"required,url"`\n    RoutingAPI *string `yaml:"routing_api" validate:"omitempty,url"`\n}\n')),(0,r.kt)("h2",{id:"type-uaaconfig"},"type UaaConfig"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-go"},'type UaaConfig struct {\n    URL    string                      `yaml:"url"`\n    Client promconfig.HTTPClientConfig `yaml:"client"`\n}\n')),(0,r.kt)("h2",{id:"type-zapconfig"},"type ZapConfig"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-go"},'type ZapConfig struct {\n    Level      string `yaml:"level" validate:"required,oneof=debug info warn error dpanic panic fatal"`\n    Production bool   `yaml:"production"`\n}\n')),(0,r.kt)("p",null,"Generated by ",(0,r.kt)("a",{parentName:"p",href:"https://github.com/princjef/gomarkdoc"},"gomarkdoc")))}g.isMDXComponent=!0}}]);