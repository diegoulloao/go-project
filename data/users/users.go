package users

// Structs
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Data types
type Users []User

// Collections
var List = Users{
	{
		ID:    0,
		Name:  "diego",
		Email: "diego@gmail.com",
	},
	{
		ID:    1,
		Name:  "hernán",
		Email: "hernan@gmail.com",
	},
	{
		ID:    2,
		Name:  "felipe",
		Email: "felipe@gmail.com",
	},
}
