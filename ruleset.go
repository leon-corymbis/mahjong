package mahjong

/**
* Ruleset defines an API for plug-and-play scoring of hands
 */
type Ruleset interface {
	Score(hand Hand, round Round) (error, uint64)
	Suits() []Suit
	CheckValidIntent(intent Intent) (error, bool)
}

type rulesetImpl struct {
}

func (rulesetImpl) Score(hand Hand, round Round) (error, uint64) {
	return nil, 0
}

var testingSuits []Suit = []Suit{
	"stick",
	"number",
	"circle",
	"wind",
	"flower",
}

func (rulesetImpl) Suits() []Suit {
	return testingSuits
}

func (rulesetImpl) CheckValidIntent(Intent) (error, bool) {
	return nil, true
}
