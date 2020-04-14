package mypkg

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	if got != "Hello World" {
		t.Errorf("Hello() = %s, want \"Hello World\"", got)
	}
}
