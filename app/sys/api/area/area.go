package area

import (
	"gopkg.in/goyy/goyy.v0/app/sys/internal/area"
)

// Get the full name of the area.
func Name(id string) string {
	a := area.NewEntity()
	a.SetId(id)
	err := area.Mgr.Get(a)
	if err != nil {
		return ""
	}
	return a.Fullname()
}
