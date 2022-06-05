package mlog

import (
	"github.com/astaxie/beego/logs"
	"os"
	"path"
)

var mainLogger *logs.BeeLogger

func init() {
	// MainLogger
	// Common for all
	{
		fullPath := path.Join("SERVER", "mainLogs.log")
		lc := `{"filename":"` + fullPath + `", "level":7, "maxlines":0, "maxsize":5000000, "daily":true, "maxdays":14, "color":true, "rotate":true }`
		mainLogger = logs.NewLogger(5000)
		// Включаем отображение файла и номер строки
		mainLogger.EnableFuncCallDepth(true)
		// Поднимаем выше уровень выше для EnableFuncCallDepth, чтобы не отображалось [mlog-tmp:30]
		mainLogger.SetLogFuncCallDepth(3)

		if err := mainLogger.SetLogger(logs.AdapterConsole, lc); err != nil {
			panic(err)
		}

		if err := mainLogger.SetLogger(logs.AdapterFile, lc); err != nil {
			panic(err)
		}
	}
}

// Trace ...
func Trace(format string, v ...interface{}) {
	mainLogger.Trace(format, v...)
}

// Debug ...
func Debug(format string, v ...interface{}) {
	mainLogger.Debug(format, v...)
}

// Notice ...
func Notice(format string, v ...interface{}) {
	mainLogger.Notice(format, v...)
}

// Info ...
func Info(format string, v ...interface{}) {
	mainLogger.Info(format, v...)
}

// Alert ...
func Alert(format string, v ...interface{}) {
	mainLogger.Alert(format, v...)
}

// Warn ...
func Warn(format string, v ...interface{}) {
	mainLogger.Warn(format, v...)
}

// Error ...
func Error(format string, v ...interface{}) {
	mainLogger.Error(format, v...)
}

// Critical ...
func Critical(format string, v ...interface{}) {
	mainLogger.Critical(format, v...)
	os.Exit(1)
}
