package service

import (
	"errors"
	"fmt"
	"math/rand"
	"payment-gwf/entity"
	"payment-gwf/input"
	"payment-gwf/repository"
	"strings"
	"time"

	"github.com/gosimple/slug"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input input.RegisterUserInput) (*entity.User, error)
	Login(input input.LoginInput) (*entity.User, error)
	IsEmailAvailability(input string) (bool, error)
	GetUserByid(ID int) (*entity.User, error)
	DeleteUser(slug string) (*entity.User, error)
	// SaveAvatar(ID int, fileLocation string) (User, error)
	UpdateUser(slugs string, input input.UpdateUserInput) (*entity.User, error)
}

type service struct {
	repository repository.RepositoryUser
}

func NewService(repository repository.RepositoryUser) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input input.RegisterUserInput) (*entity.User, error) {
	user := &entity.User{}

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	slugTitle := strings.ToLower(input.Username)

	mySlug := slug.Make(slugTitle)

	randomNumber := seededRand.Intn(1000000) // Angka acak 0-999999

	user.Slug = fmt.Sprintf("%s-%d", mySlug, randomNumber)

	user.Username = input.Username
	user.Email = input.Email
	user.Password = input.Password
	user.Role = 0
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *service) Login(input input.LoginInput) (*entity.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("user not found that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil

}

func (s *service) DeleteUser(slug string) (*entity.User, error) {
	user, err := s.repository.FindBySlug(slug)
	if err != nil {
		return user, err
	}
	userDel, err := s.repository.Delete(user)

	if err != nil {
		return userDel, err
	}
	return userDel, nil
}

func (s *service) IsEmailAvailability(email string) (bool, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	// Jika user ditemukan, berarti email sudah terdaftar
	if user != nil && user.ID != 0 {
		return false, errors.New("email has been registered")
	}

	return true, nil
}

func (s *service) GetUserByid(ID int) (*entity.User, error) {
	user, err := s.repository.FindById(ID)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("user Not Found With That ID")
	}

	return user, nil

}

func (s *service) UpdateUser(slugs string, input input.UpdateUserInput) (*entity.User, error) {
	user, err := s.repository.FindBySlug(slugs)
	if err != nil {
		return user, err
	}

	oldSlug := user.Slug

	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	slugTitle := strings.ToLower(input.Username)
	mySlug := slug.Make(slugTitle)
	randomNumber := seededRand.Intn(1000000) // Angka acak 0-999999
	user.Slug = fmt.Sprintf("%s-%d", mySlug, randomNumber)

	// Ubah nilai slug kembali ke nilai slug lama untuk mencegah perubahan slug dalam database
	user.Slug = oldSlug

	user.Username = input.Username
	user.Email = input.Email
	user.Password = input.Password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)

	newUser, err := s.repository.Update(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}
