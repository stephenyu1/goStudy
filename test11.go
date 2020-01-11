package main

import (
	"fmt"
)

type user struct {
	name string
	email string
}


func (u user) notify() {
	fmt.Printf("Sending user Email to %s<%s>\n", u.name, u.email)
}

func (u *user) changeEmail(newEmail string) {
	u.email = newEmail
}

func changeEmail(u *user, newEmail string) {
	u.email = newEmail
}


func main() {
	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()

	//lisa.changeEmail("new.lisa@email.com")
	//lisa.notify()
	changeEmail(lisa, "new.lisa@email.com")
	lisa.notify()
}

