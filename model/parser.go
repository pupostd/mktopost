package model

import (
	"os"
	"gopkg.in/russross/blackfriday.v2"
	"bytes"
	"io"
	"log"
	"path/filepath"
	"strings"
)

const parserHtmlFolder = "parser.html.folder"

type Parser struct {

}

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

	st.Save(nn, out, conf.For(parserHtmlFolder))

	return string(out)
}
