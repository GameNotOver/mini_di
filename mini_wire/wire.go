//go:build wireinject
// +build wireinject

package mini_wire

import "github.com/google/wire"

func NewMissionImpl(name string) *Mission {
	wire.Build(NewMonster, NewHero, NewMission)
	return nil
}
