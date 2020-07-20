(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-fdfc453e"],{"135f":function(t,e,n){},"1c0b":function(t,e,n){"use strict";n.d(e,"a",(function(){return r})),n.d(e,"b",(function(){return o}));var a=n("b775");function r(t){return Object(a["a"])({url:"/createRealEstate",method:"post",data:t})}function o(t){return Object(a["a"])({url:"/queryRealEstateList",method:"post",data:t})}},"49fc":function(t,e,n){"use strict";var a=n("135f"),r=n.n(a);r.a},"8d49":function(t,e,n){"use strict";n.d(e,"b",(function(){return r})),n.d(e,"c",(function(){return o})),n.d(e,"d",(function(){return i})),n.d(e,"a",(function(){return l}));var a=n("b775");function r(t){return Object(a["a"])({url:"/queryDonatingList",method:"post",data:t})}function o(t){return Object(a["a"])({url:"/queryDonatingListByGrantee",method:"post",data:t})}function i(t){return Object(a["a"])({url:"/updateDonating",method:"post",data:t})}function l(t){return Object(a["a"])({url:"/createDonating",method:"post",data:t})}},"945e":function(t,e,n){"use strict";n.d(e,"c",(function(){return r})),n.d(e,"d",(function(){return o})),n.d(e,"b",(function(){return i})),n.d(e,"e",(function(){return l})),n.d(e,"a",(function(){return c}));var a=n("b775");function r(t){return Object(a["a"])({url:"/querySellingList",method:"post",data:t})}function o(t){return Object(a["a"])({url:"/querySellingListByBuyer",method:"post",data:t})}function i(t){return Object(a["a"])({url:"/createSellingByBuy",method:"post",data:t})}function l(t){return Object(a["a"])({url:"/updateSelling",method:"post",data:t})}function c(t){return Object(a["a"])({url:"/createSelling",method:"post",data:t})}},e825:function(t,e,n){"use strict";n.r(e);var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"container"},[n("el-alert",{attrs:{type:"success"}},[n("p",[t._v("账户ID: "+t._s(t.accountId))]),t._v(" "),n("p",[t._v("用户名: "+t._s(t.userName))]),t._v(" "),n("p",[t._v("余额: ￥"+t._s(t.balance)+" 元")]),t._v(" "),n("p",[t._v("当发起出售、捐赠或质押操作后，担保状态为true")]),t._v(" "),n("p",[t._v("当担保状态为false时，才可发起出售、捐赠或质押操作")])]),t._v(" "),0==t.realEstateList.length?n("div",{staticStyle:{"text-align":"center"}},[n("el-alert",{attrs:{title:"查询不到数据",type:"warning"}})],1):t._e(),t._v(" "),n("el-row",{directives:[{name:"loading",rawName:"v-loading",value:t.loading,expression:"loading"}],attrs:{gutter:20}},t._l(t.realEstateList,(function(e,a){return n("el-col",{key:a,attrs:{span:6,offset:1}},[n("el-card",{staticClass:"realEstate-card"},[n("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[t._v("\n          担保状态:\n          "),n("span",{staticStyle:{color:"rgb(255, 0, 0)"}},[t._v(t._s(e.encumbrance))])]),t._v(" "),n("div",{staticClass:"item"},[n("el-tag",[t._v("房产ID: ")]),t._v(" "),n("span",[t._v(t._s(e.realEstateId))])],1),t._v(" "),n("div",{staticClass:"item"},[n("el-tag",{attrs:{type:"success"}},[t._v("业主ID: ")]),t._v(" "),n("span",[t._v(t._s(e.proprietor))])],1),t._v(" "),n("div",{staticClass:"item"},[n("el-tag",{attrs:{type:"warning"}},[t._v("总空间: ")]),t._v(" "),n("span",[t._v(t._s(e.totalArea)+" ㎡")])],1),t._v(" "),n("div",{staticClass:"item"},[n("el-tag",{attrs:{type:"danger"}},[t._v("居住空间: ")]),t._v(" "),n("span",[t._v(t._s(e.livingSpace)+" ㎡")])],1),t._v(" "),e.encumbrance||"admin"===t.roles[0]?t._e():n("div",[n("el-button",{attrs:{type:"text"},on:{click:function(n){return t.openDialog(e)}}},[t._v("出售")]),t._v(" "),n("el-divider",{attrs:{direction:"vertical"}}),t._v(" "),n("el-button",{attrs:{type:"text"},on:{click:function(n){return t.openDonatingDialog(e)}}},[t._v("捐赠")])],1),t._v(" "),"admin"===t.roles[0]?n("el-rate"):t._e()],1)],1)})),1),t._v(" "),n("el-dialog",{directives:[{name:"loading",rawName:"v-loading",value:t.loadingDialog,expression:"loadingDialog"}],attrs:{visible:t.dialogCreateSelling,"close-on-click-modal":!1},on:{"update:visible":function(e){t.dialogCreateSelling=e},close:function(e){return t.resetForm("realForm")}}},[n("el-form",{ref:"realForm",attrs:{model:t.realForm,rules:t.rules,"label-width":"100px"}},[n("el-form-item",{attrs:{label:"价格 (元)",prop:"price"}},[n("el-input-number",{attrs:{precision:2,step:1e4,min:0},model:{value:t.realForm.price,callback:function(e){t.$set(t.realForm,"price",e)},expression:"realForm.price"}})],1),t._v(" "),n("el-form-item",{attrs:{label:"有效期 (天)",prop:"salePeriod"}},[n("el-input-number",{attrs:{min:1},model:{value:t.realForm.salePeriod,callback:function(e){t.$set(t.realForm,"salePeriod",e)},expression:"realForm.salePeriod"}})],1)],1),t._v(" "),n("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[n("el-button",{attrs:{type:"primary"},on:{click:function(e){return t.createSelling("realForm")}}},[t._v("立即出售")]),t._v(" "),n("el-button",{on:{click:function(e){t.dialogCreateSelling=!1}}},[t._v("取 消")])],1)],1),t._v(" "),n("el-dialog",{directives:[{name:"loading",rawName:"v-loading",value:t.loadingDialog,expression:"loadingDialog"}],attrs:{visible:t.dialogCreateDonating,"close-on-click-modal":!1},on:{"update:visible":function(e){t.dialogCreateDonating=e},close:function(e){return t.resetForm("DonatingForm")}}},[n("el-form",{ref:"DonatingForm",attrs:{model:t.DonatingForm,rules:t.rulesDonating,"label-width":"100px"}},[n("el-form-item",{attrs:{label:"业主",prop:"proprietor"}},[n("el-select",{attrs:{placeholder:"请选择业主"},on:{change:t.selectGet},model:{value:t.DonatingForm.proprietor,callback:function(e){t.$set(t.DonatingForm,"proprietor",e)},expression:"DonatingForm.proprietor"}},t._l(t.accountList,(function(e){return n("el-option",{key:e.accountId,attrs:{label:e.userName,value:e.accountId}},[n("span",{staticStyle:{float:"left"}},[t._v(t._s(e.userName))]),t._v(" "),n("span",{staticStyle:{float:"right",color:"#8492a6","font-size":"13px"}},[t._v(t._s(e.accountId))])])})),1)],1)],1),t._v(" "),n("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[n("el-button",{attrs:{type:"primary"},on:{click:function(e){return t.createDonating("DonatingForm")}}},[t._v("立即捐赠")]),t._v(" "),n("el-button",{on:{click:function(e){t.dialogCreateDonating=!1}}},[t._v("取 消")])],1)],1)],1)},r=[],o=(n("8e6e"),n("ac6a"),n("456d"),n("bd86")),i=n("2f62"),l=n("5723"),c=n("1c0b"),s=n("945e"),u=n("8d49");function d(t,e){var n=Object.keys(t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(t);e&&(a=a.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),n.push.apply(n,a)}return n}function g(t){for(var e=1;e<arguments.length;e++){var n=null!=arguments[e]?arguments[e]:{};e%2?d(Object(n),!0).forEach((function(e){Object(o["a"])(t,e,n[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(n)):d(Object(n)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(n,e))}))}return t}var f={name:"RealeState",data:function(){var t=function(t,e,n){e<=0?n(new Error("必须大于0")):n()};return{loading:!0,loadingDialog:!1,realEstateList:[],dialogCreateSelling:!1,dialogCreateDonating:!1,realForm:{price:0,salePeriod:0},rules:{price:[{validator:t,trigger:"blur"}],salePeriod:[{validator:t,trigger:"blur"}]},DonatingForm:{proprietor:""},rulesDonating:{proprietor:[{required:!0,message:"请选择业主",trigger:"change"}]},accountList:[],valItem:{}}},computed:g({},Object(i["b"])(["accountId","roles","userName","balance"])),created:function(){var t=this;"admin"===this.roles[0]?Object(c["b"])().then((function(e){null!==e&&(t.realEstateList=e),t.loading=!1})).catch((function(e){t.loading=!1})):Object(c["b"])({proprietor:this.accountId}).then((function(e){null!==e&&(t.realEstateList=e),t.loading=!1})).catch((function(e){t.loading=!1}))},methods:{openDialog:function(t){this.dialogCreateSelling=!0,this.valItem=t},openDonatingDialog:function(t){var e=this;this.dialogCreateDonating=!0,this.valItem=t,Object(l["b"])().then((function(t){null!==t&&(e.accountList=t.filter((function(t){return"管理员"!==t.userName&&t.accountId!==e.accountId})))}))},createSelling:function(t){var e=this;this.$refs[t].validate((function(t){if(!t)return!1;e.$confirm("是否立即出售?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"success"}).then((function(){e.loadingDialog=!0,Object(s["a"])({objectOfSale:e.valItem.realEstateId,seller:e.valItem.proprietor,price:e.realForm.price,salePeriod:e.realForm.salePeriod}).then((function(t){e.loadingDialog=!1,e.dialogCreateSelling=!1,null!==t?e.$message({type:"success",message:"出售成功!"}):e.$message({type:"error",message:"出售失败!"}),setTimeout((function(){window.location.reload()}),1e3)})).catch((function(t){e.loadingDialog=!1,e.dialogCreateSelling=!1}))})).catch((function(){e.loadingDialog=!1,e.dialogCreateSelling=!1,e.$message({type:"info",message:"已取消出售"})}))}))},createDonating:function(t){var e=this;this.$refs[t].validate((function(t){if(!t)return!1;e.$confirm("是否立即捐赠?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"success"}).then((function(){e.loadingDialog=!0,Object(u["a"])({objectOfDonating:e.valItem.realEstateId,donor:e.valItem.proprietor,grantee:e.DonatingForm.proprietor}).then((function(t){e.loadingDialog=!1,e.dialogCreateDonating=!1,null!==t?e.$message({type:"success",message:"捐赠成功!"}):e.$message({type:"error",message:"捐赠失败!"}),setTimeout((function(){window.location.reload()}),1e3)})).catch((function(t){e.loadingDialog=!1,e.dialogCreateDonating=!1}))})).catch((function(){e.loadingDialog=!1,e.dialogCreateDonating=!1,e.$message({type:"info",message:"已取消捐赠"})}))}))},resetForm:function(t){this.$refs[t].resetFields()},selectGet:function(t){this.DonatingForm.proprietor=t}}},p=f,m=(n("49fc"),n("2877")),v=Object(m["a"])(p,a,r,!1,null,null,null);e["default"]=v.exports}}]);