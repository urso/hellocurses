package curses

//go:generate go run ../internal/gen/gen.go -o curses.stub.c -t cursesLinkTable initscr endwin wgetch wrefresh printw

// #include <curses.h>
//
// WINDOW *stdscr;
//
// WINDOW* doInitscr(void) {
//   stdscr = initscr();
//   return stdscr;
// }
//
// extern void* cursesLinkTable;
//
// int printstr(const char *str) { return printw(str); }
import "C"
import "github.com/urso/hellocurses/linktbl"

func LoadLibs() error {
	paths := []string{
		// common linux paths
		"/usr/lib/libncurses.so",
		"/usr/lib/libncurses.so.5",
		"/lib/libncurses.so",
		"/lib/libncurses.so.5",
		"/lib/x86_64-linux-gnu/libncurses.so",
		"/lib/x86_64-linux-gnu/libncurses.so.5",

		// OS X
		"/usr/lib/libncurses.dylib",
	}

	return linktbl.Load(paths, C.cursesLinkTable)
}

func Initscr()         { C.doInitscr() }
func Endwin()          { C.endwin() }
func Getch()           { C.getch() }
func Refresh()         { C.refresh() }
func Print(str string) { C.printstr(C.CString(str)) }
