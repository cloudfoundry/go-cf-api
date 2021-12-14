"use strict";(self.webpackChunkmy_website=self.webpackChunkmy_website||[]).push([[966],{3905:function(e,r,n){n.d(r,{Zo:function(){return f},kt:function(){return d}});var t=n(7294);function a(e,r,n){return r in e?Object.defineProperty(e,r,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[r]=n,e}function o(e,r){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var t=Object.getOwnPropertySymbols(e);r&&(t=t.filter((function(r){return Object.getOwnPropertyDescriptor(e,r).enumerable}))),n.push.apply(n,t)}return n}function u(e){for(var r=1;r<arguments.length;r++){var n=null!=arguments[r]?arguments[r]:{};r%2?o(Object(n),!0).forEach((function(r){a(e,r,n[r])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(r){Object.defineProperty(e,r,Object.getOwnPropertyDescriptor(n,r))}))}return e}function c(e,r){if(null==e)return{};var n,t,a=function(e,r){if(null==e)return{};var n,t,a={},o=Object.keys(e);for(t=0;t<o.length;t++)n=o[t],r.indexOf(n)>=0||(a[n]=e[n]);return a}(e,r);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(t=0;t<o.length;t++)n=o[t],r.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var l=t.createContext({}),i=function(e){var r=t.useContext(l),n=r;return e&&(n="function"==typeof e?e(r):u(u({},r),e)),n},f=function(e){var r=i(e.components);return t.createElement(l.Provider,{value:r},e.children)},p={inlineCode:"code",wrapper:function(e){var r=e.children;return t.createElement(t.Fragment,{},r)}},s=t.forwardRef((function(e,r){var n=e.components,a=e.mdxType,o=e.originalType,l=e.parentName,f=c(e,["components","mdxType","originalType","parentName"]),s=i(n),d=a,m=s["".concat(l,".").concat(d)]||s[d]||p[d]||o;return n?t.createElement(m,u(u({ref:r},f),{},{components:n})):t.createElement(m,u({ref:r},f))}));function d(e,r){var n=arguments,a=r&&r.mdxType;if("string"==typeof e||a){var o=n.length,u=new Array(o);u[0]=s;var c={};for(var l in r)hasOwnProperty.call(r,l)&&(c[l]=r[l]);c.originalType=e,c.mdxType="string"==typeof e?e:a,u[1]=c;for(var i=2;i<o;i++)u[i]=n[i];return t.createElement.apply(null,u)}return t.createElement.apply(null,n)}s.displayName="MDXCreateElement"},2174:function(e,r,n){n.r(r),n.d(r,{frontMatter:function(){return c},contentTitle:function(){return l},metadata:function(){return i},toc:function(){return f},default:function(){return s}});var t=n(7462),a=n(3366),o=(n(7294),n(3905)),u=["components"],c={},l=void 0,i={unversionedId:"Packages/internal/apicommon/v3",id:"Packages/internal/apicommon/v3",title:"v3",description:"Index",source:"@site/godocs/Packages/internal/apicommon/v3.md",sourceDirName:"Packages/internal/apicommon",slug:"/Packages/internal/apicommon/v3",permalink:"/go-cf-api/godocs/Packages/internal/apicommon/v3",editUrl:"https://github.com/cloudfoundry/go-cf-api/edit/main/docs/godocs/Packages/internal/apicommon/v3.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"securitygroups",permalink:"/go-cf-api/godocs/Packages/internal/api/v3/securitygroups"},next:{title:"auth",permalink:"/go-cf-api/godocs/Packages/internal/apicommon/v3/auth"}},f=[{value:"Index",id:"index",children:[],level:2},{value:"func GetResourcePath",id:"func-getresourcepath",children:[],level:2},{value:"type CfAPIError",id:"type-cfapierror",children:[{value:"func BadQueryParameter",id:"func-badqueryparameter",children:[],level:3},{value:"func InvalidAuthToken",id:"func-invalidauthtoken",children:[],level:3},{value:"func NotAuthenticated",id:"func-notauthenticated",children:[],level:3},{value:"func NotAuthorized",id:"func-notauthorized",children:[],level:3},{value:"func ResourceNotFound",id:"func-resourcenotfound",children:[],level:3},{value:"func TooManyRequests",id:"func-toomanyrequests",children:[],level:3},{value:"func UnknownError",id:"func-unknownerror",children:[],level:3},{value:"func UnprocessableEntity",id:"func-unprocessableentity",children:[],level:3},{value:"func (*CfAPIError) Error",id:"func-cfapierror-error",children:[],level:3},{value:"func (*CfAPIError) Unwrap",id:"func-cfapierror-unwrap",children:[],level:3}],level:2},{value:"type CfAPIErrors",id:"type-cfapierrors",children:[{value:"func AsErrors",id:"func-aserrors",children:[],level:3}],level:2}],p={toc:f};function s(e){var r=e.components,n=(0,a.Z)(e,u);return(0,o.kt)("wrapper",(0,t.Z)({},p,n,{components:r,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"v3"},"v3"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},'import "github.com/cloudfoundry/go-cf-api/internal/apicommon/v3"\n')),(0,o.kt)("h2",{id:"index"},"Index"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"#func-getresourcepath"},"func GetResourcePath(c echo.Context) string")),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"#type-cfapierror"},"type CfAPIError"),(0,o.kt)("ul",{parentName:"li"},(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"#func-badqueryparameter"},"func BadQueryParameter(err error) *CfAPIError")),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"#func-invalidauthtoken"},"func InvalidAuthToken(err error) *CfAPIError")),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"#func-notauthenticated"},"func NotAuthenticated(err error) *CfAPIError")),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"#func-notauthorized"},"func NotAuthorized(err error) *CfAPIError")),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"#func-resourcenotfound"},"func ResourceNotFound(resourceType string, err error) *CfAPIError")),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"#func-toomanyrequests"},"func TooManyRequests(err error) *CfAPIError")),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"#func-unknownerror"},"func UnknownError(err error) *CfAPIError")),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"#func-unprocessableentity"},"func UnprocessableEntity(detail string, err error) *CfAPIError")),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"#func-cfapierror-error"},"func (e *CfAPIError) Error() string")),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"#func-cfapierror-unwrap"},"func (e *CfAPIError) Unwrap() error")))),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"#type-cfapierrors"},"type CfAPIErrors"),(0,o.kt)("ul",{parentName:"li"},(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"#func-aserrors"},"func AsErrors(errors ...CfAPIError) CfAPIErrors"))))),(0,o.kt)("h2",{id:"func-getresourcepath"},"func GetResourcePath"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"func GetResourcePath(c echo.Context) string\n")),(0,o.kt)("p",null,"GetResourcePath returns the full url that was called by the web request"),(0,o.kt)("h2",{id:"type-cfapierror"},"type CfAPIError"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},'type CfAPIError struct {\n    HTTPStatus int    `json:"-"`\n    Detail     string `json:"detail,omitempty"`\n    Title      string `json:"title"`\n    Code       int    `json:"code"`\n    Err        error  `json:"-"`\n}\n')),(0,o.kt)("h3",{id:"func-badqueryparameter"},"func BadQueryParameter"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"func BadQueryParameter(err error) *CfAPIError\n")),(0,o.kt)("h3",{id:"func-invalidauthtoken"},"func InvalidAuthToken"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"func InvalidAuthToken(err error) *CfAPIError\n")),(0,o.kt)("h3",{id:"func-notauthenticated"},"func NotAuthenticated"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"func NotAuthenticated(err error) *CfAPIError\n")),(0,o.kt)("h3",{id:"func-notauthorized"},"func NotAuthorized"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"func NotAuthorized(err error) *CfAPIError\n")),(0,o.kt)("h3",{id:"func-resourcenotfound"},"func ResourceNotFound"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"func ResourceNotFound(resourceType string, err error) *CfAPIError\n")),(0,o.kt)("h3",{id:"func-toomanyrequests"},"func TooManyRequests"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"func TooManyRequests(err error) *CfAPIError\n")),(0,o.kt)("h3",{id:"func-unknownerror"},"func UnknownError"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"func UnknownError(err error) *CfAPIError\n")),(0,o.kt)("h3",{id:"func-unprocessableentity"},"func UnprocessableEntity"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"func UnprocessableEntity(detail string, err error) *CfAPIError\n")),(0,o.kt)("h3",{id:"func-cfapierror-error"},"func ","(","*","CfAPIError",")"," Error"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"func (e *CfAPIError) Error() string\n")),(0,o.kt)("h3",{id:"func-cfapierror-unwrap"},"func ","(","*","CfAPIError",")"," Unwrap"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"func (e *CfAPIError) Unwrap() error\n")),(0,o.kt)("h2",{id:"type-cfapierrors"},"type CfAPIErrors"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},'type CfAPIErrors struct {\n    Errors []CfAPIError `json:"errors"`\n}\n')),(0,o.kt)("h3",{id:"func-aserrors"},"func AsErrors"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"func AsErrors(errors ...CfAPIError) CfAPIErrors\n")),(0,o.kt)("p",null,"Generated by ",(0,o.kt)("a",{parentName:"p",href:"https://github.com/princjef/gomarkdoc"},"gomarkdoc")))}s.isMDXComponent=!0}}]);