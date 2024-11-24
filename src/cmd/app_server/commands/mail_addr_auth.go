package commands

import "log"

type MailAddrAuthCommand struct {
	ResCh chan<- string
}

func (c *MailAddrAuthCommand) Exec() {
	log.Println("mail addr auth exec")
	c.ResCh <- "mail addr auth fin"
}
