package main

import "errors"

// Dari inisiasi stack dengan maksimal elemen sebanyak 10, implementasikan operasi pop.

var ErrStackUnderflow = errors.New("stack underflow")

type Stack struct {
	// TODO: answer here
	Top  int
	Data []int
}

func NewStack(size int) Stack {
	// TODO: answer here
	return Stack{
		Top:  -1,
		Data: []int{},
	}
}

func (s *Stack) Push(Elemen int) error {
	// TODO: answer here
	if len(s.Data) == 10 {
		return ErrStackUnderflow
	}

	s.Top += 1
	s.Data = append(s.Data, Elemen)

	return nil
}

func (s *Stack) Pop() (int, error) {
	// TODO: answer here
	if s.Top == -1 {
		return 0, ErrStackUnderflow
	}

	e := s.Data[s.Top]
	s.Top -= 1
	s.Data = s.Data[:s.Top+1]

	return e, nil
}
