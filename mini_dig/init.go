package mini_dig

func Init() {
	ctr := GetAppContainer()
	ctr.MustRegister(func() string { return "A" })
	ctr.MustRegister(NewHero)
	ctr.MustRegister(NewMonster)
	ctr.MustRegister(NewMission)
}
