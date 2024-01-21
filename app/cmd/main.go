package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"projects/LDmitryLD/repository/app/config"
	"projects/LDmitryLD/repository/app/internal/db"
	"projects/LDmitryLD/repository/app/internal/infrastructure/router"
	"projects/LDmitryLD/repository/app/internal/modules"
	"projects/LDmitryLD/repository/app/internal/storages"
	"syscall"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	confDB := config.NewAppConf().DB
	_, sqlAdapter, err := db.NewSqlDB(confDB)
	if err != nil {
		log.Fatal("Ошибка 1:", err)
	}

	storages := storages.NewStorages(sqlAdapter)

	services := modules.NewServices(storages)

	controllers := modules.NewControllers(services)

	r := router.NewRouter(controllers)

	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Starting server")
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server error: ", err.Error())
		}
	}()

	<-sigChan

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Server stopped")
}
