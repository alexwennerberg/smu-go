package main

import "testing"

func TestSMU(t *testing.T) {
	type Test struct {
		name     string
		in       string
		expected string
		option   *Option
	}

	for _, tc := range []Test{
		{"simple emphasis", "Hello, **World**!", "Hello, <strong>World</strong>!", nil},
		{"emphasis with spaces", "Hello, ** World **!", "Hello, <strong>World</strong>!", nil},
		{"short emphasis", "Hello, *World*!", "Hello, <em>World</em>!", nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := smuRender([]byte(tc.in))
			if string(out) != tc.expected {
				t.Errorf("Failed '%s':\n%s\n\n%s", tc.in, string(out), tc.expected)
			}
		})
	}
}
