package kv

import "fmt"

type Op int

const (
	OpInvalid Op = iota // 0: never a valid operation
	OpGet
	OpPut
	OpDelete
)

type Command struct {
	Op    Op
	Key   string
	Value string // only used by OpPut
}

type Result struct {
	Value string
	Found bool
}

func (s *Store) Apply(cmd Command) (Result, error) {
	switch cmd.Op {
	case OpGet:
		v, ok := s.Get(cmd.Key)
		return Result{Value: v, Found: ok}, nil
	case OpPut:
		s.Put(cmd.Key, cmd.Value)
		return Result{}, nil
	case OpDelete:
		s.Delete(cmd.Key)
		return Result{}, nil
	default:
		return Result{}, fmt.Errorf("unknown op: %d", cmd.Op)
	}
}
