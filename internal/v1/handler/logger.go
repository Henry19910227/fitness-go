package handler

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/gin-gonic/gin"
)

type logger struct {
	logTool tool.Logger
	jwtTool tool.JWT
}

func NewLogger (logTool tool.Logger, jwtTool tool.JWT) Logger {
	return &logger{logTool: logTool, jwtTool: jwtTool}
}

func (handler *logger) Set(c *gin.Context, level LogLevel, tag string, code int, msg string) {
	fields := handler.GetFields(c, tag, code)
	switch level {
	case Trace:
		handler.logTool.Trace(fields, msg)
		break
	case Debug:
		handler.logTool.Debug(fields, msg)
		break
	case Info:
		handler.logTool.Info(fields, msg)
		break
	case Warn:
		handler.logTool.Warn(fields, msg)
		break
	case Error:
		handler.logTool.Error(fields, msg)
		break
	case Fatal:
		handler.logTool.Fatal(fields, msg)
		break
	case Panic:
		handler.logTool.Panic(fields, msg)
		break
	}
}

func (handler *logger) GetFields(c *gin.Context, tag string, code int) map[string]interface{} {
	var uid int64
	var path string
	var hostIP string
	var body interface{}
	var rid interface{}
	if c != nil {
		uid, _ = handler.jwtTool.GetIDByToken(c.Request.Header.Get("token"))
		path = c.FullPath()
		hostIP = c.ClientIP()
		body = c.Value("Body")
		rid = c.Value("X-Request-Id")
	}

	fields := map[string]interface{}{
		"tag":     tag,
		"code":    code,
		"path":    path,
		"host_ip": hostIP,
		"uid":     uid,
		"body":    body,
		"rid":     rid,
	}
	return fields
}
