package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// anonymous struct embedding - similar to class inheritance
// anonymous struct embedding is used here to get access of 'firstName', 'lastName' etc fields of Admin user in the 'structs.go' file
type Admin struct {
	email    string
	password string
	User
}

func (u *User) OutputUserDetails() {
	// fmt.Println((*u).firstName, (*u).lastName, (*u).birthDate)
	fmt.Println(u.firstName, u.lastName, u.birthDate) // shortcut allowed by GO
}

// mutation method
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// 'admin' constructor function
// not using pointer return feature this time just to have variety => copy of Admin struct will be created by this constructor function
func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}
}

// 'user' constructor function
func New(firstName, lastName, birthdate string) (*User, error) {
	// constructor validation (if needed)
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name and birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}, nil
}

// ------------------

// named struct embedding - similar to class inheritance
// in this case all Admin fields, including the User fields, will be unexposed in the 'structs.go' file
// type Admin struct {
// 	email    string
// 	password string
// 	user     User
// }

// 'admin' constructor function
// not using pointer return feature this time just to have variety => copy of Admin struct will be created by this constructor function
// func NewAdmin(email, password string) Admin {
// 	return Admin{
// 		email:    email,
// 		password: password,
// 		user: User{
// 			firstName: "ADMIN",
// 			lastName:  "ADMIN",
// 			birthDate: "---",
// 			createdAt: time.Now(),
// 		},
// 	}
// }

// ------------------

// named struct embedding - similar to class inheritance
// in this case all Admin fields, including the User fields, will be unexposed in the 'structs.go' file
// type Admin struct {
// 	email    string
// 	password string
// 	User     User
// }

// 'admin' constructor function
// not using pointer return feature this time just to have variety => copy of Admin struct will be created by this constructor function
// func NewAdmin(email, password string) Admin {
// 	return Admin{
// 		email:    email,
// 		password: password,
// 		User: User{
// 			firstName: "ADMIN",
// 			lastName:  "ADMIN",
// 			birthDate: "---",
// 			createdAt: time.Now(),
// 		},
// 	}
// }

// In above, even if we capitalize the 'user' field to 'User' we have to use admin.User.OutputUserDetails() to access it in the 'structs.go' file which is annoying, so the best solution is to write as:

// type Admin struct {
// 	email    string
// 	password string
// 	User
// }
