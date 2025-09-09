package container

import (
	"go-api/configs"
	"go-api/internal/auth"
	"go-api/internal/link"
	"go-api/internal/user"
	"go-api/pkg/db"
)

type Container struct {
	Config         *configs.Config
	Database       db.DatabaseInterface
	LinkRepository link.LinkRepositoryInterface
	UserRepository user.UserRepositoryInterface
	// AuthService    auth.AuthServiceInterface
}

func NewContainer() *Container {
	config := configs.LoadConfig()
	database := db.NewDb(config)

	linkRepo := link.NewLinkRepository(database)
	userRepo := user.NewUserRepository(database)
	// authService := auth.NewAuthService(config)

	return &Container{
		Config:         config,
		Database:       database,
		LinkRepository: linkRepo,
		UserRepository: userRepo,
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
