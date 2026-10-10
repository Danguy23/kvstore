package kv

import "testing"

func TestGetMissingKey( t* testing.T) {
	s:= New()
	if _, ok := s.Get("nope"); ok {
		t.Fatal("expected missing key to return ok=false")
	}
}