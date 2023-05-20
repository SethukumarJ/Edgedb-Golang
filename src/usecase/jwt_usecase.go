package usecase

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/SethukumarJ/trx/src/models"
	"github.com/golang-jwt/jwt"
)

type jwtUsecase struct {
	UserSecretKey string
}

// GenerateRefreshToken implements interfaces.JWTUsecase
func (j *jwtUsecase) GenerateRefreshToken(userid string, username string, role string) (string, error) {
	claims := &models.SignedDetails{
		UserId:   userid,
		UserName: username,
		Source:   "refreshtoken",
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(150)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(j.UserSecretKey))

	if err != nil {
		log.Println(err)
	}

	return signedToken, err
}

// GenerateRefreshToken implements interfaces.JWTUsecase
func (j *jwtUsecase) GenerateAccessToken(userid string, username string, role string) (string, error) {

	fmt.Println("usernem for accesstoken",username)
	claims := &models.SignedDetails{
		UserId:   userid,
		UserName: username,
		Source:   "accesstoken",
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(j.UserSecretKey))

	if err != nil {
		log.Println(err)
	}

	return signedToken, err
}

// GetTokenFromString implements interfaces.JWTUsecase
func (j *jwtUsecase) GetTokenFromStringUser(signedToken string, claims *models.SignedDetails) (*jwt.Token, error) {
	return jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.UserSecretKey), nil
	})

}

// VerifyToken implements interfaces.JWTUsecase
func (j *jwtUsecase) VerifyTokenUser(signedToken string) (bool, *models.SignedDetails) {
	claims := &models.SignedDetails{}
	token, _ := j.GetTokenFromStringUser(signedToken, claims)

	if token.Valid {
		if e := claims.Valid(); e == nil {
			return true, claims
		}
	}
	return false, claims
}



func NewJWTUsecase() *jwtUsecase {
	return &jwtUsecase{
		UserSecretKey: os.Getenv("USER_KEY"),
	}
}