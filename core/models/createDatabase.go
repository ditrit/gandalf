package models

type CreateDatabase struct {
	Login    []string
	Password []string
}

func NewCreateDatabase(login, password []string) *CreateDatabase {
	createDatabase := new(CreateDatabase)
	createDatabase.Login = login
	createDatabase.Password = password

	return createDatabase
}
