"use strict";(self.webpackChunkmy_website=self.webpackChunkmy_website||[]).push([[6445],{3905:function(e,t,n){n.d(t,{Zo:function(){return u},kt:function(){return m}});var o=n(7294);function i(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function a(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);t&&(o=o.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,o)}return n}function r(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?a(Object(n),!0).forEach((function(t){i(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):a(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,o,i=function(e,t){if(null==e)return{};var n,o,i={},a=Object.keys(e);for(o=0;o<a.length;o++)n=a[o],t.indexOf(n)>=0||(i[n]=e[n]);return i}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(o=0;o<a.length;o++)n=a[o],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(i[n]=e[n])}return i}var s=o.createContext({}),p=function(e){var t=o.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):r(r({},t),e)),n},u=function(e){var t=p(e.components);return o.createElement(s.Provider,{value:t},e.children)},c={inlineCode:"code",wrapper:function(e){var t=e.children;return o.createElement(o.Fragment,{},t)}},d=o.forwardRef((function(e,t){var n=e.components,i=e.mdxType,a=e.originalType,s=e.parentName,u=l(e,["components","mdxType","originalType","parentName"]),d=p(n),m=i,h=d["".concat(s,".").concat(m)]||d[m]||c[m]||a;return n?o.createElement(h,r(r({ref:t},u),{},{components:n})):o.createElement(h,r({ref:t},u))}));function m(e,t){var n=arguments,i=t&&t.mdxType;if("string"==typeof e||i){var a=n.length,r=new Array(a);r[0]=d;var l={};for(var s in t)hasOwnProperty.call(t,s)&&(l[s]=t[s]);l.originalType=e,l.mdxType="string"==typeof e?e:i,r[1]=l;for(var p=2;p<a;p++)r[p]=n[p];return o.createElement.apply(null,r)}return o.createElement.apply(null,n)}d.displayName="MDXCreateElement"},7793:function(e,t,n){n.r(t),n.d(t,{frontMatter:function(){return l},contentTitle:function(){return s},metadata:function(){return p},assets:function(){return u},toc:function(){return c},default:function(){return m}});var o=n(7462),i=n(3366),a=(n(7294),n(3905)),r=["components"],l={slug:"ADR - Traffic Splitter",title:"Deploy haproxy to route specific endpoints/methods to the new implementation",authors:[{name:"Andrew Paine",title:"go-cf-api Team",url:"https://github.com/andy-paine",image_url:"https://avatars1.githubusercontent.com/u/14118619?v=4"},{name:"Sven Krieger",title:"go-cf-api Team",url:"https://github.com/svkrieger",image_url:"https://avatars1.githubusercontent.com/u/37476281?v=4"},{name:"Philipp Thun",title:"go-cf-api Team",url:"https://github.com/philippthun",image_url:"https://avatars1.githubusercontent.com/u/618301?v=4"}],tags:["adr"]},s="Deploy [HAProxy](https://www.haproxy.org/) to route specific endpoints/methods to the new implementation",p={permalink:"/go-cf-api/adrs/ADR - Traffic Splitter",editUrl:"https://github.com/cloudfoundry/go-cf-api/edit/main/docs/adrs/2021-08-13-split-traffic-between-both-implementations.md",source:"@site/adrs/2021-08-13-split-traffic-between-both-implementations.md",title:"Deploy haproxy to route specific endpoints/methods to the new implementation",description:"* Status: accepted",date:"2021-08-13T00:00:00.000Z",formattedDate:"August 13, 2021",tags:[{label:"adr",permalink:"/go-cf-api/adrs/tags/adr"}],readingTime:2.385,truncated:!1,authors:[{name:"Andrew Paine",title:"go-cf-api Team",url:"https://github.com/andy-paine",image_url:"https://avatars1.githubusercontent.com/u/14118619?v=4",imageURL:"https://avatars1.githubusercontent.com/u/14118619?v=4"},{name:"Sven Krieger",title:"go-cf-api Team",url:"https://github.com/svkrieger",image_url:"https://avatars1.githubusercontent.com/u/37476281?v=4",imageURL:"https://avatars1.githubusercontent.com/u/37476281?v=4"},{name:"Philipp Thun",title:"go-cf-api Team",url:"https://github.com/philippthun",image_url:"https://avatars1.githubusercontent.com/u/618301?v=4",imageURL:"https://avatars1.githubusercontent.com/u/618301?v=4"}],prevItem:{title:"Use client_golang prometheus library to expose metrics",permalink:"/go-cf-api/adrs/ADR - Prometheus Framework"},nextItem:{title:"Use Zap logging as a performant and customizable structured logging framework",permalink:"/go-cf-api/adrs/ADR - Logging Framework"}},u={authorsImageUrls:[void 0,void 0,void 0]},c=[{value:"Context and Problem Statement",id:"context-and-problem-statement",children:[],level:2},{value:"Decision Drivers  optional ",id:"decision-drivers--optional-",children:[],level:2},{value:"Considered Options",id:"considered-options",children:[],level:2},{value:"Decision Outcome",id:"decision-outcome",children:[{value:"Positive Consequences  optional ",id:"positive-consequences--optional-",children:[],level:3},{value:"Negative Consequences  optional ",id:"negative-consequences--optional-",children:[],level:3}],level:2},{value:"Pros and Cons of the Options  optional ",id:"pros-and-cons-of-the-options--optional-",children:[{value:"Option 1 (Deploy new CC alongside existing CC in same instance group, using <code>nginx</code> for routing)",id:"option-1-deploy-new-cc-alongside-existing-cc-in-same-instance-group-using-nginx-for-routing",children:[],level:3},{value:"Option 2 (Complete an entire endpoint at once (all HTTP methods) and use <code>gorouter</code> for path based routing)",id:"option-2-complete-an-entire-endpoint-at-once-all-http-methods-and-use-gorouter-for-path-based-routing",children:[],level:3}],level:2},{value:"Links  optional ",id:"links--optional-",children:[],level:2}],d={toc:c};function m(e){var t=e.components,n=(0,i.Z)(e,r);return(0,a.kt)("wrapper",(0,o.Z)({},d,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"Status: accepted"),(0,a.kt)("li",{parentName:"ul"},"Deciders: ",(0,a.kt)("a",{parentName:"li",href:"https://github.com/andy-paine"},"Andy Paine"),", ",(0,a.kt)("a",{parentName:"li",href:"https://github.com/svkrieger"},"Sven Krieger"),", ",(0,a.kt)("a",{parentName:"li",href:"https://github.com/philippthun"},"Philipp Thun")),(0,a.kt)("li",{parentName:"ul"},"Date: 2021-08-04")),(0,a.kt)("h2",{id:"context-and-problem-statement"},"Context and Problem Statement"),(0,a.kt)("p",null,"The existing Cloud Controller is the reference implementation of the CF v3 API.\nThis project should avoid replacing the entire ",(0,a.kt)("inlineCode",{parentName:"p"},"cloud_controller_ng"),' project in a single "big bang" migration.\nIn order to build the new implementation iteratively, this project should be deployable in parallel with the existing implementation.\nIt should be possible to route individual API calls to the new Cloud Controller as soon as each endpoint is complete.'),(0,a.kt)("h2",{id:"decision-drivers--optional-"},"Decision Drivers "),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"Discover bugs early"),(0,a.kt)("li",{parentName:"ul"},"Deliver value from reimplementation quickly"),(0,a.kt)("li",{parentName:"ul"},"Minimise mean-time-to-recovery when bugs are discovered")),(0,a.kt)("h2",{id:"considered-options"},"Considered Options"),(0,a.kt)("ol",null,(0,a.kt)("li",{parentName:"ol"},"Deploy new CC alongside existing CC in same instance group, using ",(0,a.kt)("inlineCode",{parentName:"li"},"nginx")," for routing"),(0,a.kt)("li",{parentName:"ol"},"Complete an entire endpoint at once (all HTTP methods) and use ",(0,a.kt)("inlineCode",{parentName:"li"},"gorouter")," for path based routing"),(0,a.kt)("li",{parentName:"ol"},"Deploy a dedicated path and HTTP method based router/proxy in front of old and new implementations and split traffic based on that")),(0,a.kt)("h2",{id:"decision-outcome"},"Decision Outcome"),(0,a.kt)("p",null,"Chosen option: 3 (deploy a path and HTTP method based router/proxy in front of old and new implementations and split traffic based on that),\nbecause it is the only option that allows for separate scaling of old and new implementations as well as routing based on HTTP method and path.\nFollowing images shows a rough routing example with just the verry specific ",(0,a.kt)("inlineCode",{parentName:"p"},"GET /v3/buildpacks/:guid")," endpoint beeing routed to the go implementation. Everything else will be routed to the cloudcontroller_ng.\n",(0,a.kt)("img",{parentName:"p",src:"https://user-images.githubusercontent.com/5863788/145360535-a02ebd16-e339-461e-bff5-612b3c4c8f46.png",alt:"image"})),(0,a.kt)("h3",{id:"positive-consequences--optional-"},"Positive Consequences "),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"New implementation can be built in small units (endpoint + HTTP method)"),(0,a.kt)("li",{parentName:"ul"},"Proxy can be registered only for certain routes, minimising throughput"),(0,a.kt)("li",{parentName:"ul"},"Networking such as TLS in new implementation can be delayed until closer to completion (as proxy can perform this function)"),(0,a.kt)("li",{parentName:"ul"},"Good support for HAProxy BOSH release as it is maintained by SAP team")),(0,a.kt)("h3",{id:"negative-consequences--optional-"},"Negative Consequences "),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"Additional software to manage"),(0,a.kt)("li",{parentName:"ul"},"HAProxy BOSH release is not that well suited to this use case")),(0,a.kt)("h2",{id:"pros-and-cons-of-the-options--optional-"},"Pros and Cons of the Options "),(0,a.kt)("h3",{id:"option-1-deploy-new-cc-alongside-existing-cc-in-same-instance-group-using-nginx-for-routing"},"Option 1 (Deploy new CC alongside existing CC in same instance group, using ",(0,a.kt)("inlineCode",{parentName:"h3"},"nginx")," for routing)"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"Good, because does not change network architecture"),(0,a.kt)("li",{parentName:"ul"},"Good, because does not add any new VMs"),(0,a.kt)("li",{parentName:"ul"},"Bad, because cannot independently scale each implementation"),(0,a.kt)("li",{parentName:"ul"},"Bad, because would require changes to the CAPI release to support this use case")),(0,a.kt)("h3",{id:"option-2-complete-an-entire-endpoint-at-once-all-http-methods-and-use-gorouter-for-path-based-routing"},"Option 2 (Complete an entire endpoint at once (all HTTP methods) and use ",(0,a.kt)("inlineCode",{parentName:"h3"},"gorouter")," for path based routing)"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"Good, because does not change network architecture"),(0,a.kt)("li",{parentName:"ul"},"Good, because can independently scale each implementation"),(0,a.kt)("li",{parentName:"ul"},"Bad, because requires large amount of work to complete a whole endpoint"),(0,a.kt)("li",{parentName:"ul"},"Bad, because requires new implementation to support TLS etc. for secure communication")),(0,a.kt)("h2",{id:"links--optional-"},"Links "),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("a",{parentName:"li",href:"https://github.com/cloudfoundry-incubator/haproxy-boshrelease"},"HAProxy BOSH release"))))}m.isMDXComponent=!0}}]);