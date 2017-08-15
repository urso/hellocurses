
/*
 * DO NOT EDIT. This file is generate.
 */

struct tableEntry {
	const char* name;
	void* ptr;
} ;

struct linkTable {
	void* handle;
	struct tableEntry* symbols;
};


void (*_initscr)();
void initscr() { _initscr(); }

void (*_endwin)();
void endwin() { _endwin(); }

void (*_wgetch)();
void wgetch() { _wgetch(); }

void (*_wrefresh)();
void wrefresh() { _wrefresh(); }

void (*_printw)();
void printw() { _printw(); }


struct linkTable* cursesLinkTable = &(struct linkTable){
	.symbols = (struct tableEntry[]) {
		{"initscr", &_initscr},
		{"endwin", &_endwin},
		{"wgetch", &_wgetch},
		{"wrefresh", &_wrefresh},
		{"printw", &_printw},
		{},
	}
};

