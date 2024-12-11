package receiver

/* automatically converts c into a pointer (&c) to call the grow() method. */
/* automatically converts e into a pointer (&e) */
/* for refrence : https://www.boot.dev/lessons/b57c7224-fcdd-4dbc-8a2e-9b3ae1a36108 */

func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}


type email struct {
	message string
	fromAddress string
	toAddress string
}
