"use strict";(self.webpackChunkmy_website=self.webpackChunkmy_website||[]).push([[1997],{3905:function(e,t,n){n.d(t,{Zo:function(){return p},kt:function(){return k}});var r=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function l(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function s(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?l(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):l(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function o(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},l=Object.keys(e);for(r=0;r<l.length;r++)n=l[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(r=0;r<l.length;r++)n=l[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var i=r.createContext({}),c=function(e){var t=r.useContext(i),n=t;return e&&(n="function"==typeof e?e(t):s(s({},t),e)),n},p=function(e){var t=c(e.components);return r.createElement(i.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},d=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,l=e.originalType,i=e.parentName,p=o(e,["components","mdxType","originalType","parentName"]),d=c(n),k=a,f=d["".concat(i,".").concat(k)]||d[k]||u[k]||l;return n?r.createElement(f,s(s({ref:t},p),{},{components:n})):r.createElement(f,s({ref:t},p))}));function k(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var l=n.length,s=new Array(l);s[0]=d;var o={};for(var i in t)hasOwnProperty.call(t,i)&&(o[i]=t[i]);o.originalType=e,o.mdxType="string"==typeof e?e:a,s[1]=o;for(var c=2;c<l;c++)s[c]=n[c];return r.createElement.apply(null,s)}return r.createElement.apply(null,n)}d.displayName="MDXCreateElement"},7930:function(e,t,n){n.r(t),n.d(t,{frontMatter:function(){return o},contentTitle:function(){return i},metadata:function(){return c},toc:function(){return p},default:function(){return d}});var r=n(7462),a=n(3366),l=(n(7294),n(3905)),s=["components"],o={},i=void 0,c={unversionedId:"Packages/internal/api/v3/buildpacks",id:"Packages/internal/api/v3/buildpacks",title:"buildpacks",description:"Index",source:"@site/godocs/Packages/internal/api/v3/buildpacks.md",sourceDirName:"Packages/internal/api/v3",slug:"/Packages/internal/api/v3/buildpacks",permalink:"/go-cf-api/godocs/Packages/internal/api/v3/buildpacks",editUrl:"https://github.com/cloudfoundry/go-cf-api/edit/main/docs/godocs/Packages/internal/api/v3/buildpacks.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"v3",permalink:"/go-cf-api/godocs/Packages/internal/api/v3"},next:{title:"info",permalink:"/go-cf-api/godocs/Packages/internal/api/v3/info"}},p=[{value:"Index",id:"index",children:[],level:2},{value:"Constants",id:"constants",children:[],level:2},{value:"Variables",id:"variables",children:[],level:2},{value:"func GetBuildpackState",id:"func-getbuildpackstate",children:[],level:2},{value:"func filters",id:"func-filters",children:[],level:2},{value:"type Controller",id:"type-controller",children:[{value:"func (*Controller) Get",id:"func-controller-get",children:[],level:3},{value:"func (*Controller) List",id:"func-controller-list",children:[],level:3},{value:"func (*Controller) Post",id:"func-controller-post",children:[],level:3}],level:2},{value:"type FilterParams",id:"type-filterparams",children:[{value:"func DefaultFilters",id:"func-defaultfilters",children:[],level:3}],level:2},{value:"type ListResponse",id:"type-listresponse",children:[],level:2},{value:"type PostRequest",id:"type-postrequest",children:[{value:"func DefaultPostRequest",id:"func-defaultpostrequest",children:[],level:3}],level:2},{value:"type Presenter",id:"type-presenter",children:[{value:"func NewPresenter",id:"func-newpresenter",children:[],level:3}],level:2},{value:"type Response",id:"type-response",children:[],level:2},{value:"type presenter",id:"type-presenter-1",children:[{value:"func (*presenter) ListResponseObject",id:"func-presenter-listresponseobject",children:[],level:3},{value:"func (*presenter) ResponseObject",id:"func-presenter-responseobject",children:[],level:3}],level:2}],u={toc:p};function d(e){var t=e.components,n=(0,a.Z)(e,s);return(0,l.kt)("wrapper",(0,r.Z)({},u,n,{components:t,mdxType:"MDXLayout"}),(0,l.kt)("h1",{id:"buildpacks"},"buildpacks"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-go"},'import "github.com/cloudfoundry/go-cf-api/internal/api/v3/buildpacks"\n')),(0,l.kt)("h2",{id:"index"},"Index"),(0,l.kt)("ul",null,(0,l.kt)("li",{parentName:"ul"},(0,l.kt)("a",{parentName:"li",href:"#constants"},"Constants")),(0,l.kt)("li",{parentName:"ul"},(0,l.kt)("a",{parentName:"li",href:"#variables"},"Variables")),(0,l.kt)("li",{parentName:"ul"},(0,l.kt)("a",{parentName:"li",href:"#func-getbuildpackstate"},"func GetBuildpackState(buildpack *models.Buildpack) string")),(0,l.kt)("li",{parentName:"ul"},(0,l.kt)("a",{parentName:"li",href:"#func-filters"},"func filters(filters FilterParams) []qm.QueryMod")),(0,l.kt)("li",{parentName:"ul"},(0,l.kt)("a",{parentName:"li",href:"#type-controller"},"type Controller"),(0,l.kt)("ul",{parentName:"li"},(0,l.kt)("li",{parentName:"ul"},(0,l.kt)("a",{parentName:"li",href:"#func-controller-get"},"func (cont *Controller) Get(echoCtx echo.Context) error")),(0,l.kt)("li",{parentName:"ul"},(0,l.kt)("a",{parentName:"li",href:"#func-controller-list"},"func (cont *Controller) List(echoCtx echo.Context) error")),(0,l.kt)("li",{parentName:"ul"},(0,l.kt)("a",{parentName:"li",href:"#func-controller-post"},"func (cont *Controller) Post(echoCtx echo.Context) error")))),(0,l.kt)("li",{parentName:"ul"},(0,l.kt)("a",{parentName:"li",href:"#type-filterparams"},"type FilterParams"),(0,l.kt)("ul",{parentName:"li"},(0,l.kt)("li",{parentName:"ul"},(0,l.kt)("a",{parentName:"li",href:"#func-defaultfilters"},"func DefaultFilters() FilterParams")))),(0,l.kt)("li",{parentName:"ul"},(0,l.kt)("a",{parentName:"li",href:"#type-listresponse"},"type ListResponse")),(0,l.kt)("li",{parentName:"ul"},(0,l.kt)("a",{parentName:"li",href:"#type-postrequest"},"type PostRequest"),(0,l.kt)("ul",{parentName:"li"},(0,l.kt)("li",{parentName:"ul"},(0,l.kt)("a",{parentName:"li",href:"#func-defaultpostrequest"},"func DefaultPostRequest() PostRequest")))),(0,l.kt)("li",{parentName:"ul"},(0,l.kt)("a",{parentName:"li",href:"#type-presenter"},"type Presenter"),(0,l.kt)("ul",{parentName:"li"},(0,l.kt)("li",{parentName:"ul"},(0,l.kt)("a",{parentName:"li",href:"#func-newpresenter"},"func NewPresenter() Presenter")))),(0,l.kt)("li",{parentName:"ul"},(0,l.kt)("a",{parentName:"li",href:"#type-response"},"type Response")),(0,l.kt)("li",{parentName:"ul"},(0,l.kt)("a",{parentName:"li",href:"#type-presenter"},"type presenter"),(0,l.kt)("ul",{parentName:"li"},(0,l.kt)("li",{parentName:"ul"},(0,l.kt)("a",{parentName:"li",href:"#func-presenter-listresponseobject"},"func (p ",(0,l.kt)("em",{parentName:"a"},"presenter) ListResponseObject(buildpacks models.BuildpackSlice, totalResults int64, paginationParams pagination.Params, resourcePath string) ("),"ListResponse, error)")),(0,l.kt)("li",{parentName:"ul"},(0,l.kt)("a",{parentName:"li",href:"#func-presenter-responseobject"},"func (p ",(0,l.kt)("em",{parentName:"a"},"presenter) ResponseObject(buildpack "),"models.Buildpack, resourcePath string) (*Response, error)"))))),(0,l.kt)("h2",{id:"constants"},"Constants"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-go"},'const (\n    StateAwaitingUpload = "AWAITING_UPLOAD"\n    StateReady          = "READY"\n)\n')),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-go"},'const GUIDParam = "guid"\n')),(0,l.kt)("h2",{id:"variables"},"Variables"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-go"},"var (\n    buildpackQuerier                           = func(qm ...qm.QueryMod) models.BuildpackFinisher { return models.Buildpacks(qm...) }\n    buildpackInserter models.BuildpackInserter = models.Buildpacks()\n    buildpackUpdater  models.BuildpackUpdater  = models.Buildpacks()\n)\n")),(0,l.kt)("h2",{id:"func-getbuildpackstate"},"func GetBuildpackState"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-go"},"func GetBuildpackState(buildpack *models.Buildpack) string\n")),(0,l.kt)("h2",{id:"func-filters"},"func filters"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-go"},"func filters(filters FilterParams) []qm.QueryMod\n")),(0,l.kt)("h2",{id:"type-controller"},"type Controller"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-go"},"type Controller struct {\n    DB                  *sql.DB\n    Presenter           Presenter\n    LabelSelectorParser metadata.LabelSelectorParser\n}\n")),(0,l.kt)("h3",{id:"func-controller-get"},"func ","(","*","Controller",")"," Get"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-go"},"func (cont *Controller) Get(echoCtx echo.Context) error\n")),(0,l.kt)("p",null,"Get godoc @Summary Show a buildpack @Description Retrieve all buildpacks the user has access to",".",' @Tags Buildpacks @Param guid path string true "Buildpack GUID" @Success 200 ',"{","object","}"," Response @Success 404 ","{","object","}"," interface","{","}"," @Failure 400 ","{","object","}"," v3",".","CfAPIError @Failure 500 ","{","object","}"," v3",".","CfAPIError @Router /buildpacks/","{","guid","}"," ","[","get","]"),(0,l.kt)("h3",{id:"func-controller-list"},"func ","(","*","Controller",")"," List"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-go"},"func (cont *Controller) List(echoCtx echo.Context) error\n")),(0,l.kt)("h3",{id:"func-controller-post"},"func ","(","*","Controller",")"," Post"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-go"},"func (cont *Controller) Post(echoCtx echo.Context) error\n")),(0,l.kt)("p",null,"PostBuildpack godoc @Summary Create a buildpack @Description Create a new buildpack @Tags Buildpacks @accept json @produce json @Success 201 ","{","object","}"," Response @Success 404 ","{","object","}"," interface","{","}"," @Failure 400 ","{","object","}"," v3",".","CfAPIError @Failure 500 ","{","object","}"," v3",".","CfAPIError @Router /buildpacks ","[","post","]"),(0,l.kt)("h2",{id:"type-filterparams"},"type FilterParams"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-go"},'type FilterParams struct {\n    Names         *string `query:"names"`\n    Stacks        *string `query:"stacks"`\n    OrderBy       string  `query:"order_by" validate:"oneof=created_at -created_at updated_at -updated_at position -position"`\n    LabelSelector string  `query:"label_selector"`\n}\n')),(0,l.kt)("h3",{id:"func-defaultfilters"},"func DefaultFilters"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-go"},"func DefaultFilters() FilterParams\n")),(0,l.kt)("h2",{id:"type-listresponse"},"type ListResponse"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-go"},'type ListResponse struct {\n    Pagination *pagination.Pagination `json:"pagination"`\n    Resources  []*Response            `json:"resources"`\n}\n')),(0,l.kt)("h2",{id:"type-postrequest"},"type PostRequest"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-go"},'type PostRequest struct {\n    Name     string  `json:"name" validate:"required,min=1,max=250"`\n    Stack    *string `json:"stack" validate:"omitempty,max=250"`\n    Position int     `json:"position" validate:"gte=1"`\n    Enabled  bool    `json:"enabled"`\n    Locked   bool    `json:"locked"`\n}\n')),(0,l.kt)("h3",{id:"func-defaultpostrequest"},"func DefaultPostRequest"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-go"},"func DefaultPostRequest() PostRequest\n")),(0,l.kt)("h2",{id:"type-presenter"},"type Presenter"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-go"},"type Presenter interface {\n    ResponseObject(buildpack *models.Buildpack, resourcePath string) (*Response, error)\n    ListResponseObject(\n        buildpacks models.BuildpackSlice,\n        totalResults int64,\n        paginationParams pagination.Params,\n        resourcePath string) (*ListResponse, error)\n}\n")),(0,l.kt)("h3",{id:"func-newpresenter"},"func NewPresenter"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-go"},"func NewPresenter() Presenter\n")),(0,l.kt)("h2",{id:"type-response"},"type Response"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-go"},'type Response struct {\n    GUID      string            `json:"guid"`\n    CreatedAt string            `json:"created_at"`\n    UpdatedAt string            `json:"updated_at"`\n    Name      string            `json:"name"`\n    Stack     null.String       `json:"stack"`\n    State     string            `json:"state"`\n    Filename  null.String       `json:"filename"`\n    Position  int               `json:"position"`\n    Enabled   null.Bool         `json:"enabled"`\n    Locked    null.Bool         `json:"locked"`\n    Metadata  metadata.Metadata `json:"metadata"`\n    Links     struct {\n        Self   pagination.Link `json:"self"`\n        Upload pagination.Link `json:"upload"`\n    }   `json:"links"`\n}\n')),(0,l.kt)("h2",{id:"type-presenter-1"},"type presenter"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-go"},"type presenter struct{}\n")),(0,l.kt)("h3",{id:"func-presenter-listresponseobject"},"func ","(","*","presenter",")"," ListResponseObject"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-go"},"func (p *presenter) ListResponseObject(buildpacks models.BuildpackSlice, totalResults int64, paginationParams pagination.Params, resourcePath string) (*ListResponse, error)\n")),(0,l.kt)("h3",{id:"func-presenter-responseobject"},"func ","(","*","presenter",")"," ResponseObject"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-go"},"func (p *presenter) ResponseObject(buildpack *models.Buildpack, resourcePath string) (*Response, error)\n")),(0,l.kt)("p",null,"Generated by ",(0,l.kt)("a",{parentName:"p",href:"https://github.com/princjef/gomarkdoc"},"gomarkdoc")))}d.isMDXComponent=!0}}]);