package tclog

type Logger interface {
	Log(prefix, message string)
	Info(message string)
	Infof(format string, v ...interface{})
	Warn(message string)
	Warnf(format string, v ...interface{})
	Error(message string)
	Errorf(format string, v ...interface{})
}
