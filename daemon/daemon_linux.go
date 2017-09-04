package daemon

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/facebookgo/grace/gracehttp"
)

var (
	IsDebug bool = false
)

func Daemon() {

	debug := flag.Bool("debug", false, "debug mode ?")
	flag.Parse()
	IsDebug = *debug
	if os.Getenv("LISTEN_FDS") == "" && os.Getppid() != 1 && !*debug {
		var se string = os.Args[0]
		var args []string = os.Args[1:]
		filter(args)
		cmd := exec.Command(se, args...)
		err := cmd.Start()

		if err != nil {
			Log.Error("Start unsuccessfully")
		}
		os.Exit(0)
	}
}

//Filter the flag of daemon
func filter(args []string) {

	for arg := range args {
		if strings.EqualFold(args[arg], "--debug") || strings.EqualFold(args[arg], "-debug") {
			args[arg] = ""
			return
		}
	}

}

//initialize the signal controal
func SetSignal(serve *http.Server) {

	gracehttp.Serve(serve)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	serve.Shutdown(ctx)
	Log.Info("Server already shutdown gracefully!")
}
