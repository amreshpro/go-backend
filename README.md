# GO Dev

#### Logging

```go

package logger

import "go.uber.org/zap"

var Log *zap.SugaredLogger

func InitLogger(env string){

	var zapLogger *zap.Logger
	var err error
	if(env == "production"){
    zapLogger,err = zap.NewProduction()
	}else{
    zapLogger,err = zap.NewDevelopment()
	}
	if(err!=nil){
		panic("Failed to start logger service : " + err.Error())
	}

	Log = zapLogger.Sugar()

}

```