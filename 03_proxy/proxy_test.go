package proxy

import (
	"testing"
	"fmt"
)

func TestProxy(t *testing.T) {
	var image Image

	image = &ImageProxy{}

	res := image.Get()

	fmt.Println(res)

	if res != "pre:real_image_url:after" {
		t.Fail()
	}
}