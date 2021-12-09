"use strict";(self.webpackChunkmy_website=self.webpackChunkmy_website||[]).push([[4193],{3905:function(e,n,t){t.d(n,{Zo:function(){return u},kt:function(){return d}});var r=t(7294);function i(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function o(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);n&&(r=r.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,r)}return t}function a(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?o(Object(t),!0).forEach((function(n){i(e,n,t[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):o(Object(t)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))}))}return e}function c(e,n){if(null==e)return{};var t,r,i=function(e,n){if(null==e)return{};var t,r,i={},o=Object.keys(e);for(r=0;r<o.length;r++)t=o[r],n.indexOf(t)>=0||(i[t]=e[t]);return i}(e,n);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(r=0;r<o.length;r++)t=o[r],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(i[t]=e[t])}return i}var l=r.createContext({}),p=function(e){var n=r.useContext(l),t=n;return e&&(t="function"==typeof e?e(n):a(a({},n),e)),t},u=function(e){var n=p(e.components);return r.createElement(l.Provider,{value:n},e.children)},s={inlineCode:"code",wrapper:function(e){var n=e.children;return r.createElement(r.Fragment,{},n)}},f=r.forwardRef((function(e,n){var t=e.components,i=e.mdxType,o=e.originalType,l=e.parentName,u=c(e,["components","mdxType","originalType","parentName"]),f=p(t),d=i,g=f["".concat(l,".").concat(d)]||f[d]||s[d]||o;return t?r.createElement(g,a(a({ref:n},u),{},{components:t})):r.createElement(g,a({ref:n},u))}));function d(e,n){var t=arguments,i=n&&n.mdxType;if("string"==typeof e||i){var o=t.length,a=new Array(o);a[0]=f;var c={};for(var l in n)hasOwnProperty.call(n,l)&&(c[l]=n[l]);c.originalType=e,c.mdxType="string"==typeof e?e:i,a[1]=c;for(var p=2;p<o;p++)a[p]=t[p];return r.createElement.apply(null,a)}return r.createElement.apply(null,t)}f.displayName="MDXCreateElement"},6482:function(e,n,t){t.r(n),t.d(n,{frontMatter:function(){return c},contentTitle:function(){return l},metadata:function(){return p},toc:function(){return u},default:function(){return f}});var r=t(7462),i=t(3366),o=(t(7294),t(3905)),a=["components"],c={},l=void 0,p={unversionedId:"Packages/internal/api/v3/info",id:"Packages/internal/api/v3/info",isDocsHomePage:!1,title:"info",description:"`go",source:"@site/godocs/Packages/internal/api/v3/info.md",sourceDirName:"Packages/internal/api/v3",slug:"/Packages/internal/api/v3/info",permalink:"/cloudfoundry/go-cf-api/godocs/Packages/internal/api/v3/info",editUrl:"https://github.com/cloudfoundry/go-cf-api/edit/main/docs/godocs/Packages/internal/api/v3/info.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"buildpacks",permalink:"/cloudfoundry/go-cf-api/godocs/Packages/internal/api/v3/buildpacks"},next:{title:"securitygroups",permalink:"/cloudfoundry/go-cf-api/godocs/Packages/internal/api/v3/securitygroups"}},u=[{value:"Index",id:"index",children:[]},{value:"func ExternalURL",id:"func-externalurl",children:[]},{value:"func NewV3InfoEndpoint",id:"func-newv3infoendpoint",children:[]},{value:"type CLIVersion",id:"type-cliversion",children:[]},{value:"type Link",id:"type-link",children:[]},{value:"type V3Info",id:"type-v3info",children:[]},{value:"type V3Links",id:"type-v3links",children:[]}],s={toc:u};function f(e){var n=e.components,t=(0,i.Z)(e,a);return(0,o.kt)("wrapper",(0,r.Z)({},s,t,{components:n,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"info"},"info"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},'import "github.com/cloudfoundry/go-cf-api/internal/api/v3/info"\n')),(0,o.kt)("h2",{id:"index"},"Index"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"#func-externalurl"},"func ExternalURL(endpoint string, conf *config.CfAPIConfig) string")),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"#func-newv3infoendpoint"},"func NewV3InfoEndpoint(config *config.CfAPIConfig) echo.HandlerFunc")),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"#type-cliversion"},"type CLIVersion")),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"#type-link"},"type Link")),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"#type-v3info"},"type V3Info")),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"#type-v3links"},"type V3Links"))),(0,o.kt)("h2",{id:"func-externalurl"},"func ExternalURL"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"func ExternalURL(endpoint string, conf *config.CfAPIConfig) string\n")),(0,o.kt)("h2",{id:"func-newv3infoendpoint"},"func NewV3InfoEndpoint"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"func NewV3InfoEndpoint(config *config.CfAPIConfig) echo.HandlerFunc\n")),(0,o.kt)("h2",{id:"type-cliversion"},"type CLIVersion"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},'type CLIVersion struct {\n    Minimum     string `json:"minimum"`\n    Recommended string `json:"recommended"`\n}\n')),(0,o.kt)("h2",{id:"type-link"},"type Link"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},'type Link struct {\n    HREF string            `json:"href"`\n    Meta map[string]string `json:"meta,omitempty"`\n}\n')),(0,o.kt)("h2",{id:"type-v3info"},"type V3Info"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},'type V3Info struct {\n    Name        string     `json:"name"`\n    Description string     `json:"description"`\n    Build       string     `json:"build"`\n    Version     int        `json:"version"`\n    CLIVersion  CLIVersion `json:"cli_version"`\n    Links       V3Links    `json:"links"`\n}\n')),(0,o.kt)("h2",{id:"type-v3links"},"type V3Links"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},'type V3Links struct {\n    Self    Link `json:"self"`\n    Support Link `json:"support"`\n}\n')),(0,o.kt)("p",null,"Generated by ",(0,o.kt)("a",{parentName:"p",href:"https://github.com/princjef/gomarkdoc"},"gomarkdoc")))}f.isMDXComponent=!0}}]);