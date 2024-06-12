package route

import (
	"gq_volunteer_platform/controller"
	"gq_volunteer_platform/controller/admin"
	"gq_volunteer_platform/middleware"

	"github.com/gin-gonic/gin"
)

func Admin(router *gin.Engine) {

	// router.Static("/upload", "upload") // 服务静态文件 相对路径为 项目启动时 用户所在目录

	router.GET("time", controller.Time)
	router.GET("/admin/CaptchaId", admin.CaptchaId)
	router.GET("/admin/Captcha", admin.Captcha)
	router.POST("/admin/Login", admin.Login)

	// admin
	adminGroup := router.Group("admin")
	adminGroup.Use(middleware.Record, middleware.Login, middleware.Permission)
	{
		adminGroup.GET("test", admin.Test)
		adminGroup.GET("Info", admin.Info)
		adminGroup.GET("MyList", admin.MyList)
		adminGroup.POST("Edit", admin.Edit)
		adminGroup.POST("Delete", admin.Delete)
		adminGroup.POST("Logout", admin.Logout)
		adminGroup.POST("ChangePassword", admin.ChangePassword)

		adminGroup.GET("RecordList", admin.RecordList)

		adminGroup.GET("RoleAllList", admin.RoleAllList)
		adminGroup.GET("RoleList", admin.RoleList)
		adminGroup.GET("RoleDetail", admin.RoleDetail)
		adminGroup.POST("RoleEdit", admin.RoleEdit)
		adminGroup.POST("RoleDelete", admin.RoleDelete)
		adminGroup.POST("RolePermissionEdit", admin.RolePermissionEdit)

		adminGroup.GET("MenuTree", admin.MenuTree)
		adminGroup.GET("MenuDetail", admin.MenuDetail)
		adminGroup.POST("MenuEdit", admin.MenuEdit)
		adminGroup.POST("MenuDelete", admin.MenuDelete)
	}

}
