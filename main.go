package main

import "GolangWsLibrary/client"

func main() {
	go client.StartConnect()
	select {}
}
