package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"regexp"
	"time"

	"github.com/SethukumarJ/trx/src/models"
	"github.com/SethukumarJ/trx/src/response"
	"github.com/SethukumarJ/trx/src/usecase"
	"github.com/SethukumarJ/trx/src/utils"
	"github.com/gin-gonic/gin"
)


// @Summary Get user
// @ID Get user
// @Tags Profile
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/get [get]
func GetUser(c *gin.Context) {
	user_id := (c.Writer.Header().Get("user_id"))
	fmt.Println("user_id", user_id)
	if user_id == "" {
		response := response.ErrorResponse("id not found in the header", "try again", nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userUsecase := usecase.NewUserUseCase()
	user, err := userUsecase.FindByID(c, user_id)

	if err != nil {
		response := response.ErrorResponse("Failed to fetch user", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "SUCCESS", user)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}

// UserSignup handles the signup
// @Summary SignUp for Admin
// @ID Admin signup
// @Tags Authentication
// @Produce json
// @param Signup body models.SignUp{} true "user signup"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/signup [post]
// AdminSignup handles the user signup
func UserSignup(c *gin.Context) {

	var newUser models.SignUp
	var User models.User
	var UserResponse models.SignUpResponse
	c.Bind(&newUser)

	User.Email = newUser.Email
	User.Password = newUser.Password
	fmt.Println(User.Email)
	// Validate email
	if !validateEmail(User.Email) {
		response := response.ErrorResponse("Invalid email", "Please provide a valid email address", nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//check username exit or not
	userUsecase := usecase.NewUserUseCase()
	user, err := userUsecase.Save(c, User)

	log.Println(newUser)

	if err != nil {
		response := response.ErrorResponse("Failed to create user", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	user, _ = userUsecase.FindByName(c, user.Email)
	jwtUsecase := usecase.NewJWTUsecase()
	accesstoken, err := jwtUsecase.GenerateAccessToken(user.ID.String(), user.Email, "user")
	if err != nil {
		response := response.ErrorResponse("Failed to generate access token", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(*c, response)
		return
	}
	UserResponse.AccessToken = accesstoken
	refreshtoken, err := jwtUsecase.GenerateRefreshToken(user.ID.String(), user.Email, "user")
	if err != nil {
		response := response.ErrorResponse("Failed to generate refresh token", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(*c, response)
		return
	}
	UserResponse.RefreshToken = refreshtoken

	UserResponse.ID = user.ID.String()

	response := response.SuccessResponse(true, "SUCCESS", UserResponse)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}

// @Summary update profile for users
// @ID User profile
// @Tags Authentication
// @Produce json
// @Security BearerAuth
// @param updateProfile body models.Profile{} true "user profile update"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/profile [patch]
// Update profile handles the user profile updation
func UpdateProfile(c *gin.Context) {

	user_id := c.Writer.Header().Get("user_id")
	fmt.Println("user_id", user_id, "err")
	if user_id == "nil" {
		response := response.ErrorResponse("id not found in the header", "try again", nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	email := c.Writer.Header().Get("userName")
	fmt.Println("email", email)
	if email == "" {
		response := response.ErrorResponse("email not got in the header", "try again", nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var newUser models.Profile
	fmt.Println("user signup")
	//fetching data
	c.Bind(&newUser)

	// Validate phone number
	if !validatePhoneNumber(newUser.Phone) {
		response := response.ErrorResponse("Invalid phone number", "Phone number must have 10 digits", nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Validate date of birth
	if !validateDOB(newUser.DOB) {
		response := response.ErrorResponse("Invalid date of birth", "Age must be 18 years or older", nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userUsecase := usecase.NewUserUseCase()
	//check username exit or not
	user, err := userUsecase.UpdateProfile(c, newUser, user_id)

	log.Println(newUser)

	if err != nil {
		response := response.ErrorResponse("Failed to create user", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	user, err = userUsecase.FindByName(c, email)

	if err != nil {
		response := response.ErrorResponse("Failed to get user", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}
	err = userUsecase.SendSubmissionMail(user.Email)

	fmt.Println(err)

	if err != nil {
		fmt.Println("error wile seding mail")
	}
	user.Password = ""
	response := response.SuccessResponse(true, "SUCCESS", user)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}

// validate phone number
func validatePhoneNumber(phone string) bool {

	match, _ := regexp.MatchString(`^\d{10}$`, phone)
	return match
}

// validate email
func validateEmail(email string) bool {

	_, err := mail.ParseAddress(email)
	return err == nil
}

// validate dob
func validateDOB(dob string) bool {
	// Parse DOB string into time.Time
	dobTime, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return false
	}

	// Calculate age based on current time
	age := time.Since(dobTime).Hours() / 24 / 365

	// Check if age is less than 18
	return age >= 18
}

// UserLogin handles the user login

// @Summary Login for users
// @ID User Login
// @Tags Authentication
// @Produce json
// @Param  UserLogin   body  models.Login{}  true  "userlogin: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/login [post]
// UserLogin handles the user login
func UserLogin(c *gin.Context) {

	var userLogin models.Login

	c.Bind(&userLogin)
	fmt.Println("userLOgin", userLogin.Password, userLogin.Email)
	//verify User details
	authUsecase := usecase.NewAuthUsecase()
	err := authUsecase.VerifyUser(c, userLogin.Email, userLogin.Password)

	if err != nil {
		response := response.ErrorResponse("Failed to login", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(*c, response)
		return
	}

	//fetching user details
	userUsecase := usecase.NewUserUseCase()
	jwtUsecase := usecase.NewJWTUsecase()
	user, _ := userUsecase.FindByName(c, userLogin.Email)
	accesstoken, err := jwtUsecase.GenerateAccessToken(user.ID.String(), user.Email, "user")
	if err != nil {
		response := response.ErrorResponse("Failed to generate access token", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(*c, response)
		return
	}
	user.AccessToken = accesstoken
	refreshtoken, err := jwtUsecase.GenerateRefreshToken(user.ID.String(), user.Email, "user")
	if err != nil {
		response := response.ErrorResponse("Failed to generate refresh token", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(*c, response)
		return
	}
	user.RefreshToken = refreshtoken

	c.Writer.Header().Set("accesstoken", user.AccessToken)
	c.Writer.Header().Set("refreshtoken", user.RefreshToken)
	Tokens := map[string]string{"accesstoken": user.AccessToken, "refreshtoken": user.RefreshToken}
	response := response.SuccessResponse(true, "SUCCESS", Tokens)
	utils.ResponseJSON(*c, response)

	fmt.Println("login function returned successfully")

}

// user refresh token

// @Summary Refresh token for users
// @ID User RefreshToken
// @Tags Authentication
// @Produce json
// @Param  refreshToken   body  models.RefreshToken{}  true  "generate refresh token: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/token-refresh [post]
// UserLogin handles the user login
func UserRefreshToken(c *gin.Context) {

	var token models.RefreshToken

	c.Bind(&token)

	jwtUsecase := usecase.NewJWTUsecase()
	ok, claims := jwtUsecase.VerifyTokenUser(token.RefreshToken)
	if !ok {
		err := errors.New("your refresh token is not valid")
		response := response.ErrorResponse("Error", err.Error(), claims.Source)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(*c, response)
		c.Abort()
		return

	}

	if claims.Source == "accesstoken" {
		response := response.ErrorResponse("cannot use access token to generate refresh token", claims.Source, nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return

	}

	fmt.Println("//////////////////////////////////", claims.UserName)
	accesstoken, err := jwtUsecase.GenerateAccessToken(claims.UserId, claims.UserName, claims.Role)

	if err != nil {
		response := response.ErrorResponse("error generating refresh token", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}
	c.Writer.Header().Set("accesstoken", accesstoken)

	response := response.SuccessResponse(true, "SUCCESS", accesstoken)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}
