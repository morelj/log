package log

import "os"

var dft = New(os.Stdout, LevelDebug)

// Default returns the default Logger instance
func Default() *Logger {
	return dft
}

// Debug calls Debug(v...) on the default Logger
func Debug(v ...interface{}) {
	dft.Debug(v...)
}

// Debugln calls Debugln(v...) on the default Logger
func Debugln(v ...interface{}) {
	dft.Debugln(v...)
}

// Debugf calls Debugf(format, v...) on the default Logger
func Debugf(format string, v ...interface{}) {
	dft.Debugf(format, v...)
}

// Info calls Info(v...) on the default Logger
func Info(v ...interface{}) {
	dft.Info(v...)
}

// Infoln calls Infoln(v...) on the default Logger
func Infoln(v ...interface{}) {
	dft.Infoln(v...)
}

// Infof calls Infof(format, v...) on the default Logger
func Infof(format string, v ...interface{}) {
	dft.Infof(format, v...)
}

// Warn calls Warn(v...) on the default Logger
func Warn(v ...interface{}) {
	dft.Warn(v...)
}

// Warnln calls Warnln(v...) on the default Logger
func Warnln(v ...interface{}) {
	dft.Warnln(v...)
}

// Warnf calls Warnf(format, v...) on the default Logger
func Warnf(format string, v ...interface{}) {
	dft.Warnf(format, v...)
}

// Error calls Error(v...) on the default Logger
func Error(v ...interface{}) {
	dft.Error(v...)
}

// Errorln calls Errorln(v...) on the default Logger
func Errorln(v ...interface{}) {
	dft.Errorln(v...)
}

// Errorf calls Errorf(format, v...) on the default Logger
func Errorf(format string, v ...interface{}) {
	dft.Errorf(format, v...)
}

// SetLevel sets the loggin level of the default Logger
func SetLevel(level Level) {
	dft.SetLevel(level)
}
