package mitocho

import (
	"context"
	"crypto/tls"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yzaimoglu/mitocho/config"
	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Mitocho struct {
	Echo  *echo.Echo
	Debug bool
}

func NewMitocho() *Mitocho {
	e := echo.New()
	e.Use(middleware.HTTPSRedirect())
	e.Use(middleware.HTTPSNonWWWRedirect())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.Secure())
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.CSRF())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Static("/static", "templ/static")

	return &Mitocho{
		Echo:  e,
		Debug: config.Debug(),
	}
}

func (mitocho *Mitocho) StartHTTPS() error {
	if mitocho.Debug {
		return mitocho.Echo.StartTLS(":443", "ssl/cert.pem", "ssl/key.pem")
	}

	autoTLSManager := autocert.Manager{
		HostPolicy: autocert.HostWhitelist(config.Host()),
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache(".cache"),
	}

	tlsServer := http.Server{
		Addr:    ":443",
		Handler: mitocho.Echo,
		TLSConfig: &tls.Config{
			GetCertificate: autoTLSManager.GetCertificate,
			NextProtos:     []string{acme.ALPNProto},
		},
	}

	return tlsServer.ListenAndServeTLS("", "")
}

func (mitocho *Mitocho) StartHTTP() error {
	return mitocho.Echo.Start(":80")
}

func (mitocho *Mitocho) Start() error {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer cancel()

	// Start HTTP Server
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := mitocho.StartHTTP(); !errors.Is(err, http.ErrServerClosed) {
			mitocho.Echo.Logger.Fatalf("Error when starting HTTP server: %w", err)
		} else {
			mitocho.Echo.Logger.Info("HTTP server stopped")
		}
	}()

	// Start HTTPS Server
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := mitocho.StartHTTPS(); !errors.Is(err, http.ErrServerClosed) {
			mitocho.Echo.Logger.Fatalf("Error when starting HTTPS server: %w", err)
		} else {
			mitocho.Echo.Logger.Info("HTTPS server stopped")
		}
	}()

	// Wait for termination signal SIGTERM
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()

		timeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := mitocho.Echo.Server.Shutdown(timeout); !errors.Is(err, http.ErrServerClosed) {
			mitocho.Echo.Logger.Fatalf("Error when shutting down HTTPS server: %v", err)
		} else {
			mitocho.Echo.Logger.Info("HTTPS server shutdown")
		}
	}()

	wg.Wait()
	return nil
}
