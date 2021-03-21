package mahjong

import (
	"testing"
)

func TestAddTileAndRetrieveTile(t *testing.T) {
	p := &Player{}

	p.addTile(emptyTileImpl)

	found := false
	for _, t := range p.Hand() {
		if t == emptyTileImpl {
			found = true
		}
	}

	if !found {
		t.Error("Empty Tile should have been added to hand")
	}
}

func TestAddingMultipleTiles(t *testing.T) {
	p := &Player{}

	for i := 0; i < 14; i++ {
		p.addTile(emptyTileImpl)
	}

	if len(p.Hand()) != 14 {
		t.Error("Player's hand should be the same as the number of tiles added")
	}
}

func TestDiscardingTile(tt *testing.T) {
	p := &Player{}

	for i := 0; i < 14; i++ {
		t := &tileImpl{
			ID: TileID(i),
		}

		p.addTile(t)
	}

	for i := 13; i > 0; i-- {
		err, t := p.discard(TileID(i))
		if err != nil {
			tt.Error(err)
		}

		if t.Id() != TileID(i) {
			tt.Errorf("tile discarded %v does not match requested %v", t.Id(), i)
		}
	}

	err, t := p.discard(TileID(14))
	if err == nil {
		tt.Error("no error discarding from empty hand")
	}

	if t != nil {
		tt.Errorf("Expected no tiles, received %v", t)
	}
}

func TestAddIntent(t *testing.T) {
	p := &Player{}

	ii := intentImpl{}

	p.addIntent(ii)
}
