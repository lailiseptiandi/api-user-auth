package services

import (
	"errors"

	"github.com/lailiseptiandi/api-user-auth/app/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	RegisterUser(models.UserRegiserInput) (*models.User, error)
	LoginUser(models.UserLoginInput) (*models.User, error)
	UpdateUser(uint, models.UserUpdateInput) error
	DeleteUser(uint) error
	FindUserById(uint) (*models.User, error)
	GetUser() ([]*models.User, error)
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *userService {
	return &userService{db}
}

func (us *userService) RegisterUser(userInput models.UserRegiserInput) (*models.User, error) {

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashedPassword)

	userModel := models.User{
		Name:       userInput.Name,
		Email:      userInput.Email,
		Password:   string(hashedPassword),
		IsVerified: true,
	}
	newUser, err := userModel.CreateUser(us.db)
	if err != nil {
		return newUser, err
	}
	return newUser, nil

}

func (us *userService) LoginUser(userInput models.UserLoginInput) (*models.User, error) {
	userModel := models.User{}
	user, err := userModel.FindUserByEmal(userInput.Email, us.db)
	if err != nil {
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))
	if err != nil {
		return nil, errors.New("Invalid Password or Email")
	}

	if !user.IsVerified {
		return nil, errors.New("Account not verified")
	}

	return user, nil
}

func (us *userService) UpdateUser(id uint, userInput models.UserUpdateInput) error {

	userModel := models.User{}
	checkUser, err := userModel.FindUserById(id, us.db)
	if err != nil {
		return err
	}

	if checkUser.ID == 0 {
		return err
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashedPassword)
	userModel = models.User{
		Name:     userInput.Name,
		Password: string(hashedPassword),
	}
	err = userModel.UpdateUser(id, us.db)
	if err != nil {
		return err
	}
	return nil
}

func (us *userService) DeleteUser(id uint) error {
	userModel := models.User{}
	checkUser, err := userModel.FindUserById(id, us.db)
	if err != nil {
		return errors.New("user not found")
	}
	if checkUser.ID == 0 {
		return errors.New("user not found")
	}
	err = userModel.DeleteUser(id, us.db)
	if err != nil {
		return err
	}
	return nil
}

func (us *userService) FindUserById(userId uint) (*models.User, error) {
	userModel := models.User{}
	user, err := userModel.FindUserById(userId, us.db)
	if err != nil {
		return user, errors.New("user not found")
	}
	return user, nil
}

func (us *userService) GetUser() ([]*models.User, error) {
	userModel := models.User{}

	user, err := userModel.GetUser(us.db)
	if err != nil {
		return user, err
	}
	return user, nil
}
