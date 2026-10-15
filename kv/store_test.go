package kv

import "testing"

func TestGetMissingKey(t *testing.T) { //ensures empty store doesn't claim to have a key
	s := New()
	if _, ok := s.Get("nope"); ok {
		t.Fatal("expected missing key to return ok=false")
	}
}

func TestPutThenGet(t *testing.T) {
	s := New()

	s.Put("a", "1")
	got, ok := s.Get("a")

	if !ok {
		t.Fatal("expected key \"a\" to exist after Put")
	}
	if got != "1" {
		t.Fatalf("got %q, want %q", got, "1")
	}
}
func TestOverwrite(t *testing.T) {
	s := New()

	s.Put("a", "1")
	s.Put("a", "2")
	got, ok := s.Get("a")

	if !ok {
		t.Fatal("expected key \"a\" to exist after overwrite")
	}
	if got != "2" {
		t.Fatalf("got %q, want %q", got, "2")
	}
}

func TestDeleteThenGet(t *testing.T) {
	s := New()

	s.Put("a", "1")
	if _, ok := s.Get("a"); !ok {
		t.Fatal("setup failed: expected key \"a\" to exist before Delete")
	}

	s.Delete("a")

	if _, ok := s.Get("a"); ok {
		t.Fatal("expected key \"a\" to be gone after Delete")
	}
}

func TestDoublePut(t *testing.T) {
	s := New()

	s.Put("a", "1")
	s.Put("a", "1")
	got, ok := s.Get("a")

	if !ok {
		t.Fatal("expected key \"a\" to exist after double Put")
	}
	if got != "1" {
		t.Fatalf("got %q, want %q", got, "1")
	}
}

func TestEmptyStringValue(t *testing.T) {
	s := New()

	s.Put("a", "")
	got, ok := s.Get("a")

	if !ok {
		t.Fatal("expected key with empty value to exist")
	}
	if got != "" {
		t.Fatalf("got %q, want %q", got, "")
	}
}

func TestDeleteMissingKey(t *testing.T) {
	s := New()

	s.Delete("nope")

	if _, ok := s.Get("nope"); ok {
		t.Fatal("expected missing key to stay missing after Delete")
	}
}
