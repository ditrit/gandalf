package connector

type GandalfApplication struct {
	workerCommandsMap        map[string][]string
	workerCommandSendFileMap map[string]string
	command                  Command
	event                    Event
}

func (cg GandalfApplication) new() {
	cg.workerCommandsMap = make(map[string][]string)
	cg.workerCommandSendFileMap = make(map[string]string)
	cg.command = Command.new()
	cg.event = Event.new()

	//RUN
	go cg.command.run()
	go cg.event.run()
}

func (cg GandalfApplication) getWorkerCommands(worker string) []string {
	return cg.workerCommandsMap[worker]
}

func (cg GandalfApplication) addWorkerCommands(worker, command string) {
	var sizeList = len(cg.workerCommandsMap[worker])
	cg.workerCommandsMap[worker][sizeList] = command

}

func (cg GandalfApplication) getWorkerCommandSendFile(worker string) string {
	return cg.workerCommandSendFileMap[worker]
}

func (cg GandalfApplication) addWorkerCommandSendFile(worker, commandSendFile string) {
	cg.workerCommandSendFileMap[worker] = commandSendFile
}
