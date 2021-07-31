package greetings

import "testing"
import "regexp"

func TestHelloName(t *testing.T) {
	name := "Tungnt"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Tungnt")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Tungnt") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
