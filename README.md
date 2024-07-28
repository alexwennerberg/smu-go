# simple markup (smu) in Go

A rewrite of [smu](https://github.com/Gottox/smu) in Go -- a markdown-like or
markdown-lite language.

## Why?
Markdown is an extraordinarily complex standard that is often far more
heavy-weight than some situations may require. Working with and extending
"standards-compliant" markdown can be a headache -- and in many real-world use
cases, not necessary. I also didn't love the APIs of popular Go markdown
libraries. 

This is basically a direct translation from `smu.c`, which means it may look a little
un-Go-like.

smu syntax is very similar to markdown. See https://github.com/Gottox/smu/blob/master/documentation
