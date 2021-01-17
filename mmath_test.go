package mmath

import "testing"

func TestTopla(t *testing.T) {
	toplam := Topla(2, 5)
	if toplam != 7 {
		t.Error("Beklenen sonuc 7, elde edilen ", toplam)
	}
}
func TestBol(t *testing.T) {
	bol := Bol(15, 3)
	if bol != 5 {
		t.Error("Beklenen sonuc 5, elde edilen ", bol)
	}
}
