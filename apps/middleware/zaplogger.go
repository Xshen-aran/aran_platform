package middleware

import (
	"bytes"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Xshen-aran/aran_platform/apps/utils/json"
	"github.com/Xshen-aran/aran_platform/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var lg *zap.Logger

// func getLogwriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
// 	// filename = strings.Replace(filename, "$$", time.Now().Format("2006_01_02"), 1)
// 	lumberJackLogger := &lumberjack.Logger{
// 		Filename:   filename,
// 		MaxSize:    maxSize,
// 		MaxBackups: maxBackup,
// 		MaxAge:     maxAge,
// 	}
// 	return zapcore.AddSync(lumberJackLogger)
// }

func getLogwriterData(filename, maxSize, maxBackup, maxAge string) zapcore.WriteSyncer {
	// filename = strings.Replace(filename, "$$", time.Now().Format("2006_01_02"), 1)
	lumberJackLogger, _ := rotatelogs.New(
		filename+".%Y%m%d.log",
		rotatelogs.WithMaxAge(30*24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func InitLogger() (err error) {
	// logerWriter := getLogwriter(c.Filename, c.MaxSize, c.MaxBackups, c.MaxAge)
	// logerWriter := getLogwriterData(c.Filename, c.MaxSize, c.MaxBackups, c.MaxAge)
	logerWriter := getLogwriterData(config.Env["LOG_FILE_NAME"], config.Env["LOG_FILE_MAX_SIZE"], config.Env["LOG_FILE_MAX_BACKUPS"], config.Env["LOG_FILE_MAX_AGE"])
	encoder := getEncoder()
	var l = new(zapcore.Level)
	var core zapcore.Core
	// if config.Env["GIN_MODE"] == "release" {
	// 	c.Level = "info"
	// }
	// err = l.UnmarshalText([]byte(c.Level))
	// 控制台打印
	if config.Env["GIN_MODE"] == "debug" {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, logerWriter, l),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, logerWriter, l)
	}

	lg = zap.New(core, zap.AddCaller())
	// zapcore.AddSync()
	zap.ReplaceGlobals(lg)
	return
}
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					lg.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					lg.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					lg.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			bodyBytes []byte
			username  string
		)
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter
		c.Next()

		responseBody := bodyLogWriter.body.String()
		// var responseCode int
		// var responseMsg string
		var responseData interface{}
		if responseBody != "" {
			json.Unmarshal([]byte(responseBody), &responseData)

		}
		if c.Keys["username"] == nil {
			username = "-"
		} else {
			username = c.Keys["username"].(string)
		}
		cost := time.Since(start)
		lg.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("user", username),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("body", string(bodyBytes)),
			zap.String("query", query),
			zap.Any("response", responseData),
			// zap.String("responseData", string(responseData)),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
