package app

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
	defaultrolemanager "github.com/casbin/casbin/v2/rbac/default-role-manager"
	"github.com/casbin/casbin/v2/util"
	"github.com/osg_task/api"
	"github.com/osg_task/internal/controller/storage"
	"github.com/osg_task/internal/pkg/config"
	"github.com/osg_task/internal/pkg/db"
	"github.com/osg_task/internal/pkg/logger"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.LogLevel)

	casbinEnforcer, err := casbin.NewEnforcer(cfg.AuthConfigPath, cfg.CsvFilePath)
	if err != nil {
		l.Error("casbin enforcer error", err)
		return
	}

	err = casbinEnforcer.LoadPolicy()
	if err != nil {
		l.Error("casbin error load policy ", err)
		return
	}

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

	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddDomainMatchingFunc("keyMatch", util.KeyMatch)
	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddDomainMatchingFunc("keyMatch3", util.KeyMatch3)

	strg := storage.NewStoragePg(pg)

	server := api.New(&api.Options{
		Conf:           *cfg,
		Logger:         *l,
		Storage:        strg,
		CasbinEnforcer: casbinEnforcer,
	})

	if err := server.Run(cfg.HttpPort); err != nil {
		log.Fatal("failed to run http server ", err)
		panic(err)
	}

	log.Print("Server stopped")

}
