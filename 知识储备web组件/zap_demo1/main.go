package main

import (
	"net/http"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"github.com/natefinch/lumberjack"
)

var sugarLogger *zap.SugaredLogger
var logger *zap.Logger
func main(){
	InitLogger()

	defer logger.Sync()
	simpleHttpGet("http://www.google.com")
}

func InitLogger(){
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zap.DebugLevel)
	
	logger = zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()

}


func getEncoder() zapcore.Encoder {
	// return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,

	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// func getLogWriter() zapcore.WriteSyncer {
// 	file, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0740)
// 	return zapcore.AddSync(file)
// }

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger {
		Filename: "./test.log",
		MaxSize: 10, // 到10M后自动压缩
		MaxBackups: 5,
		MaxAge: 30,
		Compress: false,
	}

	return zapcore.AddSync(lumberJackLogger)
}

func simpleHttpGet(url string){
	resp, err := http.Get(url)
	if err != nil{

		logger.Error(
			"ERROR fetching url",
			zap.String("url", url),
			zap.Error(err),
		)

	}else {
		logger.Info(
			"success....",
			zap.String("url", url),
			zap.Int("code", http.StatusOK),
		)
		resp.Body.Close()
	}
}