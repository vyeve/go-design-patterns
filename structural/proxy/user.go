package proxy

import "fmt"

type UserFinder interface {
	FindUser(id int32) (User, error)
}

type User struct {
	ID int32
}

type UserList []User

func (ul *UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(*ul); i++ {
		if (*ul)[i].ID == id {
			return (*ul)[i], nil
		}
	}
	return User{}, fmt.Errorf("user %d could not be found\n", id)
}

func (ul *UserList) addUser(newUser User) {
	*ul = append(*ul, newUser)
}
