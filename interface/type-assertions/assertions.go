package asseratation

func getExpenseReport(e expense) (string, float64) {
	// i is instence of email struct, e. instence of expese interface
	i, ok := e.(email)

	if ok {
		return i.toAddress, i.cost()
	}

	// s is instence of sms struct, e. instence of expese interface
	s, ok := e.(sms)

	if ok {
		return s.toPhoneNumber, s.cost()
	}

	return "", 0.0
}

// don't touch below this line

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}
