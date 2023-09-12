package service

import (
	"crypto/sha1"
	"fmt"

	todo "github.com/khlkvmgmd/to-do-app"
	"github.com/khlkvmgmd/to-do-app/pkg/repository"
)

const salt = "bgyu8yytg78yutgh78yucvbnuj"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthSevice(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser (user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}