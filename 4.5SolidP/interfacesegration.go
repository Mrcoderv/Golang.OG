// a api is used to fetch user data from a database

in the previous the userreportry has the varoius method .
type UserRepository interface {
	GetUser(id int)
	CreateUser()
	UpdateUser()
	DeleteUser()
}
// now the user reader interface is used to fetch user data from a database  mwe can replace that with the user repository interface and it will still work as expected.

type UserReader interface {
    GetUser(id int)
}