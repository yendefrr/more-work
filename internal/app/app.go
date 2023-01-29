package app

import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/yendefrr/more-work/internal/config"
	"github.com/yendefrr/more-work/internal/repository"
	"github.com/yendefrr/more-work/internal/server"
	"github.com/yendefrr/more-work/internal/service"
	"github.com/yendefrr/more-work/internal/transport/http/handler"

	_ "github.com/go-sql-driver/mysql"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatalf("error occured while init config: %s", err.Error())
		return
	}

	log.Info(cfg.MySQL.DatabaseName)

	db, err := newDB(cfg.MySQL)
	if err != nil {
		log.Fatalf("error occured while connecting database: %s", err.Error())
		return
	}

	repos := repository.NewRepositories(db)
	jobsService := service.NewJobsService(repos.Jobs)
	employersService := service.NewEmployersService(repos.Employers)
	seekersService := service.NewSeekersService(repos.Seekers)

	handler := handler.NewHandler(jobsService, employersService, seekersService)

	srv := server.NewServer(cfg, handler.Init())

	srv.Run()

	// go func() {
	// 	if err := srv.Run(); err != nil {
	// 		log.Fatalf("error occured while running server: %s", err.Error())
	// 	}
	// }()

	log.Info("Server started")
}

func newDB(cfg config.MySQLConfig) (*sql.DB, error) {
	log.Info(cfg.Password)
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", cfg.User, cfg.Password, cfg.URI, cfg.DatabaseName))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
