package server

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/alexliesenfeld/health"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/sirupsen/logrus"
)

type Server struct {
	db     *sql.DB
	l      *zerolog.Logger
	config ServerConf
	wg     sync.WaitGroup
	health health.Checker
	// entities  entities.Entities
	validator *validator.Validate
	trans     ut.Translator
}

type ServerConf struct {
	Addr string
	Port int
}

func New(logger *logrus.Logger, db *sql.DB, srvConf ServerConf) *Server {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	writer := zerolog.ConsoleWriter{Out: os.Stderr}
	// writer := os.Stderr
	zl := zerolog.New(writer).With().Timestamp().Logger()
	a := &Server{
		l:      &zl,
		db:     db,
		config: srvConf,
		// entities:  entities.New(db, logger),
		validator: validator.New(),
	}

	en := en.New()
	uni := ut.New(en, en)

	a.trans, _ = uni.GetTranslator("en")

	en_translations.RegisterDefaultTranslations(a.validator, a.trans)

	return a
}

func (a *Server) Serve() error {

	shutdownError := make(chan error)
	a.health = health.NewChecker(

		// Set the time-to-live for our cache to 1 second (default).
		health.WithCacheDuration(1*time.Second),

		// Configure a global timeout that will be applied to all checks.
		health.WithTimeout(10*time.Second),

		// A check configuration to see if our database connection is up.
		// The check function will be executed for each HTTP request.
		health.WithCheck(health.Check{
			Name:    "database",      // A unique check name.
			Timeout: 2 * time.Second, // A check specific timeout.
			Check:   a.db.PingContext,
		}),

		// // The following check will be executed periodically every 15 seconds
		// // started with an initial delay of 3 seconds. The check function will NOT
		// // be executed for each HTTP request.
		// health.WithPeriodicCheck(15*time.Second, 3*time.Second, health.Check{
		// 	Name: "search",
		// 	// The check function checks the health of a component. If an error is
		// 	// returned, the component is considered unavailable (or "down").
		// 	// The context contains a deadline according to the configured timeouts.
		// 	Check: func(ctx context.Context) error {
		// 		// return fmt.Errorf("this makes the check fail")
		// 		return nil
		// 	},
		// }),

		// Set a status listener that will be invoked when the health status changes.
		// More powerful hooks are also available (see docs).
		health.WithStatusListener(func(ctx context.Context, state health.CheckerState) {
			a.l.Info().Msg(fmt.Sprintf("health status changed to %s", state.Status))
		}),
	)
	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", a.config.Addr, a.config.Port),
		Handler:      a.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		// ErrorLog:     log.New(a.logger, "", 0),
	}

	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		s := <-quit
		a.l.Info().Str("signal", s.String()).Msg("caught signal")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// shutdownError <- srv.Shutdown(ctx)
		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		a.l.Info().Str("addr", srv.Addr).Msg("completing background tasks")

		a.wg.Wait()
		shutdownError <- nil

	}()

	a.l.Info().Str("addr", srv.Addr).Str("env", a.config.Addr).Msg("starting server")

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	a.l.Info().Str("addr", srv.Addr).Msg("stopped server")

	return nil
}
