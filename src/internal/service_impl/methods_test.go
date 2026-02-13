package service_impl

import (
	"domain"
	"fmt"
	"testing"
)

func TestWin(t *testing.T) {
	testTable := []struct {
		base    domain.Base
		side    uint8
		expBool bool
	}{
		{
			base: domain.Base{
				[3][3]uint8{{1, 2, 0}, {1, 2, 0}, {1, 0, 0}},
				5,
			},
			side:    domain.X,
			expBool: true,
		},
		{
			base: domain.Base{
				[3][3]uint8{{1, 1, 1}, {1, 2, 0}, {1, 0, 0}},
				6,
			},
			side:    domain.X,
			expBool: true,
		},
		{
			base: domain.Base{
				[3][3]uint8{{0, 1, 1}, {1, 2, 0}, {1, 0, 0}},
				6,
			},
			side:    domain.X,
			expBool: false,
		},
		{
			base: domain.Base{
				[3][3]uint8{{1, 2, 0}, {2, 1, 0}, {2, 0, 1}},
				6,
			},
			side:    domain.X,
			expBool: true,
		},
	}

	for _, tcs := range testTable {
		if winOrDraw(&tcs.base, tcs.side) == tcs.expBool {
			fmt.Println("Correct")
		} else {
			t.Error("Incorrect")
		}
	}
}
