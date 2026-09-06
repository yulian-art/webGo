package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	server :=  &http.Server{
		Addr: ":8080",
		Handler: mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGABRT,
	)
	defer stop()

	go func ()  {
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			5 * time.Second,
		)

		defer cancel()

		_ = server.Shutdown(shutdownCtx)
	}()

	log.Panicln("server listening on : 8080")
	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed){
		log.Fatal(err)
	}

}