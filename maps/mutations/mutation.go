package mutation

// https://www.boot.dev/lessons/2b0807c2-7b0a-4b3e-a16d-26266dcbc460
// visit above link once in a while, for clearity
// check my first solution as well which actually works but has this unuseal behavior where 
// test case some time fails some times pass. due to first if check, which return immediately after first check. which is wrong

// Below is more descriptive
/* The problem:
This is a mistake because the code stops as soon as it finds one user whose name doesn't match the one you're searching for.
But the map may have other users that do match the name, and you never check those.
Also, a map in Go doesn't have any specific order, so you can't be sure that the first user you check is the right one.
*/

import "errors"

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	// checks if key exists
	user, ok := users[name]

	if !ok {
		return false, errors.New("not found")
	}
	
	if user.scheduledForDeletion {
		delete(users, name)
		deleted = true
		return 
	}

	return false, nil
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

/* func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	for _, user := range  users {
		if user.name != name {
			return false, errors.New("not found")
		}

		if user.name == name && !user.scheduledForDeletion {
			deleted = false
			return 
		}

		if user.name == name && user.scheduledForDeletion {
			delete(users, name)
			deleted = true
			return
		}
	}
	return
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
} */
