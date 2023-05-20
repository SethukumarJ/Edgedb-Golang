package models

import "github.com/edgedb/edgedb-go"

type User struct {
	ID           edgedb.UUID `json:"id" edgedb:"id"`
	FirstName    string      `json:"first_name" edgedb:"firstname"`
	LastName     string      `json:"last_name" edgedb:"lastname"`
	DOB          string      `json:"date_of_birth" edgedb:"dob"`
	Email        string      `json:"email" edgedb:"email"`
	Password     string      `json:"password" edgedb:"password"`
	Phone        string      `json:"phone" edgedb:"phone"`
	CV           string      `json:"cv" edgedb:"cv"`
	Verified     bool        `json:"verified" edgedb:"verified"`
	AccessToken  string      `edgedb:"accesstoken"`
	RefreshToken string      `edgedb:"refreshtoken"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUp struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	ID           string `json:"id"`
	AccessToken  string `json:"accesstoken"`
	RefreshToken string `json:"refreshtoken"`
}

type Profile struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	DOB       string `json:"date_of_birth"`
	Phone     string `json:"phone"`
	CV        string `json:"cv"`
	Verified  bool   `json:"-"`
}

type RefreshToken struct {
	RefreshToken string `json:"refreshtoken"`
}
