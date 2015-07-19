package main

type Sequencer struct {
	Total int
}

func NewSequencer() *Sequencer {
	return &Sequencer{0}
}

func (s *Sequencer) Increment(amount int) {
	s.Total += amount
}
