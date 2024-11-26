package message

func (u User) SendMessage(message string, messageLength int) (string, bool) {
	// users might send shorter messages within the allowed character limit, so that's why <= conditioea
	// like X has within limit max: 100 but i can write only 5 which should be true
	if messageLength <= u.MessageCharLimit {
		return message, true
	}
	return "", false
}

// don't touch below  this line

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	membership := Membership{Type: membershipType}
	if membershipType == "premium" {
		membership.MessageCharLimit = 1000
	} else {
		membership.Type = "standard"
		membership.MessageCharLimit = 100
	}
	return User{Name: name, Membership: membership}
}
