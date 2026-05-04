package main

import "testing"

func TestProcessText(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {
            name:     "hex conversion",
            input:    "1E (hex)",
            expected: "30",
        },
        {
            name:     "binary conversion",
            input:    "10 (bin)",
            expected: "2",
        },
        {
            name:     "single case modifiers",
            input:    "hello (up) WORLD (low) brooklyn (cap)",
            expected: "HELLO world Brooklyn",
        },
        {
            name:     "numbered modifier",
            input:    "this is great (up, 2)",
            expected: "this IS GREAT",
        },
        {
            name:     "punctuation spacing",
            input:    "Hello , world !",
            expected: "Hello, world!",
        },
        {
            name:     "single quotes",
            input:    "say ' hello ' to me",
            expected: "say 'hello' to me",
        },
        {
            name:     "article correction",
            input:    "a apple",
            expected: "an apple",
        },
    }

    for _, test := range tests {
        result := processText(test.input)

        if result != test.expected {
            t.Errorf("%s: expected %q, got %q", test.name, test.expected, result)
        }
    }
}
