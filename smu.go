package main

import "fmt"

// Tag is a struct declaring a transformation that should be made
// on a string
type Tag struct {
	search  string
	process int // ?
	before  string
	after   string
}

type Parser = func(string, int) string

var linePrefix = []Tag{
	{"    ", 0, "<pre><code>", "\n</code></pre>"},
	{"\t", 0, "<pre><code>", "\n</code></pre>"},
	{">", 2, "<blockquote>", "</blockquote>"},
	{"###### ", 1, "<h6>", "</h6>"},
	{"##### ", 1, "<h5>", "</h5>"},
	{"#### ", 1, "<h4>", "</h4>"},
	{"### ", 1, "<h3>", "</h3>"},
	{"## ", 1, "<h2>", "</h2>"},
	{"# ", 1, "<h1>", "</h1>"},
	{"- - -\n", 1, "<hr />", ""},
	{"---\n", 1, "<hr />", ""},
}

var underline = []Tag{
	{"=", 1, "<h1>", "</h1>\n"},
	{"-", 1, "<h2>", "</h2>\n"},
}

var surround = []Tag{
	{"``", 0, "<code>", "</code>"},
	{"`", 0, "<code>", "</code>"},
	{"___", 1, "<strong><em>", "</em></strong>"},
	{"***", 1, "<strong><em>", "</em></strong>"},
	{"__", 1, "<strong>", "</strong>"},
	{"**", 1, "<strong>", "</strong>"},
	{"_", 1, "<em>", "</em>"},
	{"*", 1, "<em>", "</em>"},
}

var replace = [][]string{
	{"\\\\", "\\"}, // must always remain first
	{"\\`", "`"},
	{"\\*", "*"},
	{"\\_", "_"},
	{"\\{", "{"},
	{"\\}", "}"},
	{"\\[", "["},
	{"\\]", "]"},
	{"\\(", "("},
	{"\\)", ")"},
	{"\\#", "#"},
	{"\\+", "+"},
	{"\\-", "-"},
	{"\\.", "."},
	{"\\!", "!"},
}

var insert = [][]string{
	{"  \n", "<br />"},
}

func doUnderline(s string) {
}

var defaultParsers = []Parser{}

// MD is...
type MD struct {
	data     []byte
	newblock bool
	parsers
}

func NewMD() MD {
	MD{}
}
func (md MD) process() {
}

func main() {
	fmt.Println("vim-go")
}
