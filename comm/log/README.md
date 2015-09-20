# comm-log [![GoDoc](http://godoc.org/gopkg.in/goyy/goyy.v0?status.png)](http://godoc.org/gopkg.in/goyy/goyy.v0/comm/log)
log library for Go

# Installation
`go get gopkg.in/goyy/goyy.v0/comm/log`

# Priority
Ptrace < Pdebug < Pinfo < Pwarn < Perror < Pcritical < Pprint

# Usage
	log.SetPrefix("[log]")
	log.SetPriority(log.Perror)
	log.SetLayouts(log.Lstd | log.Llongfile)
	log.Trace("=====trace=====")
	log.Debug("=====debug=====")
	log.Info("=====info=====")
	log.Warn("=====warn=====")
	log.Error("=====error=====")
	log.Critical("=====critical=====")
	log.Print("=====print=====")

	logging := log.New("[logging]")
	logging.Trace("=====trace=====")
	logging.Debug("=====debug=====")
	logging.Info("=====info=====")
	logging.Warn("=====warn=====")
	logging.Error("=====error=====")
	logging.Critical("=====critical=====")
	logging.Print("=====print=====")
