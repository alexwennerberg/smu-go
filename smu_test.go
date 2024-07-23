package main

import "testing"

func TestSMU(t *testing.T) {
	type Test struct {
		in       string
		expected string
		option   *Option
	}

	for _, tc := range []Test{
		{"Hello, **World**!", "Hello, <strong>World</strong>!", nil},
		{"Hello, *World*!", "Hello, <em>World</em>!", nil},
	} {
		out := smuRender([]byte(tc.in))
		if string(out) != tc.expected {
			t.Errorf("Unexpected result for test case '%s':\n%s\n\n%s", tc.in, string(out), tc.expected)
		}
	}
}
