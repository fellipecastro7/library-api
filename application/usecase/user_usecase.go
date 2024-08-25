package usecase

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"library-api/adapter/repository"
	"library-api/domain"
	"time"
)

type UserUseCase struct {
	userrepo repository.UserRepository
}

func NewUserUseCase(userrepo repository.UserRepository) *UserUseCase {
	return &UserUseCase{userrepo: userrepo}
}

func (uc *UserUseCase) CreateUser(userRequest *domain.User) (*domain.User, error) {
	existingUser, err := uc.userrepo.GetByEmail(userRequest.Email)
	//existingUser, err := uc.userRepo.GetByEmail(userRequest.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email já está em uso")
	}

	hashedPassword, err := uc.HashPassword(userRequest.Password)
	if err != nil {
		return nil, errors.New("erro ao gerar hash da senha")
	}

	user := &domain.User{
		Id:           uuid.New(),
		Username:     userRequest.Username,
		Email:        userRequest.Email,
		Password:     userRequest.Password,
		PasswordHash: hashedPassword,
		CreatedAt:    time.Now(),
	}

	err = user.Validate()
	if err != nil {
		return nil, err
	}

	if err = uc.userrepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUseCase) Login(userRequest domain.User) (*domain.User, error) {
	user, err := uc.userrepo.GetByEmail(userRequest.Email)
	if err != nil {
		return nil, errors.New("usuário não encontrado")
	}

	if !uc.CheckPasswordHash(userRequest.Password, user.PasswordHash) {
		return nil, errors.New("senha incorreta")
	}

	return user, nil
}

func (uc *UserUseCase) DeleteUser(id string) error {
	user, err := uc.userrepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("usuário não encontrado")
		}
		return err
	}

	err = uc.userrepo.DeleteByID(user.Id.String())
	if err != nil {
		return errors.New("erro ao deletar o usuário")
	}

	return nil
}

func (uc *UserUseCase) UpdateUser(id string, userRequest domain.User) (*domain.User, error) {
	existingUser, err := uc.userrepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("usuário não encontrado")
		}
		return nil, err
	}

	if userRequest.Email != "" && userRequest.Email != existingUser.Email {
		userWithSameEmail, _ := uc.userrepo.GetByEmail(userRequest.Email)
		if userWithSameEmail != nil {
			return nil, errors.New("email já está em uso")
		}
		existingUser.Email = userRequest.Email
	}

	if userRequest.Username != "" {
		existingUser.Username = userRequest.Username
	}

	if userRequest.Password != "" {
		hashedPassword, err := uc.HashPassword(userRequest.Password)
		if err != nil {
			return nil, errors.New("erro ao gerar hash da senha")
		}
		existingUser.PasswordHash = hashedPassword
	}

	if err := existingUser.Validate(); err != nil {
		return nil, err
	}

	updatedUser, err := uc.userrepo.UpdateUser(existingUser)
	if err != nil {
		return nil, errors.New("erro ao atualizar o usuário")
	}

	return updatedUser, nil
}

func (uc *UserUseCase) GetUserByID(id string) (*domain.User, error) {
	user, err := uc.userrepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("usuário não encontrado")
		}
		return nil, err
	}
	return user, nil
}

func (uc *UserUseCase) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (uc *UserUseCase) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
