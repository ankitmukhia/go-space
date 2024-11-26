package update

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	if membershipType == "premium" {
		user := User{
			Name: name,
			Membership: Membership{
				Type:             "premium",
				MessageCharLimit: 1000,
			},
		}
		return user
	}

	user := User{
		Name: name,
		Membership: Membership{
			Type:             "standard",
			MessageCharLimit: 100,
		},
	}
	return user
}
