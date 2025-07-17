package app

type Components struct {
	InputStates map[string]InteractableState
	InputNames  map[string]bool
}

func NewComponents() Components {
	return Components{
		InputStates: map[string]InteractableState{},
		InputNames:  map[string]bool{},
	}
}
