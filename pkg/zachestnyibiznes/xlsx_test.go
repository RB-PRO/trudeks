package zachestnyibiznes

import "testing"

func TestXlsx(t *testing.T) {
	xlsx := NewXLSX("test.xlsx")
	xlsx.CloceAndSaveXLSX()
}
