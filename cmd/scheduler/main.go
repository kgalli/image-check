package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kgalli/image-check/internal/data"
	"github.com/kgalli/image-check/internal/data/image"
	"github.com/kgalli/image-check/internal/logger"
	"github.com/kgalli/image-check/internal/scheduler"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func addImageDigest(r *image.ImageRepository) func() error {
	return func() error {
		return r.AddDigest("postgres", "latest", randSeq(48))
	}
}

func migrateDB(r *image.ImageRepository) {
	if err := r.Migrate(); err != nil {
		log.Fatal(err)
	}
}

func createImage(r *image.ImageRepository) {
	image := image.BaseImage{
		Name:   "postgres",
		Tag:    "latest",
		Digest: "12345",
	}

	if _, err := r.Create(image); err != nil {
		log.Fatal(err)
	}
}

func main() {
	logger := logger.New()
	conn := connectDB()
	imageRepo := image.NewImageRepository(conn)
	scheduler := scheduler.New(logger)

	migrateDB(imageRepo)
	createImage(imageRepo)
	scheduler.Schedule("create image update", addImageDigest(imageRepo), time.Second*3)
	handleGracefulShutdown(conn, scheduler)

	scheduler.Start()
}

func connectDB() *data.Connection {
	conn, err := data.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

func handleGracefulShutdown(db *data.Connection, scheduler *scheduler.Scheduler) {
	sigCh := signalToWaitFor()
	go func() {
		<-sigCh
		scheduler.Stop()
		db.Close()
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
