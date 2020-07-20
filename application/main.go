/**
 * @Author: 网红电商组
 * @Email: 233_666@qq.com
 * @Date: 2020/7/18 11:15 下午
 * @Description: application启动入口
 */
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/Stupid-K/blockchain-real-goods/application/blockchain"
	_ "github.com/Stupid-K/blockchain-real-goods/application/docs"
	"github.com/Stupid-K/blockchain-real-goods/application/pkg/setting"
	"github.com/Stupid-K/blockchain-real-goods/application/routers"
	"github.com/Stupid-K/blockchain-real-goods/application/service"
	"log"
	"net/http"
	"time"
)

func init() {
	setting.Setup()
}

// @title 基于区块链网红信息追根溯源系统api文档
// @version 1.0
// @description 基于区块链网红信息追根溯源系统api文档
// @contact.name 网红电商组
// @contact.email
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	timeLocal, err := time.LoadLocation("Asia/Chongqing")
	if err != nil {
		log.Printf("时区设置失败 %s", err)
	}
	time.Local = timeLocal
	blockchain.Init()
	go service.Init()
	gin.SetMode(setting.ServerSetting.RunMode)
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	if err := server.ListenAndServe(); err != nil {
		log.Printf("start http server failed %s", err)
	}
}
