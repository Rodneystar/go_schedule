package physical

import (
	"fmt"
)

type Switchable interface {
	On()
	Off()
	State() bool
}

type fakeSwitchable struct {
	Active bool
}

func (s *fakeSwitchable) On() {
	fmt.Println("---ON---")
}
func (s *fakeSwitchable) Off() {
	fmt.Println("---OFF---")
}

func (s *fakeSwitchable) State() bool {
	return s.Active
}

type radioPiHatPlug struct {
	active bool
}

func (s *radioPiHatPlug) On() {

}

func (s *radioPiHatPlug) Off() {

}

func (s *radioPiHatPlug) State() {

}
