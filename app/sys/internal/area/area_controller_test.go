package area

import (
	"testing"

	_ "gopkg.in/goyy/goyy.v0/app/tst"
	"gopkg.in/goyy/goyy.v0/test/assert"
)

func TestControllerIndex(t *testing.T) {
	if !assert.HTTPSuccess(t, ctl.Index, "GET", ctl.ApiIndex(), nil) {
		t.Errorf(`GET: %s: Fail`, ctl.ApiIndex())
	}
}
