package daemon

import (
	"context"
	"net/http"
	"time"

	//"github.com/facebookgo/grace/gracehttp"
	"os"
	"os/signal"
)

func Daemon() {

}

var (
	IsDebug bool = true
)

//initialize the signal controal
func SetSignal(serve *http.Server) {

	var quit chan int = make(chan int, 0)
	go func() {
		serve.ListenAndServe()
		quit <- 1
	}()

	var c chan os.Signal
	c = make(chan os.Signal)

	go func() {
		signal.Notify(c, os.Interrupt, os.Kill)
	}()

HERE:
	for {
		switch <-c {
		case os.Interrupt:
			//log.Log.Error("%s", "Progess was interrupted")
			break HERE
		case os.Kill:
			//log.Log.Error("%s", "Progess was killed")
			break HERE
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	serve.Shutdown(ctx)
	//shutdown gracefully
	select {
	case <-quit:
		//log.Log.Info("Server already shutdown gracefully")
	case <-ctx.Done():
		//log.Log.Info("Server already shutdown gracefully!")
	}
}
