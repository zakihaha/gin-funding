package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

type service struct {
	repository Repository
}

// inputan berupa Repository
// return berupa struct service
func NewService(repository Repository) *service {
	// membuat service
	return &service{repository}
}

// ini adalah function yang dibuat di service
// inputan berupa RegisterUserInput
// return berupa User dan error
func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	// service memanggil Repository lalu jalankan function Save
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
