/*
The Proxy pattern usually wraps an object to hide some of its characteristics.These characteistics
could be the fact that it is a remote object (remote proxy), a very heavy object such as a very big
image or the dump of e terabyte database (virtual proxy), or a restricted access object (protection
proxy).
The Proxy uses the same interface as the type it wraps.
*/

package proxy

import (
	"fmt"
)

type UserListProxy struct {
	SomeDatabase           UserList
	StackCache             UserList
	StackCapacity          int
	DidLastSearchUsedCache bool
}

func (p *UserListProxy) FindUser(id int32) (User, error) {
	user, err := p.StackCache.FindUser(id)
	if err == nil {
		fmt.Println("Returning user from cache...")
		p.DidLastSearchUsedCache = true
		return user, nil
	}
	user, err = p.SomeDatabase.FindUser(id)
	if err != nil {
		return User{}, err
	}
	p.addUserToStack(user)
	fmt.Println("Returning user from database...")
	p.DidLastSearchUsedCache = false
	return user, nil
}

func (p *UserListProxy) addUserToStack(user User) {
	if len(p.StackCache) >= p.StackCapacity {
		p.StackCache = append(p.StackCache[1:], user)
	} else {
		p.StackCache.addUser(user)
	}
}
