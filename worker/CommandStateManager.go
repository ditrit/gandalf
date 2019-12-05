package worker

type CommandStateManager struct {
	mapUUIDCommandStates map[string]List
	mapUUIDState map[string]ReferenceState
}

func (c CommandStateManager) New() {
	r.mapUUIDCommandStates = make(map[string]List)
	r.mapUUIDState = make(map[string]ReferenceState)
}

func (c CommandStateManager) GetMapUUIDCommandStates() (mapUUIDCommandStates map[string]List, err error) {
	return c.mapUUIDCommandStates
}

func (c CommandStateManager) SetMapUUIDCommandStates(mapUUIDCommandStates map[string]List) err error {
	return c.mapUUIDCommandStates = mapUUIDCommandStates
}

func (c CommandStateManager) GetMapUUIDState() (mapUUIDState map[string]ReferenceState,  err error) {
	return c.mapUUIDState
}

func (c CommandStateManager) SetMapUUIDState(mapUUIDState map[string]ReferenceState)  err error) {
	return c.mapUUIDState = mapUUIDState
}


func (c CommandStateManager) GetMapUUIDCommandStatesByUUID(uuid string) (commandState map[string]List,  err error) {
	return c.mapUUIDCommandStates.get(uuid)
}

func (c CommandStateManager) GetMapUUIDStateByUUID(uuid string) (state map[string]ReferenceState,  err error) {
	return c.mapUUIDState.get(uuid)
}
