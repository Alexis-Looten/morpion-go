package morpion

import "fmt"

func Draw(g *Game, choice string) {

	drawGrid(g)
	drawState(g, choice)
}

func drawGrid(g *Game) {

	fmt.Printf(`
	
 %v | %v | %v
---|---|---
 %v | %v | %v
---|---|---
 %v | %v | %v
	`, g.Box7, g.Box8, g.Box9, g.Box4, g.Box5, g.Box6, g.Box1, g.Box2, g.Box3)
	fmt.Print("\n")
}

func drawState(g *Game, choice string) {
	if g.State == "PlayerWin" {
		fmt.Printf("Player %v WIN ! Well played !\n", g.ActivePlayer)
	} else if g.State == "End" {
		fmt.Print("Egality ! Nobody win, sorry :(\n")
	} else {
		fmt.Printf("Player %v's turn\n", g.ActivePlayer)
	}
}
