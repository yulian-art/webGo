package main

import (
	"go.uber.org/zap"
	"net/http"
)
//定义全局logger实例
var logger *zap.Logger
var sugarlogger *zap.SugaredLogger

func main(){
	InitLogger()
	//在程序推出之前把缓冲区的日志都刷到磁盘上
	defer logger.Sync()
	simpleHttpGet("http://www.google.com")
}
func InitLogger(){
	logger, _ = zap.NewProduction()
	sugarlogger =  logger.Sugar()
}
func simpleHttpGet(url string){
	resp, err := http.Get(url)
	if err != nil {
		sugarlogger.Error(
			"Error fetching url...",
			zap.String("url", url),
			zap.Error(err),
		)
	}else {
		sugarlogger.Info(
			"success...",
			zap.String("url",url),
			zap.Int("code",http.StatusOK),
		)
		resp.Body.Close()
	}
}