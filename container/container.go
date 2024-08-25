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
	LoanRepo    repository.LoanRepository
	LoanUseCase *usecase.LoanUseCase
	LoanHandler *handler.LoanHandler
}

func NewContainer() *Container {
	db := postgres.InitDB()

	userRepo := repository.NewUserRepository(db)
	bookRepo := repository.NewBookRepository(db)
	loanRepo := repository.NewLoanRepository(db)

	userUseCase := usecase.NewUserUseCase(userRepo)
	bookUseCase := usecase.NewBookUseCase(bookRepo)
	loanUseCase := usecase.NewLoanUseCase(loanRepo)

	userHandler := handler.NewUserHandler(userUseCase)
	bookHandler := handler.NewBookHandler(bookUseCase)
	loanHandler := handler.NewLoanHandler(loanUseCase)

	return &Container{
		DB:          db,
		UserRepo:    userRepo,
		BookRepo:    bookRepo,
		LoanRepo:    loanRepo,
		UserUsecase: userUseCase,
		BookUseCase: bookUseCase,
		LoanUseCase: loanUseCase,
		UserHandler: userHandler,
		BookHandler: bookHandler,
		LoanHandler: loanHandler,
	}
}
