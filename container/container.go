package container

import (
	"gorm.io/gorm"
	"library-api/adapter/http/handler"
	"library-api/adapter/repository"
	"library-api/application/usecase"
	postgres "library-api/postgress"
)

type Container struct {
	DB          *gorm.DB
	UserRepo    repository.UserRepository
	UserUsecase *usecase.UserUseCase
	UserHandler *handler.UserHandler
}

func NewContainer() *Container {
	db := postgres.InitDB()

	userRepo := repository.NewUserRepository(db)

	userUseCase := usecase.NewUserUseCase(userRepo)

	userHandler := handler.NewUserHandler(userUseCase)

	return &Container{
		DB:          db,
		UserRepo:    userRepo,
		UserUsecase: userUseCase,
		UserHandler: userHandler,
	}
}
