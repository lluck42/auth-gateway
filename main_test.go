package main

import (
	"fmt"
	"gq_volunteer_platform/bean"
	"testing"
	"time"
)

func init() {
	// 配置项初始化
	bean.ConfigInit()
	// 日志初始化
	bean.LogInit()
	// mysql 初始化
	bean.MysqlInit()
	//11
}

func TestA(t *testing.T) {
	//1
	var start_time = ""
	StartTime, err := time.ParseInLocation(time.DateTime, start_time, time.Local)
	if err != err {
		var err = fmt.Sprintf("时间解析失败:%s", start_time)
		fmt.Println(err)
		return
	}
	//

	dieline := StartTime.Add(time.Minute * 30).Format(time.DateTime)

	fmt.Println(StartTime)
	fmt.Println(dieline)
}

func TestTest(t *testing.T) {

	// var rq weixin.URLLinkReq

	// rq.Path = "pages/index/index"
	// rq.Query = "aa=bb"
	// rq.EnvVersion = "trial"

	// var _, err = rq.URLLinkGet()
	// if err != nil {
	// 	fmt.Printf("+%v", err)
	// }
}
