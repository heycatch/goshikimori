package goshikimori

import "testing"

func TestNekoSearch(t *testing.T) {
	name_normal := "Lorem ipsum dolor Sit amet consectetur adipiscing Elit"
	compare_normal := "lorem_ipsum_dolor_sit_amet_consectetur_adipiscing_elit"

	if NekoSearch(name_normal) == compare_normal {
		t.Log("NekoSearch normal passed")
	} else {
		t.Error("NekoSearch normal failed")
	}

	name_spaces := "Lorem                ipsum             dolor          Sit             "
	compare_spaces := "lorem_ipsum_dolor_sit"
	if NekoSearch(name_spaces) == compare_spaces {
		t.Log("NekoSearch with big spaces passed")
	} else {
		t.Error("NekoSearch with big spaces failed")
	}
}
