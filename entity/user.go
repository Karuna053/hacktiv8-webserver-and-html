package entity

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID             int    `form:"id"` // Optional, typically not needed for login
	Username       string `form:"username" binding:"required,min=3,max=150"`
	Email          string `form:"email"`
	HashedPassword string `form:"password" binding:"required,min=3,max=150"`
}

var Users = []User{
	{ID: 1, Username: "Eilang", Email: "Eilang@gmail.com", HashedPassword: HashPassword("Password1")},
	{ID: 2, Username: "Whites", Email: "Whites@gmail.com", HashedPassword: HashPassword("Password2")},
	{ID: 3, Username: "Nathan", Email: "Nathan@gmail.com", HashedPassword: HashPassword("Password3")},
}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
