package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"log"
	"net/http"
)

func main() {
	// Init a HTTP server
	server := &http.Server{
		Addr: ":8080",
	}

	// Controller for the HTTP request
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		/// Simulate a long running process
		for i := 0; i < 10; i++ {
			fmt.Fprintf(w, "Hello World %d\n", i)
			time.Sleep(time.Second)
		}
		fmt.Fprintf(w, "All done!\n")
	})

	// Arrancar el servidor HTTP en una goroutine
	go func() {
		log.Println("Servidor HTTP iniciado en http://localhost:8080")
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Esperar la señal de interrupción (Ctrl+C)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt

	// Detener el servidor HTTP
	log.Println("Deteniendo servidor HTTP...")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}
	log.Println("Servidor HTTP detenido.")
}
