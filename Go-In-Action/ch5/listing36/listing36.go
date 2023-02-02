package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("sending user email to %s<%s>\n", u.name, u.email)
}

func main() {
	u := user{"Bill", "bill@email.com"}
	sendNotification(&u)
	//sendNotification(u)

	// ./listing36.go:21: cannot use u (type user) as type
	//                     notifier in argument to sendNotification:
	//   user does not implement notifier
	//                          (notify method has pointer receiver)
}

func sendNotification(n notifier) {
	n.notify()
}
