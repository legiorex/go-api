package container

import (
	"go-api/configs"
	"go-api/internal/auth"
	"go-api/internal/link"
	"go-api/internal/user"
	"go-api/pkg/db"
	"go-api/pkg/jwt"
)

type Container struct {
	Config         *configs.Config
	Database       db.DatabaseInterface
	LinkRepository link.LinkRepositoryInterface
	UserRepository user.UserRepositoryInterface
	AuthService    auth.AuthServiceInterface
	JWT            *jwt.JWT
}

func NewContainer() *Container {
	config := configs.LoadConfig()
	database := db.NewDb(config)
	jwt := jwt.NewJWT(config.Auth.Secret)

	linkRepo := link.NewLinkRepository(database)
	userRepo := user.NewUserRepository(database)

	authService := auth.NewAuthService(auth.AuthServiceDeps{
		UserRepository: userRepo,
	})

	return &Container{
		Config:         config,
		Database:       database,
		LinkRepository: linkRepo,
		UserRepository: userRepo,
		AuthService:    authService,
		JWT:            jwt,
	}
}

func (c *Container) GetLinkHandlerDeps() link.LinkHandlerDeps {
	return link.LinkHandlerDeps{
		LinkRepository: c.LinkRepository,
	}
}

func (c *Container) GetAuthHandlerDeps() auth.AuthHandlerDeps {
	return auth.AuthHandlerDeps{
		Config:      c.Config,
		AuthService: c.AuthService,
		JWT:         c.JWT,
	}
}
