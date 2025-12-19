package main

import (
	"fmt"

	"golang.org/x/crypto/ssh"
	"golang.org/x/net/html"
)

func main() {
	fmt.Println("Project 1 using x/net v0.35.0 and x/crypto (vulnerable)")
	_ = html.Node{}
	_ = ssh.ClientConfig{}
}
