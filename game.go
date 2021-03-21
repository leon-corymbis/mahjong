package mahjong

import "sync"

/**
* Game holds information about the state of the game
 */
type Game struct {
	ruleset Ruleset
	hands   map[Wind]Player

	turn Wind

	sync.Mutex
}

/**
* GameConfig lets you customize the game before it begins
 */
type GameConfig interface {
	Ruleset() Ruleset
}

/**
* New takes in a configuration for a new game and prepares to begin.
 */
func New(gc GameConfig) *Game {
	g := &Game{
		ruleset: gc.Ruleset(),
		hands:   make(map[Wind]Player, 4),
	}

	return g
}

/**
* RegisterIntent confers the desire to register a
* group of tiles as a "combo", and to satisfy it at
* the earliest possible convenience
 */
func (g *Game) RegisterIntent(w Wind, i Intent) error {

	return nil
}
