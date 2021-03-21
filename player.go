package mahjong

import (
	"fmt"
	"sync"
)

type Player struct {
	hand    Hand
	intents []Intent

	sync.Mutex
	Observable
}

func (p *Player) addTile(t Tile) {
	p.Lock()
	p.hand = append(p.hand, t)

	for _, o := range p.observers {
		po, ok := o.(PlayerObserver)
		if !ok {
			continue
		}

		po.Draw(p.hand)
	}

	p.Unlock()
}

func (p *Player) discard(tid TileID) (error, Tile) {
	p.Lock()
	defer p.Unlock()

	for i, t := range p.hand {
		if t.Id() == tid {
			p.hand = append(p.hand[:i], p.hand[i+1:]...)
			return nil, t
		}
	}

	return fmt.Errorf("could not find tile with id %v", tid), nil
}

// Hand lets us retrieve a copy of the player's hand for calculations
func (p *Player) Hand() Hand {
	p.Lock()
	mh := make([]Tile, len(p.hand))
	copy(mh, p.hand)
	p.Unlock()

	return mh
}

func (p *Player) addIntent(i Intent) {
	p.Lock()
	p.intents = append(p.intents, i)

	for _, o := range p.observers {
		po, ok := o.(PlayerObserver)
		if !ok {
			continue
		}

		po.IntentAdded(i)
	}

	p.Unlock()
}
