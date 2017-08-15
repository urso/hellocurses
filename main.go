package main

import (
	"log"

	"github.com/urso/hellocurses/curses"
)

func main() {
	if err := curses.LoadLibs(); err != nil {
		log.Fatal("Loading curses failed:\n", err)
	}

	curses.Initscr()
	defer curses.Endwin()

	curses.Print("Hello World !!!")
	curses.Refresh()
	curses.Getch()
}
