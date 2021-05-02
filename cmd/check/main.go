package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kgalli/image-check/internal/logger"
	"github.com/kgalli/image-check/internal/scheduler"
)

func printKarsten() error {
	fmt.Println("karsten")

	return nil
}

func printDave() error {
	fmt.Println("Dave")

	return nil
}

func main() {
	logger := logger.New()
	scheduler := scheduler.New(logger)

	handleGracefulShutdown(scheduler)
	scheduler.Schedule("print Karsten", printKarsten, time.Second*3)
	scheduler.Schedule("print Dave", printDave, time.Second*1)

	scheduler.Start()
}

func handleGracefulShutdown(scheduler *scheduler.Scheduler) {
	sigCh := signalToWaitFor()
	go func() {
		<-sigCh
		scheduler.Stop()
	}()
}

func signalToWaitFor() chan os.Signal {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	return sigCh
}
