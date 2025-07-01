package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/handarudwiki/mkp-tht/config"
	"github.com/handarudwiki/mkp-tht/database/connections"
	"github.com/handarudwiki/mkp-tht/internal/routes"
)

func main() {
	cfg := config.GetConfig()

	db, err := connections.GetDatabaseConnection(cfg)

	if err != nil {
		panic(err)
	}

	app := gin.New()

	routes.InitUser(db, cfg.JWT, app)

	app.Run(fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port))
}
