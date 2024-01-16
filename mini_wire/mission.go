package mini_wire

import "fmt"

type Mission struct {
	CurrentMonster *Monster
	CurrentHero    *Hero
}

func NewMission(monster *Monster, hero *Hero) *Mission {
	return &Mission{
		CurrentMonster: monster,
		CurrentHero:    hero,
	}
}

func (m *Mission) Start() {
	fmt.Printf("hero %v beat monster %v\n", m.CurrentHero.GetName(), m.CurrentMonster.GetName())
}
