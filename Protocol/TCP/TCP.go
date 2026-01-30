package main

import (
	client "project/memo/Protocol/TCP/client"
)

func main() {
	client.Start()
	server.start()
}
