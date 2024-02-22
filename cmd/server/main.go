package main

import (
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	api "sber-test/internal/api/http"
	"sber-test/internal/services/deposit"
	"sber-test/pkg/options"

	log "github.com/sirupsen/logrus"
)

// Блок переменных приложения
const (
	// applicationName - название приложения.
	applicationName = "sber-test"
)

var (
	configName    string
	exitSignal    chan bool
	signalChannel chan os.Signal
	apiServer     *api.Server
)

func init() {
	flag.StringVar(&configName, "config", "sber-test", "configuration file name")
	exitSignal = make(chan bool)
}

func Run() error {
	log.Info("Running App")
	flag.Parse()

	opt, err := options.LoadConfig(configName)
	if err != nil {
		log.Errorf("Unable to load configuration file %s: %s", configName, err)
		return err
	}

	depositService := deposit.New()

	apiServer = api.Init(opt, depositService)
	go func() {
		if err := apiServer.Serve(); err != nil && err != http.ErrServerClosed {
			log.WithError(err).Fatal("Unable to server HTTP API")
		} else if err == http.ErrServerClosed {
			log.Infof("HTTP server closed")
		}
	}()

	time.Sleep(time.Second * 1)
	log.Infof("HTTP API server started on \"%s\"", opt.APIAddr)

	go initSignals()

	<-exitSignal

	return nil

}
func main() {
	err := Run()
	if err != nil {
		log.Errorf("Error running app: %s", err)
	}

}

func initSignals() {
	log.Info("Try to initialize signals...")
	signalChannel = make(chan os.Signal)
	signal.Notify(signalChannel, syscall.SIGTERM)
	signal.Notify(signalChannel, syscall.SIGINT)
	signal.Notify(signalChannel, syscall.SIGKILL)

	for {
		select {
		case s := <-signalChannel:
			switch s {
			case syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL:
				close(signalChannel)
				log.Warnf("We got %s, shutdown application...", s)
				_ = apiServer.Stop()
				exitSignal <- true
				return
			}
		}
	}
}
