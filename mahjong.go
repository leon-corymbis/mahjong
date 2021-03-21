package mahjong

/**
* Hand describes an interface to interact with and receive information about a hand
 */
type Hand []Tile

/**
* Round describes an interface that holds information about the current game round
 */
type Round interface {
	CurrentWind() Wind
}

/**
* Wind describes what direction the player is sitting
 */
type Wind uint64

const (
	East Wind = iota
	South
	West
	North
)

/**
* Suit describes a set of predetermined tiles
 */
type Suit string
