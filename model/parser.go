package model

import "os"

const parserHtmlFolder = "parser.html.folder"

type Parser struct {

}

func (p *Parser) ToHTML(file *os.File) string {

	st := Storage{}
	conf := Configuration{}

	st.Save(file, conf.For(parserHtmlFolder))

	return "file content" // TODO implementation
}
