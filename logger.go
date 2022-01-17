package log

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/TwiN/go-color"
)

// Level is a log level
type Level uint8

// Available log levels
const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	levelCount
)

// A Logger wraps multiple loggers (one per log level).
// Individual loggers may be enabled or disabled depending on the current level (which can be set using the SetLevel function).
// Messages logged to a disabled level will be discarded.
type Logger struct {
	zero    *log.Logger
	loggers [levelCount]*log.Logger
	level   Level
	active  [levelCount]*log.Logger
}

// New returns a new Logger which outputs enabled levels to out.
// level allows to set the initial enabled level.
func New(out io.Writer, level Level) *Logger {
	hasColors := isTTY(out)

	logger := &Logger{
		zero: log.New(nilWriter{}, "", 0),
		loggers: [levelCount]*log.Logger{
			log.New(out, getPrefix("debug ", color.Purple, hasColors), log.LstdFlags|log.Lmsgprefix),
			log.New(out, getPrefix("info  ", color.Green, hasColors), log.LstdFlags|log.Lmsgprefix),
			log.New(out, getPrefix("warn  ", color.Yellow, hasColors), log.LstdFlags|log.Lmsgprefix),
			log.New(out, getPrefix("error ", color.Red, hasColors), log.LstdFlags|log.Lmsgprefix),
		},
	}
	logger.SetLevel(level)

	return logger
}

// IsEnabled returns whether level is enabled or not.
func (l *Logger) IsEnabled(level Level) bool {
	return l.level <= level
}

// For returns the Logger to use for the given level.
func (l *Logger) For(level Level) *log.Logger {
	return l.active[level]
}

func (l *Logger) resolve(level Level) *log.Logger {
	if l.IsEnabled(level) {
		return l.loggers[level]
	}
	return l.zero
}

// SetLevel changes the minimal enabled log level.
// All loggers below level are disabled, other loggers are enabled.
func (l *Logger) SetLevel(level Level) {
	if level < 0 || level >= levelCount {
		panic(fmt.Errorf("Invalid level %d", level))
	}
	l.level = level
	l.active[LevelDebug] = l.resolve(LevelDebug)
	l.active[LevelInfo] = l.resolve(LevelInfo)
	l.active[LevelWarn] = l.resolve(LevelWarn)
	l.active[LevelError] = l.resolve(LevelError)
}

// Level returns the currently enabled level
func (l *Logger) Level() Level {
	return l.level
}

// Debug is equivalent to l.For(LevelDebug).Print(v...)
func (l *Logger) Debug(v ...interface{}) {
	l.active[LevelDebug].Print(v...)
}

// Debugln is equivalent to l.For(LevelDebug).Println(v...)
func (l *Logger) Debugln(v ...interface{}) {
	l.active[LevelDebug].Println(v...)
}

// Debugf is equivalent to l.For(LevelDebug).Printf(format, v...)
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.active[LevelDebug].Printf(format, v...)
}

// Info is equivalent to l.For(LevelInfo).Print(v...)
func (l *Logger) Info(v ...interface{}) {
	l.active[LevelInfo].Print(v...)
}

// Infoln is equivalent to l.For(LevelInfo).Println(v...)
func (l *Logger) Infoln(v ...interface{}) {
	l.active[LevelInfo].Println(v...)
}

// Infof is equivalent to l.For(LevelInfo).Printf(format, v...)
func (l *Logger) Infof(format string, v ...interface{}) {
	l.active[LevelInfo].Printf(format, v...)
}

// Warn is equivalent to l.For(LevelWarn).Print(v...)
func (l *Logger) Warn(v ...interface{}) {
	l.active[LevelWarn].Print(v...)
}

// Warnln is equivalent to l.For(LevelWarn).Println(v...)
func (l *Logger) Warnln(v ...interface{}) {
	l.active[LevelWarn].Println(v...)
}

// Warnf is equivalent to l.For(LevelWarn).Printf(format, v...)
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.active[LevelWarn].Printf(format, v...)
}

// Error is equivalent to l.For(LevelError).Print(v...)
func (l *Logger) Error(v ...interface{}) {
	l.active[LevelError].Print(v...)
}

// Errorln is equivalent to l.For(LevelError).Println(v...)
func (l *Logger) Errorln(v ...interface{}) {
	l.active[LevelError].Println(v...)
}

// Errorf is equivalent to l.For(LevelError).Printf(format, v...)
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.active[LevelError].Printf(format, v...)
}

// Print is equivalent to l.Info(v...)
func (l *Logger) Print(v ...interface{}) {
	l.active[LevelInfo].Print(v...)
}

// Printf is equivalent to l.Infof(format, v...)
func (l *Logger) Printf(format string, v ...interface{}) {
	l.active[LevelInfo].Printf(format, v...)
}

// Println is equivalent to l.Infoln(v...)
func (l *Logger) Println(v ...interface{}) {
	l.active[LevelInfo].Println(v...)
}

// Fatal is equivalent to l.For(LevelError).Fatal(v...)
// Note that the program will exit even if the Error level is diabled.
func (l *Logger) Fatal(v ...interface{}) {
	l.active[LevelError].Fatal(v...)
}

// Fatalf is equivalent to l.For(LevelError).Fatalf(format, v...)
// Note that the program will exit even if the Error level is diabled.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.active[LevelError].Fatalf(format, v...)
}

// Fatalln is equivalent to l.For(LevelError).Fatalln(v...)
// Note that the program will exit even if the Error level is diabled.
func (l *Logger) Fatalln(v ...interface{}) {
	l.active[LevelError].Fatalln(v...)
}

// Panic is equivalent to l.For(LevelError).Panic(v...)
// Note that panic() will be called even if the Error level is diabled.
func (l *Logger) Panic(v ...interface{}) {
	l.active[LevelError].Panic(v...)
}

// Panicf is equivalent to l.For(LevelError).Panicf(format, v...)
// Note that panic() will be called even if the Error level is diabled.
func (l *Logger) Panicf(format string, v ...interface{}) {
	l.active[LevelError].Panicf(format, v...)
}

// Panicln is equivalent to l.For(LevelError).Panicln(v...)
// Note that panic() will be called even if the Error level is diabled.
func (l *Logger) Panicln(v ...interface{}) {
	l.active[LevelError].Panicln(v...)
}

func getPrefix(prefix, colorCode string, hasColors bool) string {
	if hasColors {
		return color.Colorize(colorCode, prefix)
	}
	return prefix
}

func isTTY(w io.Writer) bool {
	f, ok := w.(*os.File)
	if ok {
		fileInfo, err := f.Stat()
		if err == nil {
			return fileInfo.Mode()&os.ModeCharDevice != 0
		}
	}

	return false
}
