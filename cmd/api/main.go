package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	_ "github.com/go-sql-driver/mysql" // New import
	"github.com/y-o-anicca/ggg_clean_arch/internal/config"
	"github.com/y-o-anicca/ggg_clean_arch/internal/http/handler"
	httpRouter "github.com/y-o-anicca/ggg_clean_arch/internal/http/router"
	"github.com/y-o-anicca/ggg_clean_arch/internal/infra/mysql"
	"github.com/y-o-anicca/ggg_clean_arch/internal/usecase"

	"github.com/y-o-anicca/ggg_clean_arch/internal/util/logger"
)

func main() {
	os.Setenv("APP_ENV", "development")
	os.Setenv("HOST", "localhost")
	os.Setenv("PORT", "4000")
	os.Setenv("MYSQL_USERNAME", "web")
	os.Setenv("MYSQL_PASSWORD", "pass")
	os.Setenv("MYSQL_DATABASE", "snippetbox")

	if !config.IsRequiredEnvironments() {
		return
	}

	level := logger.INFO
	if config.IsDevelopment() {
		level = logger.DEBUG
	}
	log := logger.New("test_log", level,
		logger.AddAppendix(logger.NewRuntimeAppendix(logger.WARNING, logger.DefaultCallerSkip)),
	)

	db, err := openDB()
	if err != nil {
		log.Critical(fmt.Sprintf("main: failed to initialize database: %v", err), map[string]interface{}{})
	}
	defer db.Close()

	r := initHandler(db, log)

	listenAddr := config.Host() + ":" + config.Port()
	log.Debug(fmt.Sprintf("main: listenAddr: %v", listenAddr), map[string]interface{}{})
	server := http.Server{
		Addr:         listenAddr,
		Handler:      r,
		ReadTimeout:  time.Duration(config.ReadTimeout()) * time.Second,
		WriteTimeout: time.Duration(config.WriteTimeout()) * time.Second,
	}

	// Listen for syscall signals for process to interrupt/quit
	serverCtx, cancelServerCtx := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		shutdownCtx, cancelShutdownCtx := context.WithTimeout(serverCtx, time.Duration(config.ShutdownWaitTime())*time.Second)
		defer cancelShutdownCtx()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Critical(fmt.Sprintln("main: graceful shutdown timed out.. forcing exit."), map[string]interface{}{})
				os.Exit(1)
			}
		}()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Critical(fmt.Sprintf("main: an error has occurred while shutting down: %v", err), map[string]interface{}{})
			os.Exit(1)
		}
		cancelServerCtx()
	}()

	// Run the server
	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Critical(fmt.Sprintf("main: an error has occurred in server.ListenAndServe: %v", err), map[string]interface{}{})
		os.Exit(1)
	}
	<-serverCtx.Done()
}

func openDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.MysqlDSN())
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func initHandler(db *sql.DB, log *logger.Logger) *chi.Mux {

	repository := mysql.NewClient(log, db)
	usecase := usecase.NewUseCase(log, repository)
	h := handler.NewHandler(usecase, log)
	router := httpRouter.NewRouter(h, log)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Routing(r)
	return r
}
