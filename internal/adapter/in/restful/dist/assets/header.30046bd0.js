import{r as d,b as s,j as e,b2 as p,C as f,a_ as h,aQ as x,av as _,ai as g,a5 as u,al as v}from"./vendor.9dccb7a3.js";/* empty css               *//* empty css               */import{u as S}from"./index.3df1e727.js";import{l as N}from"./index.69444b8c.js";var i={"info-wrapper":"_info-wrapper_g8kyo_1","info-avatar":"_info-avatar_g8kyo_4","info-content":"_info-content_g8kyo_14","verified-tag":"_verified-tag_g8kyo_20","edit-btn":"_edit-btn_g8kyo_25"};function C({userInfo:t={},loading:n}){const a=S(N),[l,o]=d.exports.useState("");function m(j,c){o(c.originFile?URL.createObjectURL(c.originFile):"")}d.exports.useEffect(()=>{o(t.avatar)},[t]);const b=e(v,{text:{rows:0},style:{width:"100px",height:"100px"},animation:!0}),r=e(v,{text:{rows:1},animation:!0});return s("div",{className:i["info-wrapper"],children:[e(p,{showUploadList:!1,onChange:m,children:n?b:e(f,{size:100,triggerIcon:e(h,{}),className:i["info-avatar"],children:l?e("img",{src:l}):e(x,{})})}),e(_,{className:i["info-content"],column:2,colon:"\uFF1A",labelStyle:{textAlign:"right"},data:[{label:a["userSetting.label.name"],value:n?r:t.name},{label:a["userSetting.label.verified"],value:n?r:s("span",{children:[t.verified?e(g,{color:"green",className:i["verified-tag"],children:a["userSetting.value.verified"]}):e(g,{color:"red",className:i["verified-tag"],children:a["userSetting.value.notVerified"]}),e(u,{role:"button",className:i["edit-btn"],children:a["userSetting.btn.edit"]})]})},{label:a["userSetting.label.accountId"],value:n?r:t.accountId},{label:a["userSetting.label.phoneNumber"],value:n?r:s("span",{children:[t.phoneNumber,e(u,{role:"button",className:i["edit-btn"],children:a["userSetting.btn.edit"]})]})},{label:a["userSetting.label.registrationTime"],value:n?r:t.registrationTime}]})]})}export{C as default};
