package models

import "github.com/edgedb/edgedb-go"

type User struct {
	ID           edgedb.UUID `edgedb:"id"`
	FirstName    string      `edgedb:"firstname"`
	LastName     string      `edgedb:"lastname"`
	DOB          string      `edgedb:"dob"`
	Email        string      `edgedb:"email"`
	Password     string      `edgedb:"password"`
	Phone        string      `edgedb:"phone"`
	CV           string      `edgedb:"cv"`
	Verified     bool        `edgedb:"verified"`
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
