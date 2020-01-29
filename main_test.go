package main

import (
	"testing"
)

func TestGetPrize(t *testing.T) {
	gs := GameState{Clicks: PrizeBigClicks}

	res := gs.getPrize()
	if res != PrizeBig {
		t.Errorf("\nGOT:[%d]\nORG:[%d]\n", PrizeBig, res)
	}

	gs.Clicks = PrizeMediumClicks
	res = gs.getPrize()
	if res != PrizeMedium {
		t.Errorf("\nGOT:[%d]\nORG:[%d]\n", PrizeBig, res)
	}

	gs.Clicks = PrizeSmallClicks
	res = gs.getPrize()
	if res != PrizeSmall {
		t.Errorf("\nGOT:[%d]\nORG:[%d]\n", PrizeBig, res)
	}
}
