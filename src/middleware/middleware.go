package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/SethukumarJ/trx/src/response"
	"github.com/SethukumarJ/trx/src/usecase"
	"github.com/SethukumarJ/trx/src/utils"

	"github.com/gin-gonic/gin"
)

type Middleware interface {
	AuthorizeJwt() gin.HandlerFunc
	
}
type middleware struct {
}

func NewMiddlewareUser() Middleware {
	return &middleware{
	}

}

func (cr *middleware) AuthorizeJwt() gin.HandlerFunc {
	return (func(c *gin.Context) {

		//getting from header
		autheader := c.Request.Header["Authorization"]
		auth := strings.Join(autheader, " ")
		bearerToken := strings.Split(auth, " ")
		fmt.Printf("\n\ntoken : %v\n\n", autheader)

		if len(bearerToken) != 2 {
			err := errors.New("request does not contain an access token")
			response := response.ErrorResponse("Failed to autheticate jwt", err.Error(), nil)
			c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(*c, response)
			c.Abort()

			return
		}

		authtoken := bearerToken[1]
		jwtUsecase := usecase.NewJWTUsecase()
		ok, claims := jwtUsecase.VerifyTokenUser(authtoken)
		source := fmt.Sprint(claims.Source)
		fmt.Println("///////////////token role", claims.Role)
		if claims.Role != "user" {
			err := errors.New("your role of the token is not valid")
			response := response.ErrorResponse("Error", err.Error(), source)
			c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(*c, response)
			c.Abort()
			return
		}

		if !ok && source == "accesstoken" {
			err := errors.New("your access token is not valid")
			response := response.ErrorResponse("Error", err.Error(), source)
			c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(*c, response)
			c.Abort()
			return
		} else if !ok && source == "refreshtoken" {
			err := errors.New("your refresh token is not valid")
			response := response.ErrorResponse("Error", err.Error(), source)
			c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(*c, response)
			c.Abort()
			return
		} else if ok && source == "refreshtoken" {

			c.Writer.Header().Set("Authorization", authtoken)
			c.Next()
		} else {

			userName := fmt.Sprint(claims.UserName)
			userId := fmt.Sprint(claims.UserId)
			fmt.Println(userId,userName, "from middleware")
			c.Writer.Header().Set("userName", userName)
			c.Writer.Header().Set("user_id", userId)
			c.Next()

		}

	})
}
