import{r as e,j as t,b1 as s,al as c,C as l,f as x,a6 as m}from"./vendor.9dccb7a3.js";/* empty css               *//* empty css               *//* empty css               */const{Text:u}=x;function j(){const[r,i]=e.exports.useState(new Array(4).fill({})),[o,p]=e.exports.useState(!0),d=async()=>{const{data:a}=await m.get("/api/users/teamList").finally(()=>p(!1));i(a)};return e.exports.useEffect(()=>{d()},[]),t(s,{dataSource:r,render:(a,n)=>t(s.Item,{style:n!==r.length-1?{padding:"8px 0px"}:{padding:"8px 0px 0px 0px"},children:o?t(c,{animation:!0,text:{width:["80%","20%"],rows:2},image:{shape:"circle"}}):t(s.Item.Meta,{avatar:t(l,{size:48,children:t("img",{src:a.avatar})}),title:a.name,description:t(u,{type:"secondary",style:{fontSize:"12px"},children:`\u5171${(a.members||0).toLocaleString()}\u4EBA`})})},n)})}export{j as default};
