package builder

import "testing"

func TestBuilder(t *testing.T) {
	builder := NewBuilder()
	director := &Director{}
	director.SetBuilder(builder)

	house := director.BuildHouse()

	if house.foundation != "concrete foundation" {
		t.Error("Foundation not built correctly")
	}
	if house.walls != "brick walls" {
		t.Error("Walls not built correctly")
	}
	if house.roof != "tile roof" {
		t.Error("Roof not built correctly")
	}
}
