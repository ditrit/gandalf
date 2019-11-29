package worker

type GandalfApplication struct {
	mapUUIDCommandStates map[string][]string
	mapUUIDState map[string]*ReferenceState
	routine Routine
}

func (ga GandalfApplication) new() {
	g.mapUUIDCommandStates = make(map[string][]string)
	g.mapUUIDState = make(map[string]*ReferenceState)
	g.routine = Routine.new(type)

	go g.routine.run()
}

func (ga GandalfApplication) GetMapUUIDCommandStates() map[string]List {
	return c.mapUUIDCommandStates
}

func (ga GandalfApplication) SetMapUUIDCommandStates(mapUUIDCommandStates map[string]List) {
	return c.mapUUIDCommandStates = mapUUIDCommandStates
}

func (ga GandalfApplication) GetMapUUIDState() map[string]ReferenceState {
	return c.mapUUIDState
}

func (ga GandalfApplication) SetMapUUIDState(mapUUIDState map[string]ReferenceState) {
	return c.mapUUIDState = mapUUIDState
}

func (ga GandalfApplication) New() {
	r.mapUUIDCommandStates = make(map[string]List)
	r.mapUUIDState = make(map[string]ReferenceState)
}

func (ga GandalfApplication) GetMapUUIDCommandStatesByUUID(uuid string) map[string]List {
	return c.mapUUIDCommandStates.get(uuid)
}

func (ga GandalfApplication) GetMapUUIDStateByUUID(uuid string) map[string]ReferenceState {
	return c.mapUUIDState.get(uuid)
}