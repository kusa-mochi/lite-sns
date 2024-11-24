package commands

import "log"

type SignupRequestCommand struct {
}

func (c *SignupRequestCommand) Exec() {
	log.Println("signup request exec")
}
