package mahjong

import (
	"testing"
)

type poImpl struct {
	drawCalled        bool
	intentAddedCalled bool
}

func (poImpl) GetID() string {
	return "testing-player-observer"
}

func (p *poImpl) Draw(Hand) {
	p.drawCalled = true
}

func (p *poImpl) IntentAdded(Intent) {
	p.intentAddedCalled = true
}

func TestObserveDraw(t *testing.T) {
	po := &poImpl{
		drawCalled: false,
	}

	p := &Player{}
	p.registerObserver(po)

	p.addTile(emptyTileImpl)

	if !po.drawCalled {
		t.Error("Expected observer to receive a draw notification")
	}
}

func TestObserveIntent(t *testing.T) {
	po := &poImpl{
		intentAddedCalled: false,
	}

	p := &Player{}
	p.registerObserver(po)

	ii := intentImpl{}

	p.addIntent(ii)

	if !po.intentAddedCalled {
		t.Error("Expected observer to receive an intentAdded notification")
	}
}
