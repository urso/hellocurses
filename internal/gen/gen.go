package main

import (
	"flag"
	"html/template"
	"log"
	"os"
)

var stubTemplate = createTemplate(`
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

{{ range .symbols }}
void (*_{{ . }})();
void {{ . }}() { _{{ . }}(); }
{{ end }}

struct linkTable* {{ .tableName }} = &(struct linkTable){
	.symbols = (struct tableEntry[]) {
		{{- range .symbols }}
		{"{{ . }}", &_{{ . }}},
		{{- end }}
		{},
	}
};

`)

func main() {
	var (
		outFile   = "-"
		tableName = "linkTable"
	)

	flag.StringVar(&outFile, "o", outFile, "output file")
	flag.StringVar(&tableName, "t", tableName, "link table name")
	flag.Parse()
	symbols := flag.Args()

	out := os.Stdout
	if outFile != "-" {
		var err error
		out, err = os.Create(outFile)
		if err != nil {
			log.Fatal("Failed to create output file:", err)
		}

		defer func() {
			err := out.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	stubTemplate.Execute(out, map[string]interface{}{
		"tableName": tableName,
		"symbols":   symbols,
	})
}

func createTemplate(in string) *template.Template {
	return template.Must(template.New("").Parse(in))
}
