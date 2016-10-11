package role

import (
	"testing"

	_ "gopkg.in/goyy/goyy.v0/app/tst"
	"gopkg.in/goyy/goyy.v0/data/domain"
)

func TestPostSelectCountBySift(t *testing.T) {
	sIdEQ, _ := domain.NewSift("sIdEQ", "1")
	sDeletionEQ, _ := domain.NewSift("sDeletionEQ", "0")
	out, _ := Mgr.SelectCountBySift(sIdEQ, sDeletionEQ)
	expected := 0
	if out != expected {
		t.Errorf(`SelectCountBySift:"%v", want:"%v"`, out, expected)
	}
}
