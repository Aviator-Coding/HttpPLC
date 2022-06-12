package utils

import (
	"log"
	"os"
	"os/signal"

	"github.com/Aviator-Coding/HttpPLC/configs"
	"github.com/gofiber/fiber/v2"
)

// StartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func StartServerWithGracefulShutdown(app *fiber.App) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := app.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Run server.
	if configs.CFG.Server.TlsEnable {
		// Run server.
		if err := app.ListenTLS(configs.CFG.Server.Url, configs.CFG.Server.TlsCertificate, configs.CFG.Server.TlsKeyFile); err != nil {
			log.Fatalf("Oops... Server is not running! Reason: %v", err)
		}

	} else {
		// Run server.
		if err := app.Listen(configs.CFG.Server.Url); err != nil {
			log.Fatalf("Oops... Server is not running! Reason: %v", err)
		}
	}

	<-idleConnsClosed
}

// StartServer func for starting a simple server.
func StartServer(app *fiber.App) {
	if configs.CFG.Server.TlsEnable {
		// Run server.
		if err := app.ListenTLS(configs.CFG.Server.Url, configs.CFG.Server.TlsCertificate, configs.CFG.Server.TlsKeyFile); err != nil {
			log.Fatalf("Oops... Server is not running! Reason: %v", err)
		}

	} else {
		// Run server.
		if err := app.Listen(configs.CFG.Server.Url); err != nil {
			log.Fatalf("Oops... Server is not running! Reason: %v", err)
		}
	}

}
