package mahjong

// An intent defines an intended "combo" to be made
// The user indicates the tiles that are part of the combo, and then indicates
// which tiles are valid candidates for that combo.
type Intent interface {
	// Tiles returns an immutable slice of tiles that allow us to make comparisons
	// i.e. (checkValidIntent)
	// as well as
	Tiles() []Tile

	// Targets returns an immutable slice of tiles with "TileID"s of EmptyTile
	// We use these to compare against discarded tiles
	Targets() []Tile
}

type intentImpl struct{}

func (i intentImpl) Tiles() []Tile {
	return nil
}

func (i intentImpl) Targets() []Tile {
	return nil
}
