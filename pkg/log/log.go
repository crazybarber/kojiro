package log

type Logger interface {
	Info(v ...interface{})
	Infof(message string, args ...interface{})
	Warn(v ...interface{})
	Warnf(message string, args ...interface{})
	Fatal(v ...interface{})
	Fatalf(message string, args ...interface{})
}
