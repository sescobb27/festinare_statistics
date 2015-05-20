package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sescobb27/festinare_statistics/database"
)

func signalHandler() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGINT,
		syscall.SIGKILL,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGABRT,
	)
	<-signalChan
	println("Closing DB Connection")
	database.Session.Close()
	os.Exit(0)
}

func main() {

}
