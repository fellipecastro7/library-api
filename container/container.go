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
	BookRepo    repository.BookRepository
	BookUseCase *usecase.BookUseCase
	BookHandler *handler.BookHandler
}

func NewContainer() *Container {
	db := postgres.InitDB()

	userRepo := repository.NewUserRepository(db)
	bookRepo := repository.NewBookRepository(db)

	userUseCase := usecase.NewUserUseCase(userRepo)
	bookUseCase := usecase.NewBookUseCase(bookRepo)

	userHandler := handler.NewUserHandler(userUseCase)
	bookHandler := handler.NewBookHandler(bookUseCase)

	return &Container{
		DB:          db,
		UserRepo:    userRepo,
		BookRepo:    bookRepo,
		UserUsecase: userUseCase,
		BookUseCase: bookUseCase,
		UserHandler: userHandler,
		BookHandler: bookHandler,
	}
}
