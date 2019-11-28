package domain

type CommandStateManager struct {
	mapUUIDCommandStates map[string]List
	mapUUIDState map[string]ReferenceState
}

func (c CommandStateManager) GetMapUUIDCommandStates() map[string]List {
	return c.mapUUIDCommandStates
}

func (c CommandStateManager) SetMapUUIDCommandStates(mapUUIDCommandStates map[string]List) {
	return c.mapUUIDCommandStates = mapUUIDCommandStates
}

func (c CommandStateManager) GetMapUUIDState() map[string]ReferenceState {
	return c.mapUUIDState
}

func (c CommandStateManager) SetMapUUIDState(mapUUIDState map[string]ReferenceState) {
	return c.mapUUIDState = mapUUIDState
}

func (c CommandStateManager) New() {
	r.mapUUIDCommandStates = make(map[string]List)
	r.mapUUIDState = make(map[string]ReferenceState)
}

func (c CommandStateManager) GetMapUUIDCommandStatesByUUID(uuid string) map[string]List {
	return c.mapUUIDCommandStates.get(uuid)
}

func (c CommandStateManager) GetMapUUIDStateByUUID(uuid string) map[string]ReferenceState {
	return c.mapUUIDState.get(uuid)
}
