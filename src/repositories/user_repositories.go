package repositories

import (
	"context"
	"fmt"

	"github.com/edgedb/edgedb-go"
	"github.com/SethukumarJ/trx/src/infrastructure"
	"github.com/SethukumarJ/trx/src/models"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (c *UserRepository)  Save(ctx context.Context, user models.User) ( models.User,error) {
	client := infrastructure.NewDBClient(ctx)
	defer client.Close()
	var inserted struct{ id edgedb.UUID }
	query := `INSERT User {
		email := <str>$0, 
		password := <str>$1
	}`

	username := user.Email
	password := user.Password
	err := client.QuerySingle(ctx, query, &inserted, username, password)
	if err != nil {
		fmt.Println("save error")
		return models.User{}, err
	}
	fmt.Println(user.Email,user.ID, user.Password, "from repo save function")
	return user, nil
}



func (c *UserRepository) FindByID(ctx context.Context, id string) (models.User, error) {
	client := infrastructure.NewDBClient(ctx)
	query := "SELECT User filter .id=<uuid>$0"
	var user models.User
	err := client.Query(ctx, query, &user, id)
	
	if err != nil 	{
		return models.User{},err
	}

	return user,nil
}




func (c *UserRepository) FindByName(ctx context.Context, name string) (models.User,error) {
	
	
	var user models.User
	client := infrastructure.NewDBClient(ctx)
	query := "SELECT User FILTER .email = <str>$0"
	err := client.QuerySingle(ctx, query, &user, name)
	fmt.Print("user from find naem", user.Email, "email",user.Password, "password")
	if err != nil {
		fmt.Println("find by name errror", err)
		return models.User{},err
	}

	return user,nil
}

func (c *UserRepository) GetAll(ctx context.Context) (users []models.User, err error) {
	client := infrastructure.NewDBClient(ctx)
	query := "SELECT User{id, username, password}"
	err = client.Query(ctx, query, &users)
	if err != nil {
		return
	}
	return
}

func (c *UserRepository) UpdateProfile(ctx context.Context, profile models.Profile, id string) (user models.User, err error) {
	client := infrastructure.NewDBClient(ctx)

	// Retrieve the user from the database based on the ID
	query := "SELECT User FILTER .id = <uuid>$0"
	err = client.QuerySingle(ctx, query, &user, id)
	if err != nil {
		fmt.Println("update profile error: failed to retrieve user", err)
		return user, err
	}

	// Update the user's profile with the provided data
	user.FirstName = profile.FirstName
	user.LastName = profile.LastName
	user.DOB = profile.DOB
	user.Phone = profile.Phone
	user.CV = profile.CV
	user.Verified = true

	// Save the updated user to the database
	query = "UPDATE $0 SET {firstname := <str>$1, lastname := <str>$2, dob := <str>$3, phone := <str>$4, cv := <str>$5, verified := true} WHERE .id = <uuid>$6"
	err = client.Execute(ctx, query, "User", user.FirstName, user.LastName, user.DOB, user.Phone, user.CV, id)
	if err != nil {
		fmt.Println("update profile error: failed to update user", err)
		return user, err
	}

	return user, nil
}



