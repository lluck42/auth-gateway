package admin

import (
	"auth-gateway/bean"
	"auth-gateway/model/admin"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

func CaptchaId(c *gin.Context) {

	var captchaId = captcha.NewLen(4)

	c.String(200, captchaId)
}

func Captcha(c *gin.Context) {

	var captchaId = c.Query("captchaId")

	c.Header("captchaid", captchaId)

	if err := captcha.WriteImage(c.Writer, captchaId, 100, 50); err != nil {
		bean.Log.Warning(err)
		c.String(400, "验证码生成失败")
		return
	}
}

func Login(c *gin.Context) {

	var req struct {
		Mobile    string `binding:"max=255"`
		Password  string `binding:"max=255"`
		CaptchaId string `binding:"max=255"`
		Digits    string `binding:"max=255"`
		UserToken string `binding:"max=255"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		bean.Log.Warning(err)
		c.String(400, err.Error())
		return
	}

	// 请求记录
	var rec admin.AdminRecord
	rec.Method = c.Request.Method
	rec.ApiPath = c.Request.URL.Path
	rec.Ip = c.ClientIP()
	rec.AdminName = "尝试登录账号：" + req.Mobile
	if err := bean.MysqlDb.Save(&rec).Error; err != nil {
		bean.Log.Error(err)
		return
	}

	// 验证码判断
	if !captcha.VerifyString(req.CaptchaId, req.Digits) {
		bean.Log.Info("验证码错误:" + req.Digits)
		c.String(400, "验证码错误")
		return
	}

	// 密码加密
	if s, err := admin.Md5(req.Password); err != nil {
		bean.Log.Error(err)
		c.String(500, "内部错误m")
		return
	} else {
		req.Password = s
	}

	// 对比查找用户
	var one admin.Admin

	if err := bean.MysqlDb.Where("`mobile`=? and `password`=?", req.Mobile, req.Password).First(&one).Error; err != nil {
		bean.Log.Error(err)
		c.String(400, "账号或密码错误")
		return
	}

	c.PureJSON(200, one)

	// 更新请求记录
	rec.AdminName = one.Name
	rec.AdminID = one.ID

	if err := bean.MysqlDb.Save(&rec).Error; err != nil {
		bean.Log.Error(err)
		return
	}

	// 放入admin 信息，record 中间件使用
	c.Set("admin", one)
}

func Logout(c *gin.Context) {
	// 获取登录用户对象
	var one admin.Admin
	if err := one.Get(c); err != nil {
		bean.Log.Error(err)
		c.String(500, err.Error())
		return
	}
	//

	if err := one.NewToken(); err != nil {
		bean.Log.Error(err)
		c.String(500, err.Error())
		return
	}

	var two admin.Admin
	two.ID = one.ID
	two.Token = one.Token

	if err := bean.MysqlDb.Updates(&two).Error; err != nil {
		bean.Log.Error(err)
		c.String(500, err.Error())
		return
	}

	c.String(200, "ok")
}
