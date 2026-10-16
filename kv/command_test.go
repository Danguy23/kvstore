package kv

import (
	"reflect"
	"testing"
)

// mustApply runs Apply and fails the test if it returns an error.
func mustApply(t *testing.T, s *Store, cmd Command) Result {
	t.Helper()
	res, err := s.Apply(cmd)
	if err != nil {
		t.Fatalf("Apply(%+v) returned error: %v", cmd, err)
	}
	return res
}

func TestApplyPutThenGet(t *testing.T) {
	s := New()

	mustApply(t, s, Command{Op: OpPut, Key: "a", Value: "1"})
	res := mustApply(t, s, Command{Op: OpGet, Key: "a"})

	if !res.Found {
		t.Fatal("expected key \"a\" to be found after Put")
	}
	if res.Value != "1" {
		t.Fatalf("got %q, want %q", res.Value, "1")
	}
}

func TestApplyGetMissingKey(t *testing.T) {
	s := New()

	res := mustApply(t, s, Command{Op: OpGet, Key: "nope"})

	if res.Found {
		t.Fatal("expected missing key to return Found=false")
	}
}

func TestApplyDelete(t *testing.T) {
	s := New()

	mustApply(t, s, Command{Op: OpPut, Key: "a", Value: "1"})
	if res := mustApply(t, s, Command{Op: OpGet, Key: "a"}); !res.Found {
		t.Fatal("setup failed: expected key \"a\" to exist before Delete")
	}

	mustApply(t, s, Command{Op: OpDelete, Key: "a"})

	if res := mustApply(t, s, Command{Op: OpGet, Key: "a"}); res.Found {
		t.Fatal("expected key \"a\" to be gone after Delete")
	}
}

func TestApplyUnknownOp(t *testing.T) {
	s := New()

	_, err := s.Apply(Command{Op: Op(99), Key: "a"})

	if err == nil {
		t.Fatal("expected an error for unknown op, got nil")
	}
}

func TestApplyZeroValueCommandRejected(t *testing.T) {
	s := New()

	_, err := s.Apply(Command{})

	if err == nil {
		t.Fatal("expected an error for a Command with no Op set, got nil")
	}
}

func TestApplyReplayIsDeterministic(t *testing.T) {
	cmds := []Command{
		{Op: OpPut, Key: "a", Value: "1"},
		{Op: OpPut, Key: "b", Value: "2"},
		{Op: OpPut, Key: "a", Value: "3"},
		{Op: OpDelete, Key: "b"},
		{Op: OpPut, Key: "c", Value: ""},
		{Op: OpDelete, Key: "missing"},
	}

	s1 := New()
	s2 := New()
	for _, cmd := range cmds {
		mustApply(t, s1, cmd)
		mustApply(t, s2, cmd)
	}

	if !reflect.DeepEqual(s1.items, s2.items) {
		t.Fatalf("stores diverged:\n s1 = %v\n s2 = %v", s1.items, s2.items)
	}

	want := map[string]string{"a": "3", "c": ""}
	if !reflect.DeepEqual(s1.items, want) {
		t.Fatalf("got %v, want %v", s1.items, want)
	}
}
