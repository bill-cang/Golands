package main

import (
	"./f"
	"./m"
)

type Sone struct {
	m.Mather
	f.Father
}

func (s *Sone) NewSone() *Sone {
	return s
}
