package mahjong

type TileID uint64

const EmptyTile TileID = 0

/**
* Tile describes an interface for holding information about a tile and displaying it
 */
type Tile interface {
	// This is a unique-per-id idx that we use to identify tiles
	// specifically, for discarding
	Id() TileID
	Value() uint64
	Suit() Suit
	Wind() Wind
}

type tileImpl struct {
	ID    TileID
	suit  Suit
	wind  Wind
	value uint64
}

func (m tileImpl) Id() TileID {
	return m.ID
}

func (m tileImpl) Suit() Suit {
	return m.suit
}

func (m tileImpl) Wind() Wind {
	return m.wind
}

func (m tileImpl) Value() uint64 {
	return m.value
}

var emptyTileImpl tileImpl = tileImpl{
	ID:    EmptyTile,
	suit:  "mock",
	wind:  East,
	value: 0,
}
