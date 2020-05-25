package params

import "flag"

type Params struct {
	Development	bool
	LogFile		string
}

var instance *Params

func GetInstance() *Params  {
	if instance == nil {
		env := flag.String("env", "", "Environment")
		logFile := flag.String("log", "/var/log/1c2jira.log", "Log file")
		flag.Parse()
		instance = &Params{false,*logFile}
		if *env!="dev" { instance.Development=false }
	}
	return instance
}