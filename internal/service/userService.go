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
func mapUsers(dbusers []database.User)[]User{
    users := []User{}
	for _ , user := range dbusers {
		users = append(users,mapUser(user))
	}
    return users
}

func (s *UserService)CreateUser(ctx context.Context , name , email,password string )( User , error)  {
	user_email := sql.NullString{}
	if email != ""{
		user_email.String = email
		user_email.Valid = true
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
func (s *UserService)ListUser(ctx context.Context)( []User , error)  {
	dbUsers , err := s.q.Listusers(ctx)
	if err != nil{
		return []User{},err
	}

	return mapUsers(dbUsers),nil
}
func (s *UserService)GetUserById(ctx context.Context , id uuid.UUID) (User , error)  {
	dbUser , err  := s.q.GetUserById(ctx , id)
	if err != nil {
		return  User{},ErrNotFound
	}
	return mapUser(dbUser) , nil
}

func (s *UserService)GetUserByEmail(ctx context.Context , email string) (User , error)  {
	user_email := sql.NullString{}
	if email != "" {
	   user_email.String = email
	   user_email.Valid = true
	}	
	dbUser , err  := s.q.GetUser(ctx , user_email)
	if err != nil {
		return  User{},ErrNotFound
	}
	return mapUser(dbUser) , nil
}

func (s *UserService)UpdateUserName(ctx context.Context , id uuid.UUID , name string) (User , error){
	dbUser , err := s.q.UpdateUserName(ctx , database.UpdateUserNameParams{
		Name: name,
		ID: id,
	})

	if err != nil {
		return User{} , err
	}

	return mapUser(dbUser) , nil
}

func (s *UserService)DeleteUser(ctx context.Context , id uuid.UUID) error {
	err  := s.q.DeleteUser(ctx , id)
	if err != nil {
		return  ErrNotFound
	}
	return nil
}