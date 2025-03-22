package main

import "log"

func main() {
	defer hello()
	log.Println("oi, mundo!")
	log.Println("oi, mãe!")
}

func hello() {
	log.Println("meu programa vai encerrar!")
}
