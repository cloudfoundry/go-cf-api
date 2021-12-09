"use strict";(self.webpackChunkmy_website=self.webpackChunkmy_website||[]).push([[237],{3905:function(e,t,n){n.d(t,{Zo:function(){return p},kt:function(){return m}});var a=n(7294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function u(e,t){if(null==e)return{};var n,a,r=function(e,t){if(null==e)return{};var n,a,r={},i=Object.keys(e);for(a=0;a<i.length;a++)n=i[a],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(a=0;a<i.length;a++)n=i[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var l=a.createContext({}),c=function(e){var t=a.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},p=function(e){var t=c(e.components);return a.createElement(l.Provider,{value:t},e.children)},s={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},g=a.forwardRef((function(e,t){var n=e.components,r=e.mdxType,i=e.originalType,l=e.parentName,p=u(e,["components","mdxType","originalType","parentName"]),g=c(n),m=r,f=g["".concat(l,".").concat(m)]||g[m]||s[m]||i;return n?a.createElement(f,o(o({ref:t},p),{},{components:n})):a.createElement(f,o({ref:t},p))}));function m(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var i=n.length,o=new Array(i);o[0]=g;var u={};for(var l in t)hasOwnProperty.call(t,l)&&(u[l]=t[l]);u.originalType=e,u.mdxType="string"==typeof e?e:r,o[1]=u;for(var c=2;c<i;c++)o[c]=n[c];return a.createElement.apply(null,o)}return a.createElement.apply(null,n)}g.displayName="MDXCreateElement"},6923:function(e,t,n){n.r(t),n.d(t,{frontMatter:function(){return u},contentTitle:function(){return l},metadata:function(){return c},toc:function(){return p},default:function(){return g}});var a=n(7462),r=n(3366),i=(n(7294),n(3905)),o=["components"],u={},l=void 0,c={unversionedId:"Packages/internal/apicommon/v3/pagination",id:"Packages/internal/apicommon/v3/pagination",isDocsHomePage:!1,title:"pagination",description:"`go",source:"@site/godocs/Packages/internal/apicommon/v3/pagination.md",sourceDirName:"Packages/internal/apicommon/v3",slug:"/Packages/internal/apicommon/v3/pagination",permalink:"/cloudfoundry/go-cf-api/godocs/Packages/internal/apicommon/v3/pagination",editUrl:"https://github.com/cloudfoundry/go-cf-api/edit/main/docs/godocs/Packages/internal/apicommon/v3/pagination.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"mocks",permalink:"/cloudfoundry/go-cf-api/godocs/Packages/internal/apicommon/v3/metadata/mocks"},next:{title:"permissions",permalink:"/cloudfoundry/go-cf-api/godocs/Packages/internal/apicommon/v3/permissions"}},p=[{value:"Index",id:"index",children:[]},{value:"func totalPages",id:"func-totalpages",children:[]},{value:"type Link",id:"type-link",children:[{value:"func GetResourcePathLink",id:"func-getresourcepathlink",children:[]},{value:"func GetResourcePathLinkWithMethod",id:"func-getresourcepathlinkwithmethod",children:[]},{value:"func nextLink",id:"func-nextlink",children:[]},{value:"func previousLink",id:"func-previouslink",children:[]}]},{value:"type Pagination",id:"type-pagination",children:[{value:"func NewPagination",id:"func-newpagination",children:[]}]},{value:"type Params",id:"type-params",children:[{value:"func Default",id:"func-default",children:[]}]}],s={toc:p};function g(e){var t=e.components,n=(0,r.Z)(e,o);return(0,i.kt)("wrapper",(0,a.Z)({},s,n,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"pagination"},"pagination"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},'import "github.com/cloudfoundry/go-cf-api/internal/apicommon/v3/pagination"\n')),(0,i.kt)("h2",{id:"index"},"Index"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#func-totalpages"},"func totalPages(totalResults int64, perPage uint16) int")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#type-link"},"type Link"),(0,i.kt)("ul",{parentName:"li"},(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#func-getresourcepathlink"},"func GetResourcePathLink(resourcePath string) Link")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#func-getresourcepathlinkwithmethod"},"func GetResourcePathLinkWithMethod(resourcePath string, method string) Link")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#func-nextlink"},"func nextLink(totalResults int64, paginationParams Params, resourcePath string) *Link")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#func-previouslink"},"func previousLink(totalResults int64, paginationParams Params, resourcePath string) *Link")))),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#type-pagination"},"type Pagination"),(0,i.kt)("ul",{parentName:"li"},(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#func-newpagination"},"func NewPagination(totalResults int64, paginationParams Params, resourcePath string) *Pagination")))),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#type-params"},"type Params"),(0,i.kt)("ul",{parentName:"li"},(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#func-default"},"func Default() Params"))))),(0,i.kt)("h2",{id:"func-totalpages"},"func totalPages"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},"func totalPages(totalResults int64, perPage uint16) int\n")),(0,i.kt)("h2",{id:"type-link"},"type Link"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},'type Link struct {\n    Href   string `json:"href"`\n    Method string `json:"method,omitempty"`\n}\n')),(0,i.kt)("h3",{id:"func-getresourcepathlink"},"func GetResourcePathLink"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},"func GetResourcePathLink(resourcePath string) Link\n")),(0,i.kt)("h3",{id:"func-getresourcepathlinkwithmethod"},"func GetResourcePathLinkWithMethod"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},"func GetResourcePathLinkWithMethod(resourcePath string, method string) Link\n")),(0,i.kt)("h3",{id:"func-nextlink"},"func nextLink"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},"func nextLink(totalResults int64, paginationParams Params, resourcePath string) *Link\n")),(0,i.kt)("h3",{id:"func-previouslink"},"func previousLink"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},"func previousLink(totalResults int64, paginationParams Params, resourcePath string) *Link\n")),(0,i.kt)("h2",{id:"type-pagination"},"type Pagination"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},'type Pagination struct {\n    TotalResults int64 `json:"total_results"`\n    TotalPages   int   `json:"total_pages"`\n    First        Link  `json:"first"`\n    Last         Link  `json:"last"`\n    Next         *Link `json:"next"`\n    Previous     *Link `json:"previous"`\n}\n')),(0,i.kt)("h3",{id:"func-newpagination"},"func NewPagination"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},"func NewPagination(totalResults int64, paginationParams Params, resourcePath string) *Pagination\n")),(0,i.kt)("h2",{id:"type-params"},"type Params"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},'type Params struct {\n    Page    int    `query:"page" validate:"gte=1"`\n    PerPage uint16 `query:"per_page" validate:"gte=1,lte=5000"`\n}\n')),(0,i.kt)("h3",{id:"func-default"},"func Default"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},"func Default() Params\n")),(0,i.kt)("p",null,"Generated by ",(0,i.kt)("a",{parentName:"p",href:"https://github.com/princjef/gomarkdoc"},"gomarkdoc")))}g.isMDXComponent=!0}}]);