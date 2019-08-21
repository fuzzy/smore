package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/niklasfasching/go-org/org"
)

func Router(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request: %s %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)
	ServeMarkup(w, r)
}

func highlightCodeBlock(source, lang string) string {
	var w strings.Builder
	l := lexers.Get(lang)
	if l == nil {
		l = lexers.Fallback
	}
	l = chroma.Coalesce(l)
	it, _ := l.Tokenise(nil, source)
	_ = html.New().Format(&w, styles.Get("friendly"), it)
	return `<div class="highlight">` + "\n" + w.String() + "\n" + `</div>`
}

func ServeMarkup(w http.ResponseWriter, r *http.Request) {
	_tmpf := fmt.Sprintf("%s/%s", cfg.Dirs.Templates, cfg.Template)
	_path := fmt.Sprintf("%s%s", cfg.Dirs.Content, r.URL.Path)
	_extt := strings.Split(_path, ".")
	_ext := _extt[len(_extt)-1]

	if isFile(_path) {
		// read in the org document
		bs, err := ioutil.ReadFile(_path)
		check(err)

		// now parse that data accordingly
		switch _ext {
		case "org":
			out := ""

			// setup the template
			title := "SMORE"
			_tmpd, err := ioutil.ReadFile(_tmpf)
			check(err)
			_tmpl, err := template.New(title).Parse(string(_tmpd))
			check(err)

			// read and parse the org document
			orgDoc := org.New().Parse(bytes.NewReader(bs), _path)

			// and setup the html output
			writer := org.NewHTMLWriter()
			writer.HighlightCodeBlock = highlightCodeBlock
			out, err = orgDoc.Write(writer)
			check(err)

			// setup the payload
			_data := struct {
				Path    string
				Title   string
				Payload string
			}{
				Path:    fmt.Sprintf(_path),
				Title:   title,
				Payload: string(out),
			}

			// and actually pump out the output
			err = _tmpl.Execute(w, _data)
			check(err)
		case "md":
			log.Println("It's an Markdown file!")
		default:
			log.Printf("It's a %s file!", _ext)
		}
	}

	return
}
