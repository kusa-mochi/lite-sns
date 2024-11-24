package commands

import "log"

type SignupCommand struct {
}

func (c *SignupCommand) Exec() {
	log.Println("signup exec")
}
