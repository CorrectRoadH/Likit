import{r as i,b as t,j as e,f as u,a3 as n,t as o,a1 as r,B as c,am as x,J as g,an as y,a6 as j}from"./vendor.9dccb7a3.js";import{u as L}from"./index.3df1e727.js";import{l as S}from"./index.662e0e1f.js";import b from"./index.6bbb456e.js";import{s}from"./index.module.7b1a989f.js";/* empty css               */import"./item.dd4083a1.js";function z(){const a=L(S),[d,h]=i.exports.useState([]),[p,l]=i.exports.useState(!1);function m(){l(!0),j.get("/api/chatList").then(f=>{h(f.data||[])}).finally(()=>{l(!1)})}return i.exports.useEffect(()=>{m()},[]),t("div",{className:s["chat-panel"],children:[t("div",{className:s["chat-panel-header"],children:[e(u.Title,{style:{marginTop:0,marginBottom:16},heading:6,children:a["monitor.title.chatPanel"]}),t(n,{size:8,children:[e(o,{style:{width:80},defaultValue:"all",children:e(o.Option,{value:"all",children:a["monitor.chat.options.all"]})}),e(r.Search,{placeholder:a["monitor.chat.placeholder.searchCategory"]}),e(c,{type:"text",iconOnly:!0,children:e(x,{})})]})]}),e("div",{className:s["chat-panel-content"],children:e(g,{loading:p,style:{width:"100%"},children:e(b,{data:d})})}),e("div",{className:s["chat-panel-footer"],children:t(n,{size:8,children:[e(r,{suffix:e(y,{})}),e(c,{type:"primary",children:a["monitor.chat.update"]})]})})]})}export{z as default};
