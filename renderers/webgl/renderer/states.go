package renderer

import "github.com/tokkenno/seed/core"

type States struct {
	list map[core.Id]map[core.Id]*State
}

func (states *States) Close() {
	states.list = make(map[core.Id]map[core.Id]*State)
}

func (states *States) onSceneDispose(scene *core.Scene) {
	// TODO: scene.RemoveEventListener onSceneDispose
	delete(states.list, scene.GetId())
}

func (states *States) Get(scene *core.Scene, camera *core.Camera) *State {
	if value, ok := states.list[scene.GetId()][camera.GetId()]; ok {
		// TODO: scene.AddEventListener onSceneDispose
		return value
	} else {
		state := NewState()
		states.list[scene.GetId()][camera.GetId()] = state
		return state
	}
}

var (
	states *States = &States{
		list: make(map[core.Id]map[core.Id]*State),
	}
)

func GetStatesInstance() *States {
	return states
}