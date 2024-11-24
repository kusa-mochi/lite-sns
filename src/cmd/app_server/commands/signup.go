package commands

import "log"

type SignupCommand struct {
	ResCh chan<- string
}

func (c *SignupCommand) Exec() {
	log.Println("signup exec")
	c.ResCh <- "signup fin"
}
