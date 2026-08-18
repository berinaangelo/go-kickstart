package main

import "fmt"

func initStruct() {
	// verbose declare (w/out explicit type-struct)
	employee := struct {
		name  string
		email string
	}{
		name:  "Fido",
		email: "sample@example.com",
	}

	fmt.Println(employee.name)
	fmt.Println(employee.email)

	// or instantiate as empty
	var employee2 struct {
		name  string
		email string
	}

	employee2.name = "Test User"
	employee2.email = "sample2@example.com"

	fmt.Println(employee2)

	// struct that composed of structs
	type Contact struct {
		Email string
		Phone string
	}

	type User struct {
		ID       int
		Username string
		Contact  // Named field
	}

	user := User{
		ID:       1,
		Username: "kornbip69",
		Contact: Contact{
			Email: "kornbip@example.com",
			Phone: "09696969",
		},
	}

	fmt.Println(user)
}
