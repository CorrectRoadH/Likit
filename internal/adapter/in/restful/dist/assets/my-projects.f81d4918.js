var g=Object.defineProperty,m=Object.defineProperties;var x=Object.getOwnPropertyDescriptors;var i=Object.getOwnPropertySymbols;var y=Object.prototype.hasOwnProperty,h=Object.prototype.propertyIsEnumerable;var p=(a,t,e)=>t in a?g(a,t,{enumerable:!0,configurable:!0,writable:!0,value:e}):a[t]=e,l=(a,t)=>{for(var e in t||(t={}))y.call(t,e)&&p(a,e,t[e]);if(i)for(var e of i(t))h.call(t,e)&&p(a,e,t[e]);return a},d=(a,t)=>m(a,x(t));import{r as o,j as n,ak as b,a6 as w}from"./vendor.9dccb7a3.js";import C from"./project.b56bc9a4.js";/* empty css               *//* empty css              */function v(){const[a,t]=o.exports.useState(new Array(6).fill({})),[e,c]=o.exports.useState(!0),{Row:j,Col:u}=b,f=async()=>{c(!0);const{data:r}=await w.get("/api/user/projectList").finally(()=>{c(!1)});t(r)};return o.exports.useEffect(()=>{f()},[]),n(j,{gutter:12,children:a.map((r,s)=>n(u,{span:8,style:s>a.length-4&&s<a.length?{marginTop:"16px"}:{},children:n(C,d(l({},r),{loading:e}))},s))})}export{v as default};
