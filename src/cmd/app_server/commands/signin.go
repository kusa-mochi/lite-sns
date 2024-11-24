package commands

import "log"

type SigninCommand struct {
	ResCh chan<- string
}

func (c *SigninCommand) Exec() {
	log.Println("signin exec")
	c.ResCh <- "signin fin"
}
