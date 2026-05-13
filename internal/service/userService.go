package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/yihune21/link-shortner/internal/database"
)

type UserService struct {
    q *database.Queries
}

func NewUserService(q *database.Queries) *UserService  {
	return &UserService{q : q}
} 

var (
	ErrEmailExists   = errors.New("email  exists")
	ErrNotFound    = errors.New("not found")
	ErrWrongCredentials = errors.New("Wrong Credentails.")
)

type User struct {
	ID       uuid.UUID   `json:"id"`
	Name    	string `json:"name"`
	Email  		string `json:"email"`
	Password    string `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
}

func mapUser(dbuser database.User) User{
     return User{
        ID: dbuser.ID,
		Name: dbuser.Name,
		Email: dbuser.Email.String,
		Password: dbuser.Password.String,
		CreatedAt: dbuser.CreatedAt,
	 }
}

func (s *UserService)CreateUser(ctx context.Context , name , email,password string )( User , error)  {
	user_email := sql.NullString{}
	if user_email.String != ""{
		user_email.String = email
	}
    
	user_password := sql.NullString{}
	if user_password.String != ""{
		user_password.String = password
	}

	dbuser, err := s.q.CreateUser(ctx ,database.CreateUserParams{
		ID: uuid.New(),
		Name: name,
		Email: user_email,
	    Password:user_password,
		CreatedAt: time.Now().UTC(),
	 } )

	if err != nil {
		return User{} , errors.New(err.Error())
	}
	return mapUser(dbuser) , nil
	


	
}