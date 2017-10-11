package model

import (
	"bytes"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/russross/blackfriday.v2"
)

const parserHTMLFolder = "parser.html.folder"

// Parser A class to transform a requested resource/file in a new format.
type Parser struct {
}

// ToHTML transforms the .md file given to html.
func (p *Parser) ToHTML(file *os.File) string {
	st := Storage{}
	conf := Configuration{}

	buf := bytes.NewBuffer(nil)
	_, err := io.Copy(buf, file)

	if err != nil {
		log.Fatalf("Could not transform markdown to html.")
		return ""
	}

	out := blackfriday.Run(buf.Bytes())
	nn := filepath.Base(file.Name())
	nn = (strings.Split(nn, "."))[0] + ".html"

	st.Save(nn, out, conf.For(parserHTMLFolder))

	return string(out)
}
