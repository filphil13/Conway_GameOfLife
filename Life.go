package main

type Life struct {
	alive      bool
	posx, posy int
	generation int
}

func createLife(posx, posy int) Life {
	var life = Life{
		alive:      true,
		posx:       posx,
		posy:       posy,
		generation: 0}

	return life
}

func killLife(life Life) {
	life.alive = false
	//if you wanna do something else for data purposes
	//like find which one has the highest generation count
}
