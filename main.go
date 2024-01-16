package main

import (
	"mini_di/mini_dig"
	"mini_di/mini_wire"
)

func main() {
	// di
	mini_dig.Init()
	ctr := mini_dig.GetAppContainer()
	ctr.MustCall(func(mission *mini_dig.Mission) {
		mission.Start()
	})
	// wire
	mission := mini_wire.NewMissionImpl("A")
	mission.Start()
}
