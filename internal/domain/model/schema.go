package core

type User struct {
	Username string
}

type Team struct {
	Name    string
	Members []User
}
