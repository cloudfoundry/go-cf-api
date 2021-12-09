"use strict";(self.webpackChunkmy_website=self.webpackChunkmy_website||[]).push([[5824],{3905:function(e,t,r){r.d(t,{Zo:function(){return i},kt:function(){return f}});var c=r(7294);function o(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function n(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var c=Object.getOwnPropertySymbols(e);t&&(c=c.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,c)}return r}function l(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?n(Object(r),!0).forEach((function(t){o(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):n(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function a(e,t){if(null==e)return{};var r,c,o=function(e,t){if(null==e)return{};var r,c,o={},n=Object.keys(e);for(c=0;c<n.length;c++)r=n[c],t.indexOf(r)>=0||(o[r]=e[r]);return o}(e,t);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);for(c=0;c<n.length;c++)r=n[c],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(o[r]=e[r])}return o}var u=c.createContext({}),s=function(e){var t=c.useContext(u),r=t;return e&&(r="function"==typeof e?e(t):l(l({},t),e)),r},i=function(e){var t=s(e.components);return c.createElement(u.Provider,{value:t},e.children)},m={inlineCode:"code",wrapper:function(e){var t=e.children;return c.createElement(c.Fragment,{},t)}},p=c.forwardRef((function(e,t){var r=e.components,o=e.mdxType,n=e.originalType,u=e.parentName,i=a(e,["components","mdxType","originalType","parentName"]),p=s(r),f=o,d=p["".concat(u,".").concat(f)]||p[f]||m[f]||n;return r?c.createElement(d,l(l({ref:t},i),{},{components:r})):c.createElement(d,l({ref:t},i))}));function f(e,t){var r=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var n=r.length,l=new Array(n);l[0]=p;var a={};for(var u in t)hasOwnProperty.call(t,u)&&(a[u]=t[u]);a.originalType=e,a.mdxType="string"==typeof e?e:o,l[1]=a;for(var s=2;s<n;s++)l[s]=r[s];return c.createElement.apply(null,l)}return c.createElement.apply(null,r)}p.displayName="MDXCreateElement"},1107:function(e,t,r){r.r(t),r.d(t,{frontMatter:function(){return a},contentTitle:function(){return u},metadata:function(){return s},toc:function(){return i},default:function(){return p}});var c=r(7462),o=r(3366),n=(r(7294),r(3905)),l=["components"],a={},u=void 0,s={unversionedId:"Packages/internal/metrics/custommetrics",id:"Packages/internal/metrics/custommetrics",isDocsHomePage:!1,title:"custommetrics",description:"`go",source:"@site/godocs/Packages/internal/metrics/custommetrics.md",sourceDirName:"Packages/internal/metrics",slug:"/Packages/internal/metrics/custommetrics",permalink:"/cloudfoundry/go-cf-api/godocs/Packages/internal/metrics/custommetrics",editUrl:"https://github.com/cloudfoundry/go-cf-api/edit/main/docs/godocs/Packages/internal/metrics/custommetrics.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"metrics",permalink:"/cloudfoundry/go-cf-api/godocs/Packages/internal/metrics"},next:{title:"db",permalink:"/cloudfoundry/go-cf-api/godocs/Packages/internal/storage/db"}},i=[{value:"Index",id:"index",children:[]},{value:"Constants",id:"constants",children:[]},{value:"type CustomCollector",id:"type-customcollector",children:[{value:"func NewCustomCollector",id:"func-newcustomcollector",children:[]},{value:"func (*CustomCollector) Collect",id:"func-customcollector-collect",children:[]},{value:"func (*CustomCollector) Describe",id:"func-customcollector-describe",children:[]},{value:"func (*CustomCollector) cpuStatus",id:"func-customcollector-cpustatus",children:[]},{value:"func (*CustomCollector) upTime",id:"func-customcollector-uptime",children:[]}]}],m={toc:i};function p(e){var t=e.components,r=(0,o.Z)(e,l);return(0,n.kt)("wrapper",(0,c.Z)({},m,r,{components:t,mdxType:"MDXLayout"}),(0,n.kt)("h1",{id:"custommetrics"},"custommetrics"),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-go"},'import "github.com/cloudfoundry/go-cf-api/internal/metrics/custommetrics"\n')),(0,n.kt)("h2",{id:"index"},"Index"),(0,n.kt)("ul",null,(0,n.kt)("li",{parentName:"ul"},(0,n.kt)("a",{parentName:"li",href:"#constants"},"Constants")),(0,n.kt)("li",{parentName:"ul"},(0,n.kt)("a",{parentName:"li",href:"#type-customcollector"},"type CustomCollector"),(0,n.kt)("ul",{parentName:"li"},(0,n.kt)("li",{parentName:"ul"},(0,n.kt)("a",{parentName:"li",href:"#func-newcustomcollector"},"func NewCustomCollector(startTime time.Time) *CustomCollector")),(0,n.kt)("li",{parentName:"ul"},(0,n.kt)("a",{parentName:"li",href:"#func-customcollector-collect"},"func (e *CustomCollector) Collect(channel chan<- prometheus.Metric)")),(0,n.kt)("li",{parentName:"ul"},(0,n.kt)("a",{parentName:"li",href:"#func-customcollector-describe"},"func (e ",(0,n.kt)("em",{parentName:"a"},"CustomCollector) Describe(ch chan<- "),"prometheus.Desc)")),(0,n.kt)("li",{parentName:"ul"},(0,n.kt)("a",{parentName:"li",href:"#func-customcollector-cpustatus"},"func (e *CustomCollector) cpuStatus() (float64, error)")),(0,n.kt)("li",{parentName:"ul"},(0,n.kt)("a",{parentName:"li",href:"#func-customcollector-uptime"},"func (e *CustomCollector) upTime() float64"))))),(0,n.kt)("h2",{id:"constants"},"Constants"),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-go"},'const namespace = "go_custom_stats"\n')),(0,n.kt)("h2",{id:"type-customcollector"},"type CustomCollector"),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-go"},"type CustomCollector struct {\n    cpuUsage  *prometheus.Desc\n    uptime    *prometheus.Desc\n    startTime time.Time\n}\n")),(0,n.kt)("h3",{id:"func-newcustomcollector"},"func NewCustomCollector"),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-go"},"func NewCustomCollector(startTime time.Time) *CustomCollector\n")),(0,n.kt)("h3",{id:"func-customcollector-collect"},"func ","(","*","CustomCollector",")"," Collect"),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-go"},"func (e *CustomCollector) Collect(channel chan<- prometheus.Metric)\n")),(0,n.kt)("h3",{id:"func-customcollector-describe"},"func ","(","*","CustomCollector",")"," Describe"),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-go"},"func (e *CustomCollector) Describe(ch chan<- *prometheus.Desc)\n")),(0,n.kt)("h3",{id:"func-customcollector-cpustatus"},"func ","(","*","CustomCollector",")"," cpuStatus"),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-go"},"func (e *CustomCollector) cpuStatus() (float64, error)\n")),(0,n.kt)("h3",{id:"func-customcollector-uptime"},"func ","(","*","CustomCollector",")"," upTime"),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre",className:"language-go"},"func (e *CustomCollector) upTime() float64\n")),(0,n.kt)("p",null,"Generated by ",(0,n.kt)("a",{parentName:"p",href:"https://github.com/princjef/gomarkdoc"},"gomarkdoc")))}p.isMDXComponent=!0}}]);