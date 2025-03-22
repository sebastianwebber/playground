package main

import "log"

func main() {
	defer hello()
	log.Println("oi, mundo!")
}

func hello() {
	log.Println("meu programa vai encerrar!")
}
