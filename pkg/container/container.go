package container

import (
	"go-api/configs"
	"go-api/internal/auth"
	"go-api/internal/link"
	"go-api/pkg/db"
)

type Container struct {
	Config         *configs.Config
	Database       db.DatabaseInterface
	LinkRepository link.LinkRepositoryInterface
	// AuthService    auth.AuthServiceInterface
}

func NewContainer() *Container {
	config := configs.LoadConfig()
	database := db.NewDb(config)

	linkRepo := link.NewLinkRepository(database)
	// authService := auth.NewAuthService(config)

	return &Container{
		Config:         config,
		Database:       database,
		LinkRepository: linkRepo,
		// AuthService:    authService,
	}
}

func (c *Container) GetLinkHandlerDeps() link.LinkHandlerDeps {
	return link.LinkHandlerDeps{
		LinkRepository: c.LinkRepository,
	}
}

func (c *Container) GetAuthHandlerDeps() auth.AuthHandlerDeps {
	return auth.AuthHandlerDeps{
		Config: c.Config,
	}
}
