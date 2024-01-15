package tetris

import (
	"time"
)

const HEIGHT = 20
const WIDTH = 10

type game int

const (
	PLAYING game = iota
	PAUSED
	ENDED
)

type vector struct {
	x int
	y int
}

type Tetris struct {
	board     [HEIGHT][WIDTH]int
	pos       vector
	piece     piece
	score     int
	state     game
	Fallspeed *time.Timer
}

func (t *Tetris) blockbyPosition(v vector) vector {
	return vector{t.pos.x + v.x, t.pos.y + v.y}
}

func (t *Tetris) collision() bool {
	for _, v := range t.piece.shape {

		if t.blockbyPosition(v).y >= HEIGHT {
			return true
		}

		if t.blockbyPosition(v).x < 0 || t.blockbyPosition(v).x >= WIDTH {
			return true
		}

		if t.board[t.blockbyPosition(v).y][t.blockbyPosition(v).x] != 0 {
			return true
		}
	}
}

func (t *Tetris) Fall() {
	if t.state != PLAYING {
		return
	}

	for {
		if !t.moveIfPossible()
	}

}

func (t *Tetris) moveLeft() {
	t.moveIfPossible(vector{0,-1})
}

func (t *Tetris) moveRight() {
	t.moveIfPossible(vector{0,1})
}

func (t *Tetris) rotate() {
	if t.state != PLAYING {
		return
	}

	t.piece.rotate()
	if t.collision() {
		t.piece.rotate()
		t.piece.rotate()
		t.piece.rotate()
	}	
}

func (t *Tetris) drop() {
	if t.state != PLAYING {
		return
	}
	t.Fallspeed.Reset(50)
}

func (t *Tetris) moveIfPossible(v vector) {
	if t.state != PLAYING {
		return
	}

	t.pos.x += v.x
	if t.collision() {
		t.pos.x -= v.x
	}
}

func (t *Tetris) shapeToBoard() {
	for _, v := range t.shape.Blocks {
		t.board[t.blockbyPosition(v).y][t.blockbyPosition(v).x] = 1
	}
}

func (t *Tetris) removeFullLines() {
	for i := 0; i < HEIGHT; i++ {
		remove := true
		for j := 0; j < WIDTH; j++ {
			if t.board[i][j] == 0 {
				remove = false
				break
			}
		}

		if remove {
			for j := 0; j < WIDTH; j++ {
				t.board[i][j] = 0
			}
			t.score += 1
		}
	}
}

func (t *Tetris) Tick() {
	t.Fall()
	t.removeFullLines()
}

func (t *Tetris) Start() {
	t.state = PLAYING
}

func (t *Tetris) Pause() {
	t.state = PAUSED
}

func (t *Tetris) End() {
	t.state = ENDED
}

func (t *Tetris) State() game {
	return t.state
}

func (t *Tetris) Board() [HEIGHT][WIDTH]int {
	return t.board
}

func (t *Tetris) Pos() vector {
	return t.pos
}

func (t *Tetris) Shape() *shape.Shape {
	return t.shape
}

func (t *Tetris) Score() int {

	return t.score
}

func (t *Tetris) Speed() time.Duration {
	return *t.speed
}

func NewTetris() *Tetris {
	return &Tetris{
		pos:   vector{WIDTH / 2, 0},
		shape: shape.RandomShape(),
		score: 0,
		state: PAUSED,
		speed: time.Second,
	}
}
