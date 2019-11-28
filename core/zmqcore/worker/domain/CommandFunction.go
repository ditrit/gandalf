package domain

type CommandFunction interface {
}

func (C CommandFunction) executeCommand(command [][]byte, commandStates CommandStates, referenceState ReferenceState) string
