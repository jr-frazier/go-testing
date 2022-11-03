package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2"},
		{"equals zero", 0, false, "0 is not prime, by definition!"},
		{"less than zero", -1, false, "Negative numbers are not prime"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but go false", e.name)
		}
		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}
		if e.msg != msg {
			t.Errorf("%s: expected %s but go %s", e.name, e.msg, msg)
		}
	}
}

func Test_prompt(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	prompt()

	// close writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of our prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	// perform test

	if string(out) != "-> " {
		t.Errorf("incorrect prompt: expected -> but got %s", string(out))
	}
}

func Test_intro(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	intro()

	// close writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of our prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	// perform test

	if !strings.Contains(string(out), "Enter a whole number") {
		t.Errorf("incorrect prompt: expected -> but got %s", string(out))
	}
}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name  string
		input string
		exp   string
	}{
		{name: "empty", input: "", exp: "Please enter a whole number!"},
		{name: "zero", input: "0", exp: "0 is not prime, by definition!"},
		{name: "negative", input: "-1", exp: "Negative numbers are not prime"},
		{name: "not prime", input: "4", exp: "4 is not a prime number because it is divisible by 2"},
		{name: "is prime", input: "7", exp: "7 is a prime number!"},
		{name: "quit", input: "q", exp: ""},
	}

	for _, e := range tests {
		// Simulates user input
		input := strings.NewReader(e.input)
		// Reads user input
		reader := bufio.NewScanner(input)
		// run the function we are testing
		res, _ := checkNumbers(reader)

		// if response does not equal the expected response it will fail
		if !strings.EqualFold(res, e.exp) {
			t.Errorf("%s: expected %s, but got %s", e.name, e.exp, res)
		}
	}
}
