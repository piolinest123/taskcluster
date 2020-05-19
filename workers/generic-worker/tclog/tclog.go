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

type DevNull struct{}

func (DevNull) Log(prefix, message string)             {}
func (DevNull) Info(message string)                    {}
func (DevNull) Infof(format string, v ...interface{})  {}
func (DevNull) Warn(message string)                    {}
func (DevNull) Warnf(format string, v ...interface{})  {}
func (DevNull) Error(message string)                   {}
func (DevNull) Errorf(format string, v ...interface{}) {}
