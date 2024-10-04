package services

import (
	"errors"
	"technical-test/models"
	"technical-test/repositories"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)


var jwtSecret = []byte("your_secret_key")

type UserService interface {
    Register(user *models.User) error
    Login(username, password string) (string, error)
    GetUserByUsername(username string) (*models.User, error)
    GetAllUsers() ([]models.User, error)
    UpdateUser(username string, newEmail string, newUsername string) error
    DeleteUserByUsername(username string) error
}

type userService struct {
    userRepository repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
    return &userService{userRepo}
}

func (s *userService) Register(user *models.User) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)
    return s.userRepository.CreateUser(user)
}

func (s *userService) Login(username, password string) (string, error) {
    user, err := s.userRepository.GetUserByUsername(username)
    if err != nil {
        return "", err
    }
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return "", errors.New("invalid username or password")
    }
    token, err := GenerateJWT(user.Username)
    if err != nil {
        return "", err
    }
    return token, nil
}

func GenerateJWT(username string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)

    claims := &jwt.RegisteredClaims{
        Subject:   username,
        ExpiresAt: jwt.NewNumericDate(expirationTime),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}


func (s *userService) GetUserByUsername(username string) (*models.User, error) {
    return s.userRepository.GetUserByUsername(username)
}


func (s *userService) GetAllUsers() ([]models.User, error) {
    return s.userRepository.GetAllUsers()
}

func (s *userService) UpdateUser(username string, newEmail string, newUsername string) error {
    user, err := s.userRepository.GetUserByUsername(username)
    if err != nil {
        return err
    }

    user.Email = newEmail
    user.Username = newUsername
    return s.userRepository.UpdateUser(user)
}

func (s *userService) DeleteUserByUsername(username string) error {
    return s.userRepository.DeleteUserByUsername(username)
}
