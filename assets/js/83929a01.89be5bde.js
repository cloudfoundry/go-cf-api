"use strict";(self.webpackChunkmy_website=self.webpackChunkmy_website||[]).push([[989],{3905:function(e,t,r){r.d(t,{Zo:function(){return u},kt:function(){return m}});var n=r(7294);function a(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function i(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function o(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?i(Object(r),!0).forEach((function(t){a(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):i(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function c(e,t){if(null==e)return{};var r,n,a=function(e,t){if(null==e)return{};var r,n,a={},i=Object.keys(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||(a[r]=e[r]);return a}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}var s=n.createContext({}),l=function(e){var t=n.useContext(s),r=t;return e&&(r="function"==typeof e?e(t):o(o({},t),e)),r},u=function(e){var t=l(e.components);return n.createElement(s.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},d=n.forwardRef((function(e,t){var r=e.components,a=e.mdxType,i=e.originalType,s=e.parentName,u=c(e,["components","mdxType","originalType","parentName"]),d=l(r),m=a,h=d["".concat(s,".").concat(m)]||d[m]||p[m]||i;return r?n.createElement(h,o(o({ref:t},u),{},{components:r})):n.createElement(h,o({ref:t},u))}));function m(e,t){var r=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var i=r.length,o=new Array(i);o[0]=d;var c={};for(var s in t)hasOwnProperty.call(t,s)&&(c[s]=t[s]);c.originalType=e,c.mdxType="string"==typeof e?e:a,o[1]=c;for(var l=2;l<i;l++)o[l]=r[l];return n.createElement.apply(null,o)}return n.createElement.apply(null,r)}d.displayName="MDXCreateElement"},9785:function(e,t,r){r.r(t),r.d(t,{frontMatter:function(){return c},contentTitle:function(){return s},metadata:function(){return l},assets:function(){return u},toc:function(){return p},default:function(){return m}});var n=r(7462),a=r(3366),i=(r(7294),r(3905)),o=["components"],c={slug:"ADR - ADR Format",title:"Use Markdown Architectural Decision Records",authors:[{name:"Andrew Paine",title:"go-cf-api Team",url:"https://github.com/andy-paine",image_url:"https://avatars1.githubusercontent.com/u/14118619?v=4"}],tags:["adr"]},s="Use Markdown Architectural Decision Records",l={permalink:"/go-cf-api/adrs/ADR - ADR Format",editUrl:"https://github.com/cloudfoundry/go-cf-api/edit/main/docs/adrs/2021-07-25-use-markdown-architectural-decision-records.md",source:"@site/adrs/2021-07-25-use-markdown-architectural-decision-records.md",title:"Use Markdown Architectural Decision Records",description:"* Status: accepted",date:"2021-07-25T00:00:00.000Z",formattedDate:"July 25, 2021",tags:[{label:"adr",permalink:"/go-cf-api/adrs/tags/adr"}],readingTime:.8,truncated:!1,authors:[{name:"Andrew Paine",title:"go-cf-api Team",url:"https://github.com/andy-paine",image_url:"https://avatars1.githubusercontent.com/u/14118619?v=4",imageURL:"https://avatars1.githubusercontent.com/u/14118619?v=4"}],prevItem:{title:"Use sqlboiler together with build tags to support multiple databases",permalink:"/go-cf-api/adrs/ADR - SQL Framework"}},u={authorsImageUrls:[void 0]},p=[{value:"Context and Problem Statement",id:"context-and-problem-statement",children:[]},{value:"Considered Options",id:"considered-options",children:[]},{value:"Decision Outcome",id:"decision-outcome",children:[]}],d={toc:p};function m(e){var t=e.components,r=(0,a.Z)(e,o);return(0,i.kt)("wrapper",(0,n.Z)({},d,r,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"Status: accepted"),(0,i.kt)("li",{parentName:"ul"},"Deciders: ",(0,i.kt)("a",{parentName:"li",href:"https://github.com/andy-paine"},"Andy Paine")),(0,i.kt)("li",{parentName:"ul"},"Date: 2021-07-28")),(0,i.kt)("h2",{id:"context-and-problem-statement"},"Context and Problem Statement"),(0,i.kt)("p",null,"We want to record architectural decisions made in this project.\nWhich format and structure should these records follow?"),(0,i.kt)("h2",{id:"considered-options"},"Considered Options"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"https://adr.github.io/madr/"},"MADR")," 2.1.2 \u2013 The Markdown Architectural Decision Records"),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"http://thinkrelevance.com/blog/2011/11/15/documenting-architecture-decisions"},"Michael Nygard's template"),' \u2013 The first incarnation of the term "ADR"'),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"https://www.infoq.com/articles/sustainable-architectural-design-decisions"},"Sustainable Architectural Decisions")," \u2013 The Y-Statements"),(0,i.kt)("li",{parentName:"ul"},"Other templates listed at ",(0,i.kt)("a",{parentName:"li",href:"https://github.com/joelparkerhenderson/architecture_decision_record"},"https://github.com/joelparkerhenderson/architecture_decision_record")),(0,i.kt)("li",{parentName:"ul"},"Formless \u2013 No conventions for file format and structure")),(0,i.kt)("h2",{id:"decision-outcome"},"Decision Outcome"),(0,i.kt)("p",null,'Chosen option: "MADR 2.1.2", because'),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"Implicit assumptions should be made explicit.\nDesign documentation is important to enable people understanding the decisions later on.\nSee also ",(0,i.kt)("a",{parentName:"li",href:"https://doi.org/10.1109/TSE.1986.6312940"},"A rational design process: How and why to fake it"),"."),(0,i.kt)("li",{parentName:"ul"},"The MADR format is lean and fits our development style."),(0,i.kt)("li",{parentName:"ul"},"The MADR structure is comprehensible and facilitates usage & maintenance."),(0,i.kt)("li",{parentName:"ul"},"The MADR project is vivid."),(0,i.kt)("li",{parentName:"ul"},"Version 2.1.2 is the latest one available when starting to document ADRs.")))}m.isMDXComponent=!0}}]);