package util

import (
	"os"
	"os/signal"
	"syscall"
)

// HandleSigterm deals with SIGTERM signal to allow graceful shutdown
func HandleSigterm(handleExit func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		handleExit()
		os.Exit(1)
	}()
}
