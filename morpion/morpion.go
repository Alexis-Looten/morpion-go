package morpion

type Game struct {
	State        string //Game state
	ActivePlayer int    //Player who is playing
	Box1         string
	Box2         string
	Box3         string
	Box4         string
	Box5         string
	Box6         string
	Box7         string
	Box8         string
	Box9         string
}

func New() (*Game, error) {

	g := &Game{
		State:        "",
		ActivePlayer: 1,
		Box1:         "1",
		Box2:         "2",
		Box3:         "3",
		Box4:         "4",
		Box5:         "5",
		Box6:         "6",
		Box7:         "7",
		Box8:         "8",
		Box9:         "9",
	}

	return g, nil
}

func (g *Game) CheckACase(choice string) {
	if g.ActivePlayer == 1 {
		switch choice {
		case "1":
			g.Box1 = "X"

		case "2":
			g.Box2 = "X"

		case "3":
			g.Box3 = "X"

		case "4":
			g.Box4 = "X"

		case "5":
			g.Box5 = "X"

		case "6":
			g.Box6 = "X"

		case "7":
			g.Box7 = "X"

		case "8":
			g.Box8 = "X"

		case "9":
			g.Box9 = "X"

		}
	}

	if g.ActivePlayer == 2 {
		switch choice {
		case "1":
			g.Box1 = "O"

		case "2":
			g.Box2 = "O"

		case "3":
			g.Box3 = "O"

		case "4":
			g.Box4 = "O"

		case "5":
			g.Box5 = "O"

		case "6":
			g.Box6 = "O"

		case "7":
			g.Box7 = "O"

		case "8":
			g.Box8 = "O"

		case "9":
			g.Box9 = "O"
		}
	}
}

func (g *Game) ChangePlayer() {

	if g.ActivePlayer == 1 {
		g.ActivePlayer = 2
	} else {
		g.ActivePlayer = 1
	}

}

func (g *Game) MakeAChoice(choice string) {

	switch g.State {
	case "End", "PlyerWin":
		return
	}

	g.CheckACase(choice)

	if g.ItEnd() {
		g.State = "End"
	} else if g.PlayerWin() {
		g.State = "PlayerWin"
	} else {
		g.ChangePlayer()
	}

}

func (g *Game) PlayerWin() bool {

	if g.ActivePlayer == 1 {
		if g.Box1 == "X" && g.Box2 == "X" && g.Box3 == "X" {
			return true
		}
		if g.Box4 == "X" && g.Box5 == "X" && g.Box6 == "X" {
			return true
		}
		if g.Box7 == "X" && g.Box8 == "X" && g.Box9 == "X" {
			return true
		}
		if g.Box1 == "X" && g.Box5 == "X" && g.Box9 == "X" {
			return true
		}
		if g.Box3 == "X" && g.Box5 == "X" && g.Box7 == "X" {
			return true
		}
		if g.Box1 == "X" && g.Box4 == "X" && g.Box7 == "X" {
			return true
		}
		if g.Box2 == "X" && g.Box5 == "X" && g.Box8 == "X" {
			return true
		}
		if g.Box3 == "X" && g.Box6 == "X" && g.Box9 == "X" {
			return true
		}
	}

	if g.ActivePlayer == 2 {
		if g.Box1 == "O" && g.Box2 == "O" && g.Box3 == "O" {
			return true
		}
		if g.Box4 == "O" && g.Box5 == "O" && g.Box6 == "O" {
			return true
		}
		if g.Box7 == "O" && g.Box8 == "O" && g.Box9 == "O" {
			return true
		}
		if g.Box1 == "O" && g.Box5 == "O" && g.Box9 == "O" {
			return true
		}
		if g.Box3 == "O" && g.Box5 == "O" && g.Box7 == "O" {
			return true
		}
		if g.Box1 == "O" && g.Box4 == "O" && g.Box7 == "O" {
			return true
		}
		if g.Box2 == "O" && g.Box5 == "O" && g.Box8 == "O" {
			return true
		}
		if g.Box3 == "O" && g.Box6 == "O" && g.Box9 == "O" {
			return true
		}
	}

	return false
}

func (g *Game) ItEnd() bool {
	if g.Box1 != "1" && g.Box2 != "2" && g.Box3 != "3" && g.Box4 != "4" && g.Box5 != "5" && g.Box6 != "6" && g.Box7 != "7" && g.Box8 != "8" && g.Box9 != "9" {
		return true
	} else {
		return false
	}
}
