package commands

import "log"

type MailAddrAuthCommand struct {
}

func (c *MailAddrAuthCommand) Exec() {
	log.Println("mail addr auth exec")
}
