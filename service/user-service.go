package service

import (
	"log"

	"go1/dto"
	"go1/entity"
	"go1/repository"

	"github.com/mashingan/smapping"
)

//UserService is a contract.....
type UserService interface {
	Update(user_id string, user dto.UserDTO) entity.User
	Profile(userID string) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

//NewUserService creates a new instance of UserService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(userID string, user dto.UserDTO) entity.User {
	userToUpdate := service.userRepository.ProfileUser(userID)
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) Profile(userID string) entity.User {
	return service.userRepository.ProfileUser(userID)
}
