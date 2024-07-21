package main

import (
	"bytes"
	"fmt"
	"io"
)

// Tag is a struct declaring a transformation that should be made
// on a string
type Tag struct {
	search  string
	process int    // whether to parse sub-items. ie, bold within links (TODO define better)
	before  string // TODO convert?
	after   string
}

// return whether affected
type Parser = func([]byte, int, io.Writer) bool

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

// can exist within lines, and if they do they can potentially cross lines.
// basically, non-line-oriented
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

func doUnderline(buffer []byte, newblock int, out io.Writer) {
}

// line prefixes are commands that
// a. must precede a line and
// b. apply to a block of lines
func doLinePrefix(buffer []byte, newblock int, out io.Writer) bool {
	var p int
	var tmp []byte
	if newblock {
		p = 0
	} else if buffer[0] == '\n' {
		p += 1
	} else {
		return false
	}
	for _, pref := range linePrefix {
		if buffer[0] == '\n' {
			out.Write('\n') // TODO err handling
		}
		for bytes.HasPrefix(buffer, []byte(pref.search)) {
		}
	}
}

func doSurround(buffer []byte, newblock int, out io.Writer) bool {
	for _, pref := range surround {
	}
}

var defaultParsers = []Parser{
	doUnderline,
}

func process(buffer []byte, newblock int, out io.Writer) {

}

func SmuToHTML(md []byte) []byte {
	var buf bytes.Buffer
	process(md, 1, &buf)
	return buf
}

func main() {
	fmt.Println("vim-go")
}
