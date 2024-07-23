package main

import (
	"bytes"
	"encoding/xml"
	"io"
	"os"
)

// Tag is a struct declaring a transformation that should be made
// on a string
type Tag struct {
	search  string
	process int    // whether to parse sub-items. ie, bold within links (TODO define better)
	before  string // TODO convert?
	after   string
}

// return bytes affected
type Parser = func([]byte, bool, io.Writer) int

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

func doUnderline(buffer []byte, newblock bool, out io.Writer) int {
	return -1
}

// line prefixes are commands that
// a. must precede a line and
// b. apply to a block of lines
// func doLinePrefix(buffer []byte, newblock int, out io.Writer) int {
// 	var p int
// 	var tmp []byte
// 	if newblock {
// 		p = 0
// 	} else if buffer[0] == '\n' {
// 		p += 1
// 	} else {
// 		return false
// 	}
// 	for _, pref := range linePrefix {
// 		if buffer[0] == '\n' {
// 			out.Write('\n') // TODO err handling
// 		}
// 		for bytes.HasPrefix(buffer, []byte(pref.search)) {
// 		}
// 	}
// }

func doSurround(buffer []byte, newblock bool, out io.Writer) int {
out:
	for _, tag := range surround {
		if !bytes.HasPrefix(buffer, []byte(tag.search)) {
			continue
		}
		l := len(tag.search)
		p := l - 1
		var stop int
		for {
			stop = p
			p = bytes.Index(buffer[p:], []byte(tag.search))
			// failed to find matching tag, continue
			if p <= 0 {
				continue out
			} else if buffer[p-1] != '\\' {
				stop = p
				break
			}
		}
		out.Write([]byte(tag.before))

		inside := buffer[l : stop+1]
		// ignore single space around tags
		if inside[0] == ' ' && inside[len(inside)-2] == ' ' {
			inside = inside[1 : len(inside)-2]
			l += 1
		}
		if tag.process > 0 {
			process(inside, false, out)
		} else {
			xml.EscapeText(out, inside)
		}
		out.Write([]byte(tag.after))
		// TODO figure out this conidtion in smu.c
		return stop + l
	}
	return 0
}

// var defaultParsers = []Parser{
// 	doSurround,
// }

var nohtml = false

func process(buffer []byte, newblock bool, out io.Writer) {
	var p int
	for p < len(buffer) {
		if newblock {
			p = bytes.IndexFunc(buffer, func(r rune) bool { return r != '\n' })
			if p == -1 {
				return
			}
		}
		var affected int
		for _, parser := range []Parser{doSurround} {
			affected = parser(buffer[p:], newblock, out)
		}
		p += affected
		if affected == 0 {
			if nohtml {
				xml.EscapeText(out, buffer[p:p+1])
			} else {
				out.Write(buffer[p : p+1])
			}
			p += 1
		}
		// TODO block logic
	}
}

// TODO use
type SMU struct {
	nohtml  bool
	writer  io.Writer
	parsers []Parser
}

func SmuToHTML(md []byte) []byte {
	// var buf bytes.Buffer
	// process(md, 1, &buf)
	// return buf
	return nil
}

func main() {
	process([]byte("Hello, **World**!"), false, os.Stdout)
}
