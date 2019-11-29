package worker

type EventFunction interface {
}

func (c EventFunction) executeCommand(event [][]byte)
