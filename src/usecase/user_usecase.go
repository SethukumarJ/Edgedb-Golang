package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/SethukumarJ/trx/src/infrastructure"
	"github.com/SethukumarJ/trx/src/models"
	"github.com/SethukumarJ/trx/src/repositories"

	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct{}

// UpdateProfile implements interfaces.UserUseCase
func (c *userUseCase) UpdateProfile(ctx context.Context, profile models.Profile, id string) (models.User, error) {
	var user models.User

	userRepo := repositories.NewUserRepository()
	user, err := userRepo.UpdateProfile(ctx, profile, id)

	return user, err

}

// SendSubmissionMail implements interfaces.userUseCase used for sending mail after form submission
func (c *userUseCase) SendSubmissionMail(email string) error {
	subject := "Form submission"
	message := "Dear User,\n\nThank you for submitting the form. Your submission has been successfully received.\n\nIf you have any further inquiries or require additional information, please feel free to reach out.\n\nBest regards,\nYour Organization"

	msg := []byte(
		"From: TRX machine test <sethukumarj.76@gmail.com>\r\n" +
			"To: " + email + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/plain; charset=UTF-8\r\n\r\n" +
			message)

	// send email with text message
	mailConfig := infrastructure.NewMailConfig()
	err := mailConfig.SendMail(email, msg)

	if err != nil {
		return errors.New("unable to end mail")
	}

	return nil
}

func (c *userUseCase) FindByName(ctx context.Context, name string) (models.User, error) {
	var user models.User

	userRepo := repositories.NewUserRepository()
	user, err := userRepo.FindByName(ctx, name)

	return user, err
}

func (c *userUseCase) FindByID(ctx context.Context, id string) (models.User, error) {

	userRepo := repositories.NewUserRepository()
	user, err := userRepo.FindByID(ctx, id)
	return user, err
}

func (c *userUseCase) Save(ctx context.Context, user models.User) (models.User, error) {

	userRepo := repositories.NewUserRepository()
	_, err := userRepo.FindByName(ctx, user.Email)
	if err == nil {

		return models.User{}, errors.New("email already exists")

	}

	user.Password = HashPassword(user.Password)
	user, err = userRepo.Save(ctx, user)
	fmt.Println(user.Password, "password from usecase")
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// HashPassword hashes the password
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func NewUserUseCase() *userUseCase {
	return &userUseCase{}
}
