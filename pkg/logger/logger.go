package logger

func Infow(msg string, keysAndValues ...interface{}) {
	_SugarLog.Infow(msg, keysAndValues)
}

func Info(args ...interface{}) {
	_SugarLog.Info(args)
}

func Infof(template string, args ...interface{}) {
	_SugarLog.Infof(template, args)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	_SugarLog.Errorw(msg, keysAndValues)
}

func Error(args ...interface{}) {
	_SugarLog.Error(args)
}

func Errorf(template string, args ...interface{}) {
	_SugarLog.Errorf(template, args)
}
