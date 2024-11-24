package commands

import "log"

type SigninCommand struct {
}

func (c *SigninCommand) Exec() {
	log.Println("signin exec")
}
