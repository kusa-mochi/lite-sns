package commands

import "log"

type SignupRequestCommand struct {
	ResCh chan<- string
}

func (c *SignupRequestCommand) Exec() {
	log.Println("signup request exec")
	c.ResCh <- "signup request fin"
}
