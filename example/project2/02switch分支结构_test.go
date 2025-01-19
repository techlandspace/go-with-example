package project2

import "testing"

func TestSwitchExample(t *testing.T) {
	switch_example("frag")
}

func TestSwitchExample2(t *testing.T) {
	switch_fallthrough_example("apple")
}

func TestSwitchExample3(t *testing.T) {
	switch_type_example()
}
