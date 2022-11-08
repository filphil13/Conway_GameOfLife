package main

import (
	"math/rand"
)

type GenMap [][]Life

type Game struct {
	genMap      GenMap
	currentLife []Life
	running     bool
	size        int
	seed        int
}

func startGame(size, seed int) Game {
	rand.Seed(int64(seed))

	firstGen := createRandomGen(size / 20)
	genMap := updateGenMap(firstGen, size)

	game := Game{
		genMap:      genMap,
		currentLife: firstGen,
		running:     true,
		size:        size,
		seed:        seed}

	gameLoop(game)
	return game
}

func gameLoop(game Game) {

	//game loop
	for game.running {

		game.currentLife = checkRules(game.currentLife, game)
		game.genMap = updateGenMap(game.currentLife, game.size)
	}
}

func updateGenMap(lives []Life, size int) GenMap {
	newMap := GenMap{}
	for _, life := range lives {
		if life.alive {
			life.generation += 1
			newMap[life.posy][life.posx] = life
		}
	}
	return newMap

}

func createRandomGen(size int) []Life {
	lives := []Life{}

	for i := 0; i < size; i++ {
		life := Life{
			alive:      true,
			posx:       rand.Int() % size,
			posy:       rand.Int() % size,
			generation: 0}

		lives = append(lives, life)
	}
	return lives
}

func checkRules(lives []Life, game Game) []Life {

	for _, life := range lives {
		neighbours := checkNeighbours(life, game)
		if life.alive {
			if neighbours == 2 || neighbours == 3 {
				continue
			} else {
				life.alive = false
			}
		} else {
			if neighbours == 3 {
				life.alive = true
				life.generation = 0
			}
		}
	}
	return lives
}

func checkNeighbours(life Life, game Game) int {
	sumOfNeighbours := 0
	//check horizontal
	if !(game.genMap[(life.posy+1)%game.size][life.posx] != Life{}) {
		if game.genMap[(life.posy+1)%game.size][life.posx].alive {
			sumOfNeighbours++
		}
	} //CHECK TOP
	if !(game.genMap[(life.posy-1)%game.size][life.posx] != Life{}) {
		if game.genMap[(life.posy-1)%game.size][life.posx].alive {
			sumOfNeighbours++
		}
	} //CHECK BOTTOM
	if !(game.genMap[(life.posy)%game.size][(life.posx-1)%game.size] != Life{}) {
		if game.genMap[(life.posy)][(life.posx-1)%game.size].alive {
			sumOfNeighbours++
		}
	} //CHECK LEFT
	if !(game.genMap[(life.posy)][(life.posx+1)%game.size] != Life{}) {
		if game.genMap[(life.posy+1)%game.size][(life.posx+1)%game.size].alive {
			sumOfNeighbours++
		}
	} //CHECK RIGHT
	if !(game.genMap[(life.posy+1)%game.size][(life.posx+1)%game.size] != Life{}) {
		if game.genMap[(life.posy+1)%game.size][(life.posx+1)%game.size].alive {
			sumOfNeighbours++
		}
	} //CHECK TOP RIGHT
	if !(game.genMap[(life.posy+1)%game.size][(life.posx-1)%game.size] != Life{}) {
		if game.genMap[(life.posy+1)%game.size][(life.posx-1)%game.size].alive {
			sumOfNeighbours++
		}
	} //CHECK TOP LEFT
	if !(game.genMap[(life.posy-1)%game.size][(life.posx+1)%game.size] != Life{}) {
		if game.genMap[(life.posy-1)%game.size][(life.posx+1)%game.size].alive {
			sumOfNeighbours++
		}
	} //CHECK BOTTOM RIGHT
	if !(game.genMap[(life.posy-1)%game.size][(life.posx-1)%game.size] != Life{}) {
		if game.genMap[(life.posy-1)%game.size][(life.posx-1)%game.size].alive {
			sumOfNeighbours++
		}
	} //CHECK BOTTOM LEFT
	return sumOfNeighbours
}
