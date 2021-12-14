"use strict";(self.webpackChunkmy_website=self.webpackChunkmy_website||[]).push([[3887],{3905:function(e,n,t){t.d(n,{Zo:function(){return s},kt:function(){return d}});var r=t(7294);function o(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function i(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);n&&(r=r.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,r)}return t}function a(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?i(Object(t),!0).forEach((function(n){o(e,n,t[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):i(Object(t)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))}))}return e}function c(e,n){if(null==e)return{};var t,r,o=function(e,n){if(null==e)return{};var t,r,o={},i=Object.keys(e);for(r=0;r<i.length;r++)t=i[r],n.indexOf(t)>=0||(o[t]=e[t]);return o}(e,n);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)t=i[r],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(o[t]=e[t])}return o}var l=r.createContext({}),u=function(e){var n=r.useContext(l),t=n;return e&&(t="function"==typeof e?e(n):a(a({},n),e)),t},s=function(e){var n=u(e.components);return r.createElement(l.Provider,{value:n},e.children)},f={inlineCode:"code",wrapper:function(e){var n=e.children;return r.createElement(r.Fragment,{},n)}},p=r.forwardRef((function(e,n){var t=e.components,o=e.mdxType,i=e.originalType,l=e.parentName,s=c(e,["components","mdxType","originalType","parentName"]),p=u(t),d=o,g=p["".concat(l,".").concat(d)]||p[d]||f[d]||i;return t?r.createElement(g,a(a({ref:n},s),{},{components:t})):r.createElement(g,a({ref:n},s))}));function d(e,n){var t=arguments,o=n&&n.mdxType;if("string"==typeof e||o){var i=t.length,a=new Array(i);a[0]=p;var c={};for(var l in n)hasOwnProperty.call(n,l)&&(c[l]=n[l]);c.originalType=e,c.mdxType="string"==typeof e?e:o,a[1]=c;for(var u=2;u<i;u++)a[u]=t[u];return r.createElement.apply(null,a)}return r.createElement.apply(null,t)}p.displayName="MDXCreateElement"},6168:function(e,n,t){t.r(n),t.d(n,{frontMatter:function(){return c},contentTitle:function(){return l},metadata:function(){return u},toc:function(){return s},default:function(){return p}});var r=t(7462),o=t(3366),i=(t(7294),t(3905)),a=["components"],c={},l=void 0,u={unversionedId:"Packages/internal/api",id:"Packages/internal/api",title:"api",description:"Index",source:"@site/godocs/Packages/internal/api.md",sourceDirName:"Packages/internal",slug:"/Packages/internal/api",permalink:"/go-cf-api/godocs/Packages/internal/api",editUrl:"https://github.com/cloudfoundry/go-cf-api/edit/main/docs/godocs/Packages/internal/api.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"cmd",permalink:"/go-cf-api/godocs/Packages/cmd"},next:{title:"apicommon",permalink:"/go-cf-api/godocs/Packages/internal/apicommon"}},s=[{value:"Index",id:"index",children:[],level:2},{value:"func NewRootEndpoint",id:"func-newrootendpoint",children:[],level:2},{value:"func RegisterHandlers",id:"func-registerhandlers",children:[],level:2},{value:"func registerV3Handlers",id:"func-registerv3handlers",children:[],level:2},{value:"func routingLink",id:"func-routinglink",children:[],level:2},{value:"type Root",id:"type-root",children:[],level:2},{value:"type RootLinks",id:"type-rootlinks",children:[],level:2}],f={toc:s};function p(e){var n=e.components,t=(0,o.Z)(e,a);return(0,i.kt)("wrapper",(0,r.Z)({},f,t,{components:n,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"api"},"api"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},'import "github.com/cloudfoundry/go-cf-api/internal/api"\n')),(0,i.kt)("h2",{id:"index"},"Index"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#func-newrootendpoint"},"func NewRootEndpoint(config *config.CfAPIConfig) echo.HandlerFunc")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#func-registerhandlers"},"func RegisterHandlers(echoInstance ",(0,i.kt)("em",{parentName:"a"},"echo.Echo, database "),"sql.DB, jwtMiddleware echo.MiddlewareFunc, rateLimitMiddleware echo.MiddlewareFunc, conf *config.CfAPIConfig)")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#func-registerv3handlers"},"func registerV3Handlers(echoInstance ",(0,i.kt)("em",{parentName:"a"},"echo.Echo, database "),"sql.DB, jwtMiddleware echo.MiddlewareFunc, rateLimitMiddleware echo.MiddlewareFunc, conf *config.CfAPIConfig)")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#func-routinglink"},"func routingLink(config ",(0,i.kt)("em",{parentName:"a"},"config.CfAPIConfig) "),"info.Link")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#type-root"},"type Root")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#type-rootlinks"},"type RootLinks"))),(0,i.kt)("h2",{id:"func-newrootendpoint"},"func NewRootEndpoint"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},"func NewRootEndpoint(config *config.CfAPIConfig) echo.HandlerFunc\n")),(0,i.kt)("h2",{id:"func-registerhandlers"},"func RegisterHandlers"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},"func RegisterHandlers(echoInstance *echo.Echo, database *sql.DB, jwtMiddleware echo.MiddlewareFunc, rateLimitMiddleware echo.MiddlewareFunc, conf *config.CfAPIConfig)\n")),(0,i.kt)("h2",{id:"func-registerv3handlers"},"func registerV3Handlers"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},"func registerV3Handlers(echoInstance *echo.Echo, database *sql.DB, jwtMiddleware echo.MiddlewareFunc, rateLimitMiddleware echo.MiddlewareFunc, conf *config.CfAPIConfig)\n")),(0,i.kt)("h2",{id:"func-routinglink"},"func routingLink"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},"func routingLink(config *config.CfAPIConfig) *info.Link\n")),(0,i.kt)("h2",{id:"type-root"},"type Root"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},'type Root struct {\n    Links RootLinks `json:"links"`\n}\n')),(0,i.kt)("h2",{id:"type-rootlinks"},"type RootLinks"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},'type RootLinks struct {\n    AppSSH            *info.Link `json:"app_ssh"`\n    BitsService       *info.Link `json:"bits_service"`\n    CloudControllerV2 *info.Link `json:"cloud_controller_v2"`\n    CloudControllerV3 *info.Link `json:"cloud_controller_v3"`\n    Credhub           *info.Link `json:"credhub"`\n    LogCache          *info.Link `json:"log_cache"`\n    LogStream         *info.Link `json:"log_stream"`\n    Logging           *info.Link `json:"logging"`\n    Login             *info.Link `json:"login"`\n    NetworkPolicyV0   *info.Link `json:"network_policy_v0"`\n    NetworkPolicyV1   *info.Link `json:"network_policy_v1"`\n    Routing           *info.Link `json:"routing"`\n    Self              *info.Link `json:"self"`\n    UAA               *info.Link `json:"uaa"`\n}\n')),(0,i.kt)("p",null,"Generated by ",(0,i.kt)("a",{parentName:"p",href:"https://github.com/princjef/gomarkdoc"},"gomarkdoc")))}p.isMDXComponent=!0}}]);