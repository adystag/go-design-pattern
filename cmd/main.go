package main

import (
	"log"

	pattern "github.com/adystag/go-design-pattern"
)

func exemplifyStatePattern() {
	radio := pattern.NewRadio(nil)
	states := []pattern.RadioState{
		pattern.RadioOnState{},
		pattern.RadioOffState{},
	}

	for _, state := range states {
		next := state

		radio.SetState(next)
		radio.Execute()
		log.Println(radio.IsOn())
	}
}

func main() {
	exemplifyStatePattern()
}
