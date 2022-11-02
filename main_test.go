package main

import "testing"

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
