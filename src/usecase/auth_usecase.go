package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/SethukumarJ/trx/src/repositories"
	"golang.org/x/crypto/bcrypt"
)

// authUsecase is the struct for the authentication service
type authUsecase struct {
}

func NewAuthUsecase() *authUsecase {
	return &authUsecase{}
}

// VerifyUser verifies the user credentials
func (c *authUsecase) VerifyUser(ctx context.Context, email string, password string) error {

	userRepo := repositories.NewUserRepository()
	user, err := userRepo.FindByName(ctx, email)

	fmt.Println(user.Email, user.Password, "from verify user")
	if err != nil {
		return errors.New("failed to login. check your email")
	}

	isValidPassword := VerifyPassword(user.Password, []byte(password))
	if !isValidPassword {
		return errors.New("failed to login. check your credential")
	}

	return nil
}

func VerifyPassword(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
