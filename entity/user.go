package entity

type User struct {
	ID       int
	Username string
	Password string
}

var Users = []User{
	{ID: 1, Username: "Eilang", Password: "Password"},
	{ID: 2, Username: "Whites", Password: "Password"},
	{ID: 3, Username: "Nathan", Password: "Password"},
}
