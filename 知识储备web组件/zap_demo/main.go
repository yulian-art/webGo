package main

import (
	"go.uber.org/zap"
	"net/http"
)
var logger *zap.Logger

func main(){
	InitLogger()
	defer logger.Sync()
	simpleHttpGet("http://www.google.com")
}
func InitLogger(){
	logger, _ = zap.NewProduction()
}
func simpleHttpGet(url string){
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url...",
			zap.String("url", url),
			zap.Error(err),
		)
	}else {
		logger.Info(
			"success...",
			zap.String("url",url),
			zap.Int("code",http.StatusOK),
		)
		resp.Body.Close()
	}
}