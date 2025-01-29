package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {

	name := "Rie"
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello(name)

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello(name) = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

func TestHellos(t *testing.T) {

	names := []string{"Rie", "Kayra", "Oytun"}
	want := regexp.MustCompile((`(?i)(rie|kayra|oytun)`))

	msgs, err := Hellos(names)
	for _, msg := range msgs {
		if !want.MatchString(msg) || err != nil {
			t.Fatalf(`Hellos(name) = %q, %v, want match for %#q, nil`, msg, err, want)
		}
	}

}
