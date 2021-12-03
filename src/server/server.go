package server

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/0B1t322/CP-Rosseti-Back/config"
	"github.com/0B1t322/CP-Rosseti-Back/controllers/auth"
	"github.com/0B1t322/CP-Rosseti-Back/controllers/module"
	"github.com/0B1t322/CP-Rosseti-Back/controllers/role"
	"github.com/0B1t322/CP-Rosseti-Back/controllers/user"
	"github.com/0B1t322/CP-Rosseti-Back/db"
	log "github.com/sirupsen/logrus"
)

func StartServer() error {
	config := config.GetConfig()

	client, err := db.CreateClient(config.DB.URI)
	if err != nil {
		log.WithFields(
			log.Fields{
				"func": "StartServer",
				"err": err,
			},
		).Error("Failed to create db client")
	}

	controllers := &Controllers{
		User: user.New(client),
		Role: role.New(client),
		Auth: auth.New(client, config.Auth.AccessSecret, config.Auth.RefreshSecret),
		Module: module.New(client),
	}

	if err := controllers.Role.CreateRoles(context.Background()); err != nil {
		log.WithFields(
			log.Fields{
				"func": "StartServer",
				"err": err,
			},
		).Error("Failed to create starty up roles")
	}

	if err := controllers.User.CreateUserOnStartUp(); err != nil {
		log.WithFields(
			log.Fields{
				"func": "StartServer",
				"err": err,
			},
		).Error("Failed to create starty up roles")
	}

	r := NewRouter(controllers)

	client.Schema.Create(
		context.Background(),
		schema.WithForeignKeys(true),
		schema.WithGlobalUniqueID(true),
		schema.WithDropIndex(true),
		schema.WithDropColumn(true),
	)

	return r.Run(fmt.Sprintf(":%s", config.App.Port))
}

