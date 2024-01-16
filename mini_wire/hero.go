package mini_wire

type Hero struct {
	Name string
}

func NewHero(name string) *Hero {
	return &Hero{
		Name: name,
	}
}

func (h *Hero) GetName() string {
	return h.Name
}

func (h *Hero) SetName(name string) {
	h.Name = name
}
