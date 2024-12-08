package logdel

const (
	logDeleted  = "user deleted"
	logNotFound = "user not found"
	logAdmin    = "admin deleted"
)

type user struct {
	name   string
	number int
	admin  bool
}

func logAndDelete(users map[string]user, name string) (log string) {
	user, ok := users[name]
	if !ok {
		delete(users, name)
		return logNotFound
	}
	if user.admin {
		delete(users, name)
		return logAdmin
	}
	delete(users, name)
	return logDeleted
}

/* Above code works fine but we can fine tune for good, using defer. 
--Note: The reason could be in future we can add more if statment and have to delete. Bcom tds.
So below code will do the job. 

defer delete(users, name) // this will translate, delete user from users map before logAndDelete function returns, so it will call this delete function right before all returns.
user, ok := users[name]
if !ok {
	delete(users, name)
	return logNotFound
}
if user.admin {
	delete(users, name)
	return logAdmin
}
delete(users, name)
return logDeleted */
