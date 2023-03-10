package middleware

import (
	"bytes"
	"douyin_service/pkg/logger"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"time"
)

var Logger *logger.Logger

func SetupLogger(logSavePath, logFileName string) error {
	Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  logSavePath + "/" + logFileName,
		MaxSize:   600, // 600MB
		MaxAge:    10,  // 10天
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = bodyWriter
		c.Next()

		beginTime := time.Now().Unix()
		endTime := time.Now().Unix()
		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}

		Logger.WithFields(fields).Infof("access log: method: %s, status_code: %d, begin_time: %d, end_time: %d",
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)
	}
}
