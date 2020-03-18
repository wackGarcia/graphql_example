package model

type User struct {
	Name 	 string `json:"name"`
	LastName string `json:"last_name"`
}

func (user User) UserGet() User {
	//TODO: Generate connection to database
	user.Name = "Abel"
	user.LastName = "Garcia"

	return user 
}  