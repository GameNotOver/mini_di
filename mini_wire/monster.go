package mini_wire

type Monster struct {
	Name string
}

func NewMonster(name string) *Monster {
	return &Monster{
		Name: name,
	}
}

func (m *Monster) GetName() string {
	return m.Name
}

func (m *Monster) SetName(name string) {
	m.Name = name
}
