package web

import (
	"github.com/caarlos0/env/v6"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/rizalfadlila/go-filesvc/config"
	"github.com/rizalfadlila/go-filesvc/handler/web"
	"github.com/rizalfadlila/go-filesvc/repository"
	"github.com/rizalfadlila/go-filesvc/usecase"
	"github.com/spf13/cobra"

	_ "github.com/lib/pq"
)

var (
	serveWEBCmd = &cobra.Command{
		Use:              "web",
		PersistentPreRun: rootPreRun,
		RunE:             runWEB,
	}
)

func rootPreRun(cmd *cobra.Command, args []string) {
	err := godotenv.Load()
	if err != nil {
		log.Info("Read config from OS Environment")
	}
}

func runWEB(cmd *cobra.Command, args []string) error {
	cfg := initConfig()
	opts := web.WebOpts{
		ImageUsecase: usecase.NewImageUsecase(initImageRepository(cfg.Database)),
		Config:       initConfig(),
	}
	web := web.NewWebHandler(opts)
	return web.Run()
}

func ServeWEBCmd() *cobra.Command {
	return serveWEBCmd
}

func initConfig() config.Main {
	cfg := config.Main{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal("error parse config: %w", err)
	}
	return cfg
}

func initPostgres(cfg config.Database) *sqlx.DB {
	db, err := sqlx.Open("postgres", cfg.GetDSN())
	if err != nil {
		log.Fatal("failed to open DB connection: ", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("failed to open DB connection: %w", err)
	}

	return db
}

func initImageRepository(cfg config.Database) repository.ImageRepository {
	return repository.NewImageModule(initPostgres(cfg))
}
