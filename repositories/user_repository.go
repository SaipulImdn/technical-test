package repositories

import (
	"errors"
	"technical-test/models"

	"gorm.io/gorm"
)

type UserRepository interface {
    CreateUser(user *models.User) error
    GetAllUsers() ([]models.User, error)
    GetUserByID(id uint) (*models.User, error)
    GetUserByUsername(username string) (*models.User, error)
    UpdateUser(user *models.User) error
    DeleteUserByUsername(username string) error
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db}
}

func (r *userRepository) CreateUser(user *models.User) error {
    return r.db.Create(user).Error
}

func (r *userRepository) GetUserByID(id uint) (*models.User, error) {
    var user models.User
    err := r.db.First(&user, id).Error
    return &user, err
}

func (r *userRepository) GetUserByUsername(username string) (*models.User, error) {
    var user models.User
    err := r.db.Where("username = ?", username).First(&user).Error
    return &user, err
}

func (r *userRepository) UpdateUser(user *models.User) error {
	var existingUser models.User
	if err := r.db.Where("(email = ? OR username = ?) AND id != ?", user.Email, user.Username, user.ID).First(&existingUser).Error; err == nil {
		return errors.New("email or username already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err 
	}
	return r.db.Save(user).Error
}

func (r *userRepository) DeleteUserByUsername(username string) error {
    var user models.User
    if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
        return err
    }
    return r.db.Delete(&user).Error
}


func (r *userRepository) GetAllUsers() ([]models.User, error) {
    var users []models.User
    err := r.db.Find(&users).Error
    return users, err
}
