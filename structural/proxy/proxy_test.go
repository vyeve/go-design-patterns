package proxy

import (
	"math/rand"
	"testing"
)

func Test_UserListProxy(t *testing.T) {
	someDatabase := UserList{}
	rand.Seed(2342342)
	for i := 0; i < 1000000; i++ {
		n := rand.Int31()
		someDatabase = append(someDatabase, User{ID: n})
	}

	proxy := UserListProxy{
		SomeDatabase:  someDatabase,
		StackCapacity: 2,
		StackCache:    UserList{},
	}

	knownIDs := [3]int32{someDatabase[3].ID, someDatabase[4].ID, someDatabase[5].ID}

	t.Run("FindUser - Empty cache", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}
		if user.ID != knownIDs[0] {
			t.Error("returned user name doesn't match with expected")
		}
		if len(proxy.StackCache) != 1 {
			t.Error("after one successful search in an empty cache, the size of it must be one")
		}
		if proxy.DidLastSearchUsedCache {
			t.Error("no user can be returned from an empty cache")
		}
	})

	t.Run("FindUser - One user, ask for the same user", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}
		if user.ID != knownIDs[0] {
			t.Error("returned user name doestn't match with expected")
		}
		if len(proxy.StackCache) != 1 {
			t.Error("cache must not grow if we asked for an object that is stored on it")
		}
		if !proxy.DidLastSearchUsedCache {
			t.Error("the user should have been returned from the cache")
		}
	})

	user1, err := proxy.FindUser(knownIDs[0])
	if err != nil {
		t.Fatal(err)
	}

	user2, _ := proxy.FindUser(knownIDs[1])
	if proxy.DidLastSearchUsedCache {
		t.Error("the user wasn't stored on the proxy cache yet")
	}

	user3, _ := proxy.FindUser(knownIDs[2])
	if proxy.DidLastSearchUsedCache {
		t.Error("the user wasn't stored on the proxy cache yet")
	}

	for i := 0; i < len(proxy.StackCache); i++ {
		if proxy.StackCache[i].ID == user1.ID {
			t.Error("user that should be gone was found")
		}
	}
	if len(proxy.StackCache) != 2 {
		t.Error("after inserting 3 users the cache should not grow more than to two")
	}

	for _, v := range proxy.StackCache {
		if v != user2 && v != user3 {
			t.Error("a non expected user was found on the cache")
		}
	}
}
