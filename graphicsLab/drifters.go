package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game holds all state that needs to persist across frames.
type Game struct {
	circle *ebiten.Image // Exercise 1 & 2: the drifting sprite
	square *ebiten.Image // Exercise 3: second sprite drawn on top
	x, y   float64      // Exercise 2: position variables incremented each frame
}

// NewGame loads both sprites and returns an initialised *Game.
// Separating loading from main() keeps main() clean and makes it easy to
// return an error instead of panicking.
func NewGame() (*Game, error) {
	// Exercise 1 – load a small PNG and store it on the struct.
	// Replace the filenames with your own assets.
	circle, _, err := ebitenutil.NewImageFromFile("circle.png")
	if err != nil {
		return nil, err
	}

	square, _, err := ebitenutil.NewImageFromFile("square.png")
	if err != nil {
		return nil, err
	}

	return &Game{
		circle: circle,
		square: square,
		x:      50, // starting position for the drifter
		y:      50,
	}, nil
}

// Update is called every tick (default 60 TPS).
// This is where game logic lives – input, physics, state changes.
func (g *Game) Update() error {
	// Exercise 2 – increment x and y each frame so the sprite drifts.
	g.x += 1
	g.y += 0.5
	return nil
}

// Draw is called every frame to render the current state.
// It must only read state, never mutate it (Update owns mutation).
func (g *Game) Draw(screen *ebiten.Image) {

	// ------------------------------------------------------------------
	// Exercise 2 – draw the circle at the drifting position.
	// We create a fresh DrawImageOptions for each sprite so transforms
	// are completely independent.
	// ------------------------------------------------------------------
	circleOp := &ebiten.DrawImageOptions{}
	circleOp.GeoM.Scale(0.5, 0.5)
	circleOp.GeoM.Translate(g.x, g.y) // uses the variables updated in Update()
	screen.DrawImage(g.circle, circleOp)

	// ------------------------------------------------------------------
	// Exercise 3 – draw two different sprites at distinct positions.
	//
	// Draw order matters: whatever is drawn LAST appears ON TOP.
	// Here the square (drawn second) will overlap the circle if they share
	// the same screen region.  Swap the two blocks to see the circle win.
	// ------------------------------------------------------------------

	// First sprite – circle at (50, 50)  ← drawn first → underneath
	firstOp := &ebiten.DrawImageOptions{}
	firstOp.GeoM.Translate(50, 50)
	screen.DrawImage(g.circle, firstOp)

	// Second sprite – square at (100, 100) ← drawn second → on top
	secondOp := &ebiten.DrawImageOptions{}
	secondOp.GeoM.Translate(100, 100)
	screen.DrawImage(g.square, secondOp)
}

// Layout tells Ebiten the logical screen size.
// The window can be resized independently; Ebiten handles the scaling.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Sprite Exercises")

	game, err := NewGame()
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}