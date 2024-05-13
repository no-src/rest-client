package flag

import "flag"

var (
	ConfigFile string
	HttpFile   string
	RequestId  int
	IsRun      bool
	LogConf    string
)

func InitFlags() {
	flag.StringVar(&ConfigFile, "conf", "", "specified a config file")
	flag.StringVar(&HttpFile, "http", "request.http", "specified a http file")
	flag.IntVar(&RequestId, "id", 0, "specified the http request by http request id")
	flag.BoolVar(&IsRun, "run", false, "run the http request")
	flag.StringVar(&LogConf, "log_conf", "", "specified a log config file")
	flag.Parse()
}
