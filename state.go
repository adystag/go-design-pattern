package pattern

import "errors"

type RadioState interface {
	Execute(*Radio) error
}

type Radio struct {
	isOn  bool
	state RadioState
}

func (r *Radio) Execute() error {
	return r.state.Execute(r)
}

func (r *Radio) TurnOn() error {
	if r.isOn {
		return errors.New("radio already turned on")
	}

	r.isOn = true

	return nil
}

func (r *Radio) TurnOff() error {
	if !r.isOn {
		return errors.New("radio already turned off")
	}

	r.isOn = false

	return nil
}

func (r *Radio) SetState(state RadioState) {
	r.state = state
}

func (r Radio) IsOn() bool {
	return r.isOn
}

func NewRadio(state RadioState) Radio {
	return Radio{
		state: state,
	}
}

type RadioOnState struct{}

func (RadioOnState) Execute(r *Radio) error {
	return r.TurnOn()
}

type RadioOffState struct{}

func (RadioOffState) Execute(r *Radio) error {
	return r.TurnOff()
}
