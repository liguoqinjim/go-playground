package racer

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "http://www.baicu.com"
	fastURL := "http://www.cip.cc"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got '%s',want '%s'", got, want)
	}
}
