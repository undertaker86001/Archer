package logger

// Logger logger 对象
var loggerInterface LoggerInterface

func init() {
	loggerInterface = &DefaultLogger{}
}

// InitLogger 初始化logger对象
func InitLogger(ILogger LoggerInterface) {
	loggerInterface = ILogger
}

// LoggerInterface 日志接口
type LoggerInterface interface {
	Debugf(format string, params ...interface{})

	Infof(format string, params ...interface{})

	Warnf(format string, params ...interface{})

	Errorf(format string, params ...interface{})

	Debug(v ...interface{})

	Info(v ...interface{})

	Warn(v ...interface{})

	Error(v ...interface{})
}

// Debugf debug 格式化
func Debugf(format string, params ...interface{}) {
	loggerInterface.Debugf(format, params)
}

// Infof 打印info
func Infof(format string, params ...interface{}) {
	loggerInterface.Infof(format, params)
}

// Warnf warn格式化
func Warnf(format string, params ...interface{}) {
	loggerInterface.Warnf(format, params)
}

// Errorf error格式化
func Errorf(format string, params ...interface{}) {
	loggerInterface.Errorf(format, params)
}

// Debug 打印debug
func Debug(v ...interface{}) {
	loggerInterface.Debug(v)
}

// Info 打印Info
func Info(v ...interface{}) {
	loggerInterface.Info(v)
}

// Warn 打印Warn
func Warn(v ...interface{}) {
	loggerInterface.Warn(v)
}

// Error 打印Error
func Error(v ...interface{}) {
	loggerInterface.Error(v)
}

// DefaultLogger 默认日志实现
type DefaultLogger struct {
}

// Debugf debug 格式化
func (d *DefaultLogger) Debugf(format string, params ...interface{}) {

}

// Infof 打印info
func (d *DefaultLogger) Infof(format string, params ...interface{}) {

}

// Warnf warn格式化
func (d *DefaultLogger) Warnf(format string, params ...interface{}) {
}

// Errorf error格式化
func (d *DefaultLogger) Errorf(format string, params ...interface{}) {
}

// Debug 打印debug
func (d *DefaultLogger) Debug(v ...interface{}) {

}

// Info 打印Info
func (d *DefaultLogger) Info(v ...interface{}) {

}

// Warn 打印Warn
func (d *DefaultLogger) Warn(v ...interface{}) {
}

// Error 打印Error
func (d *DefaultLogger) Error(v ...interface{}) {
}
