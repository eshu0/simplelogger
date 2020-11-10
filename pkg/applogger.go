package simplelogger

import (
	sli "github.com/eshu0/simplelogger/pkg/interfaces"
)

type AppLogger struct {

	//inherit from interface
	sli.IAppLogger

	Log sli.ISimpleLogger `json:"-"`
}

// This is the simplest application log generator
// The os.args[0] is used for filename and the session is random
//func NewAppLogger() *AppLogger {
//	alog := AppLogger{}
//	alog.Log = NewApplicationNowLogger(RandomSessionID())
//	return alog
//}

/*
	START/FINISH LOG FUNCTIONS
*/

// the logging functions are here
func (al *AppLogger) Start() {
	//al = NewAppLogger()
	al.Log = NewApplicationNowLogger(RandomSessionID())

	// lets open a file log using the session
	al.Log.OpenAllChannels()

	//defer the close till the shell has closed
	defer al.Finish()
}

// the logging functions are here
func (al *AppLogger) Finish() {
	al.Log.CloseAllChannels()
}

/*
	APP LOG FUNCTIONS
*/

// the logging functions are here
func (al *AppLogger) LogDebug(cmd string, data ...interface{}) {
	if al.Log != nil {
		al.Log.LogDebug(cmd, data)
	}
}

func (ssl *SimpleLogger) LogWarn(cmd string, data ...interface{}) {
	if al.Log != nil {
		al.Log.LogWarn(cmd, data)
	}
}

func (ssl *SimpleLogger) LogInfo(cmd string, data ...interface{}) {
	if al.Log != nil {
		al.Log.LogInfo(cmd, data)
	}
}

func (ssl *SimpleLogger) LogError(cmd string, data ...interface{}) {
	if al.Log != nil {
		al.Log.LogError(cmd, data)
	}
}

// This Log error allows errors to be logged .Error() is the data written
func (ssl *SimpleLogger) LogErrorE(cmd string, data error) {
	if al.Log != nil {
		al.Log.LogErrorE(cmd, data)
	}
}

// the logging functions are here
func (ssl *SimpleLogger) LogDebugf(cmd string, msg string, data ...interface{}) {
	if al.Log != nil {
		al.Log.LogDebugf(cmd, data)
	}
}

func (ssl *SimpleLogger) LogWarnf(cmd string, msg string, data ...interface{}) {
	if al.Log != nil {
		al.Log.LogWarnf(cmd, data)
	}
}

func (ssl *SimpleLogger) LogInfof(cmd string, msg string, data ...interface{}) {
	if al.Log != nil {
		al.Log.LogInfof(cmd, data)
	}
}

func (ssl *SimpleLogger) LogErrorf(cmd string, msg string, data ...interface{}) {
	if al.Log != nil {
		al.Log.LogErrorf(cmd, data)
	}
}