package kv

import "testing"

func TestGetMissingKey( t* testing.T) { //ensures empty store doesn't claim to have a key
	s:= New()
	if _, ok := s.Get("nope"); ok {
		t.Fatal("expected missing key to return ok=false")
	}
} 