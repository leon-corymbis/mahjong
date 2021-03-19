package mahjong

/**
* Game holds information about the state of the game
 */
type Game struct {
	ruleset Ruleset
	hands   map[Wind]Hand
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
		hands:   make(map[Wind]Hand, 4),
	}

	return g
}

// TODO: get each wind from the ruleset, and then give each player 13 tiles
