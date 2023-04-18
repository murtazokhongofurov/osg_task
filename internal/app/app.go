package app

import (
	"fmt"

	"github.com/osg_task/internal/pkg/config"
	"github.com/osg_task/internal/pkg/db"
	"github.com/osg_task/internal/pkg/logger"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.LogLevel)

	pgxUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	)

	pg, err := db.New(pgxUrl, db.MaxPoolSize(cfg.PgxPoolSize))
	if err != nil {
		l.Fatal(fmt.Errorf("app run postgres.New %w", err))
	}
	defer pg.Close()

}
