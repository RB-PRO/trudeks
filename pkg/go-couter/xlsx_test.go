package couterparser

import (
	"fmt"
	"testing"
)

func TestInput(t *testing.T) {
	MapCouter, ErrInput := InputFileXlsxCouter("../../суды.xlsx")
	if ErrInput != nil {
		t.Error(ErrInput)
	}
	if len(MapCouter) == 0 {
		t.Error("Lenght of MapCouter = 0")
	}
	fmt.Println(MapCouter)
}
