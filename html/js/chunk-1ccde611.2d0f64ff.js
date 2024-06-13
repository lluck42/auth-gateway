(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-1ccde611"],{"0cab":function(e,t,a){},b5e4:function(e,t,a){"use strict";a("0cab")},c91e:function(e,t,a){"use strict";a.r(t);var r=function(){var e=this,t=e._self._c;return t("div",{staticClass:"content"},[t("el-container",{staticStyle:{height:"89vh","overflow-y":"auto"}},[t("el-main",[t("el-form",{staticClass:"el-form",attrs:{inline:!0}},[t("el-form-item",[t("el-input",{staticClass:"input-with-select",attrs:{placeholder:"帐号",clearable:""},on:{clear:e.loadData},model:{value:e.formSearch.Name,callback:function(t){e.$set(e.formSearch,"Name",t)},expression:"formSearch.Name"}},[t("el-button",{attrs:{slot:"append",icon:"el-icon-search"},on:{click:function(t){return e.loadData()}},slot:"append"})],1)],1),t("el-form-item",[t("el-button",{attrs:{type:"primary",icon:"el-icon-plus"},on:{click:function(t){return e.addUser()}}},[e._v("创建管理员")])],1)],1),t("el-table",{ref:"filterTable",staticStyle:{width:"100%","margin-bottom":"20px"},attrs:{data:e.tableData,border:""}},[t("el-table-column",{attrs:{prop:"Name",align:"center","show-overflow-tooltip":!0,label:"姓名","min-width":"100"}}),t("el-table-column",{attrs:{prop:"Mobile",align:"center","show-overflow-tooltip":!0,label:"账号","min-width":"100"}}),t("el-table-column",{attrs:{prop:"RoleID",align:"center",label:"角色",formatter:e.classFormatter}}),t("el-table-column",{attrs:{prop:"CreatedAt",align:"center",label:"创建时间","min-width":"100"}}),t("el-table-column",{staticClass:"handDo",attrs:{prop:"name",align:"center","min-width":"160",fixed:"right"},scopedSlots:e._u([{key:"default",fn:function(a){return[t("el-button",{attrs:{type:"primary",size:"mini"},on:{click:function(t){return e.editUser(a.row)}}},[e._v("编辑")]),t("el-button",{staticClass:"dangerColor",attrs:{size:"mini",type:"danger"},on:{click:function(t){return t.stopPropagation(),e.delPage(a.row.ID)}}},[e._v("删除")])]}}])})],1),t("el-pagination",{staticStyle:{margin:"10px 0","text-align":"right"},attrs:{"hide-on-single-page":!1,"current-page":e.page.page,"page-sizes":[10,50,200,500],"page-size":e.page.Limit,layout:"total, prev, pager, next, jumper",total:e.page.totle},on:{"size-change":e.handleSizeChange,"current-change":e.handleCurrentChange}})],1)],1),t("el-dialog",{staticClass:"admin-model",attrs:{visible:e.showAddUser,width:"800px"},on:{"update:visible":function(t){e.showAddUser=t},close:e.resetVerForm}},[t("el-form",{ref:"addUserForm",staticClass:"demo-ruleForm",attrs:{model:e.addUserForm,"status-icon":"",rules:e.rules,"label-width":"90px"}},[t("el-form-item",{attrs:{label:"账号",prop:"Mobile"}},[t("el-input",{attrs:{disabled:0!=e.addUserForm.ID},model:{value:e.addUserForm.Mobile,callback:function(t){e.$set(e.addUserForm,"Mobile",t)},expression:"addUserForm.Mobile"}})],1),t("el-form-item",{attrs:{label:"姓名",prop:"Name"}},[t("el-input",{model:{value:e.addUserForm.Name,callback:function(t){e.$set(e.addUserForm,"Name",t)},expression:"addUserForm.Name"}})],1),t("el-form-item",{attrs:{label:"角色",prop:"RoleID"}},[t("el-select",{staticStyle:{width:"100%"},on:{change:e.selectClassify},model:{value:e.addUserForm.RoleID,callback:function(t){e.$set(e.addUserForm,"RoleID",t)},expression:"addUserForm.RoleID"}},e._l(e.roleLists,(function(e,a){return t("el-option",{key:a,attrs:{label:e.Name,value:e.ID}})})),1)],1),0==e.addUserForm.ID?t("el-form-item",{attrs:{label:"密码",prop:"Password"}},[t("el-input",{attrs:{placeholder:"密码必须包含大小写字母和数字，字符8-20位之间"},model:{value:e.addUserForm.Password,callback:function(t){e.$set(e.addUserForm,"Password",t)},expression:"addUserForm.Password"}})],1):t("el-form-item",{attrs:{label:"密码",prop:"PasswordEdit"}},[t("el-input",{model:{value:e.addUserForm.PasswordEdit,callback:function(t){e.$set(e.addUserForm,"PasswordEdit",t)},expression:"addUserForm.PasswordEdit"}})],1)],1),t("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[t("el-button",{on:{click:function(t){return e.resetVerForm()}}},[e._v("取 消")]),t("el-button",{attrs:{type:"primary"},on:{click:function(t){return e.commitAddUser()}}},[e._v("确 定")])],1)],1)],1)},s=[],o={name:"systemAdmin",data:function(){return{tableData:[],page:{page:1,Limit:20,totle:0},showAddUser:!1,addUserForm:{ID:0,Mobile:"",Password:"",PasswordEdit:"",Name:"",RoleID:"",OrganizationId:1},roleLists:[],siteLists:[],rules:{Mobile:[{required:!0,message:"请输入账号",trigger:"change"}],Name:[{required:!0,message:"请输入姓名",trigger:"change"}],RoleID:[{required:!0,message:"请选择角色",trigger:"change"}],Password:[{required:!0,message:"请输入密码",trigger:"change"},{pattern:/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,20}$/,message:"密码必须包含大小写字母和数字，字符8-20位之间",trigger:"blur"}],PasswordEdit:[{pattern:/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,20}$/,message:"密码必须包含大小写字母和数字，字符8-20位之间",trigger:"blur"}]},formSearch:{id:"",Name:""},passwordInputType:"",sortType:"",isAddBtn:!0}},mounted:function(){this.isAddBtn=!0,this.getRoleList(),this.loadData()},methods:{clickTree:function(e){this.formSearch.id=e.ID,this.formSearch.Name="",this.page.page=1,this.isAddBtn=!1,this.resetVerForm(),this.addUserForm.OrganizationId=e.ID,this.loadData()},handleSizeChange:function(e){this.page.Limit=e,this.loadData()},handleCurrentChange:function(e){this.page.page=e,this.loadData()},loadData:function(){var e=this,t={Page:this.page.page,Limit:this.page.Limit,id:this.formSearch.id,Name:this.formSearch.Name};this.apiSever.get("".concat(this.URl.ADMIN_LIST+"?Page="+t.Page+"&Limit="+t.Limit+"&Name="+t.Name+"&OrganizationId="+t.id),(function(t){console.log(t.data.list),e.page.totle=t.data.count,e.tableData=t.data.list}))},getRoleList:function(){var e=this;this.apiSever.get("".concat(this.URl.ROLE_LIST+"?Page=1&Limit=100000"),(function(t){e.roleLists=t.data.list}))},classFormatter:function(e,t,a){for(var r=0;r<this.roleLists.length;r++){var s=this.roleLists[r];if(e.RoleID==s.ID)return s.Name}},addUser:function(){this.showAddUser=!0},selectClassify:function(e){console.log(e),this.addUserForm.RoleID=e},commitAddUser:function(){var e=this;this.$refs["addUserForm"].validate((function(t){if(!t)return!1;var a={ID:e.addUserForm.ID,Name:e.addUserForm.Name,Mobile:e.addUserForm.Mobile,Password:0==e.addUserForm.ID?e.addUserForm.Password:e.addUserForm.PasswordEdit,RoleID:e.addUserForm.RoleID,OrganizationId:e.addUserForm.OrganizationId};e.apiSever.post(e.URl.ADMIN_EDIT,a,(function(t){e.$message({type:"success",message:"操作成功!"}),e.showAddUser=!1,e.resetVerForm(),e.loadData()}))}))},viewPass:function(e){this.passwordInputType?this.passwordInputType="":this.passwordInputType="password"},resetVerForm:function(){var e=this;this.addUserForm={ID:0,Mobile:"",Password:"",Name:"",PasswordEdit:"",RoleID:"",OrganizationId:this.formSearch.id},this.showAddUser=!1,this.$nextTick((function(){e.$refs["addUserForm"]&&e.$refs["addUserForm"].resetFields()}))},editUser:function(e){this.addUserForm={ID:e.ID,Mobile:e.Mobile,Password:"",Name:e.Name,RoleID:e.RoleID,OrganizationId:e.OrganizationId},this.showAddUser=!0},delPage:function(e){var t=this,a={ID:e};this.$confirm("确定要将此管理员删除吗?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((function(){t.apiSever.post(t.URl.ADMIN_DELETE,a,(function(e){t.loadData(),t.$message({type:"success",message:"删除成功"})}))})).catch((function(){t.$message({type:"info",message:"取消操作"})}))}}},i=o,n=(a("b5e4"),a("2877")),l=Object(n["a"])(i,r,s,!1,null,"6ede1650",null);t["default"]=l.exports}}]);
//# sourceMappingURL=chunk-1ccde611.2d0f64ff.js.map