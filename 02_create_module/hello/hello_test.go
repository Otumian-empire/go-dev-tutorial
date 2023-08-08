package hello

import (
	"fmt"
	"regexp"
	"testing"
)

func TestSayHelloPassingAStringValue(t *testing.T) {
	name := "Daniel"
	want := regexp.MustCompile(`\b` + name + `\b`)

	value, err := SayHello(name)

	if !want.MatchString(name) {
		t.Fatalf(`SayHello("%v") %q, %v want match for %#q, nil`, name, value, err, want)
	}
}

func TestSayHelloPassingAnEmptyString(t *testing.T) {
	value, err := SayHello("")
	fmt.Println(value)
	fmt.Println(err)

	if value != "" || err == nil {
		t.Fatalf(`SayHello("") = %q, %v, want "", error`, value, err.Error())
	}
}
