package container

import (
	"go-api/configs"
	"go-api/internal/auth"
	"go-api/internal/link"
	"go-api/internal/stat"
	"go-api/internal/user"
	"go-api/pkg/db"
	"go-api/pkg/event"
	"go-api/pkg/jwt"
)

type Container struct {
	Config         *configs.Config
	Database       db.DatabaseInterface
	LinkRepository link.LinkRepositoryInterface
	UserRepository user.UserRepositoryInterface
	StatRepository stat.StatRepositoryInterface
	AuthService    auth.AuthServiceInterface
	JWT            *jwt.JWT
	EventBus       event.EventBusInterface
}

func NewContainer() *Container {
	config := configs.LoadConfig()
	database := db.NewDb(config)
	jwt := jwt.NewJWT(config.Auth.Secret)

	eventBus := event.NewEventBus()

	linkRepo := link.NewLinkRepository(database)
	userRepo := user.NewUserRepository(database)
	statRepo := stat.NewStatRepository(database)

	authService := auth.NewAuthService(auth.AuthServiceDeps{
		UserRepository: userRepo,
	})

	statService := stat.NewServiceStat(&stat.ServiceStatDeps{
		EventBus:       eventBus,
		StatRepository: statRepo,
	})

	go statService.AddClick()

	return &Container{
		Config:         config,
		Database:       database,
		LinkRepository: linkRepo,
		UserRepository: userRepo,
		StatRepository: statRepo,
		AuthService:    authService,
		JWT:            jwt,
		EventBus:       eventBus,
	}
}

func (c *Container) GetLinkHandlerDeps() link.LinkHandlerDeps {
	return link.LinkHandlerDeps{
		LinkRepository: c.LinkRepository,
		EventBus:       c.EventBus,
		JWT:            c.JWT,
	}
}

func (c *Container) GetAuthHandlerDeps() auth.AuthHandlerDeps {
	return auth.AuthHandlerDeps{
		Config:      c.Config,
		AuthService: c.AuthService,
		JWT:         c.JWT,
	}
}
