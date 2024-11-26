package method

type authenticationInfo struct {
	username string
	password string
}

func (auth authenticationInfo) getBasicAuth() string {
	prefix := "Authorization: Basic "
	return prefix + auth.username + ":" + auth.password
}
