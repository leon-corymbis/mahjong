package mahjong

import "sync"

// Observer is a pattern that allows us to notify
// other libraries
type Observer interface {
	GetID() string
}

/**
 * PlayerObserver allows the server to subscribe to events specific to a
 * single player
 */
type PlayerObserver interface {
	// Draw gives the player a hand update
	Draw(h Hand)

	IntentAdded(intent Intent)

	Observer
}

/**
 * GameObserver allows the server to subscribe to events that affect everyone
 * in the game.
 */
type GameObserver interface {
	// Discard tells everyone what tile was discarded and by whom
	Discard() (Tile, Wind)

	// Draw tells everyone who drew a tile
	Draw() Wind

	// Taken tells everyone when someone else takes a tile, and who took it
	// Note: this can only happen to the last tile discarded
	Taken() (Tile, Wind)

	Observer
}

type Observable struct {
	observers []Observer

	sync.Mutex
}

func (o *Observable) registerObserver(oo Observer) {
	o.Lock()
	defer o.Unlock()

	for _, oos := range o.observers {
		if oos.GetID() == oo.GetID() {
			return
		}
	}

	o.observers = append(o.observers, oo)
}

func (o *Observable) deregisterObserver(oo Observer) {
	o.Lock()
	defer o.Unlock()

	for i, oos := range o.observers {
		if oos.GetID() == oo.GetID() {
			lastIdx := len(o.observers) - 1
			o.observers[lastIdx], o.observers[i] = o.observers[i], o.observers[lastIdx]
			o.observers = o.observers[:lastIdx]
			return
		}
	}
}
