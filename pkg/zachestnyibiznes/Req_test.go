package zachestnyibiznes_test

import (
	"fmt"
	"testing"

	zachestnyibiznes "github.com/RB-PRO/trudeks/pkg/zachestnyibiznes"
)

func TestContacts(t *testing.T) {
	z4b, ErrLoad := zachestnyibiznes.LoadConfig("zachestnyibiznes.json")
	if ErrLoad != nil {
		t.Error(ErrLoad)
	}
	ID := "9715255412"
	contacts, ErrCont := z4b.Contacts(ID)
	if ErrCont != nil {
		t.Error(ErrCont)
	}
	fmt.Printf("%+v\n", contacts)

	xlsx := zachestnyibiznes.NewXLSX("test.xlsx")
	xlsx.WriteXLSX(ID, contacts)
	xlsx.CloceAndSaveXLSX()
}
