package user

import (
	"fmt"
)

type User struct {
	ID   int
	Name string
	Role string
}

type Users []User

func (u Users) Exists(id int) bool {

}

func (u Users) FindByName(name string) (User, error) {

}
