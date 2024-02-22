package mitocho

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yzaimoglu/mitocho/config"
	"github.com/yzaimoglu/mitocho/utils/crypto"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Mitocho struct {
	Echo  *echo.Echo
	DB    *config.Database
	Debug bool
}

func NewMitocho(db *config.Database) *Mitocho {
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
		DB:    db,
		Debug: config.Debug(),
	}
}

func (mitocho *Mitocho) StartHTTPS() error {
	if config.SSL() {
		if config.Debug() {
			if err := crypto.GenerateKeysFiles(); err != nil {
				return err
			}
			return mitocho.Echo.StartTLS(":"+config.HTTPSPort(), "ssl/cert.pem", "ssl/key.pem")
		}

		cert, key, err := crypto.GenerateKeysBytes()
		if err != nil {
			return err
		}

		return mitocho.Echo.StartTLS(":"+config.HTTPSPort(), cert, key)
	}
	return nil
}

func (mitocho *Mitocho) StartHTTP() error {
	return mitocho.Echo.Start(fmt.Sprintf(":%s", config.HTTPPort()))
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

	if config.SSL() {
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
	}

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
