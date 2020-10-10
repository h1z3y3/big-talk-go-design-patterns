package facade

import (
	"testing"
)

func TestApiImpl_Test(t *testing.T) {
	api := NewAPIFacede()

	ret := api.Test()
	want := "facade: a->A module b->B module"

	if ret != want {
		t.Fatalf("want: %s, got: %s", want, ret)
	}

}
