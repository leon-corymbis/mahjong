package mahjong

/**
* Hand describes an interface to interact with and receive information about a hand
 */
type Hand interface {
	GetTileAtPos(idx uint64) Tile
	AddTile(tile Tile)
	Size() uint64
	Tiles() []Tile
	Wind() Wind
}

/**
* Round describes an interface that holds information about the current game round
 */
type Round interface {
	CurrentWind() Wind
}

/**
* Tile describes an interface for holding information about a tile and displaying it
 */
type Tile interface {
	Value() uint64
	Suit() Suit
	Wind() Wind
}

/**
* Wind describes what direction the player is sitting
 */
type Wind interface {
}

/**
* Suit describes a set of predetermined tiles
 */
type Suit interface {
}

/**
* Ruleset defines an API for plug-and-play scoring of hands
 */
type Ruleset interface {
	Score(hand Hand, round Round) (error, uint64)
	Suits() []Suit
}
