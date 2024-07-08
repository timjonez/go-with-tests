package main

import "testing"

func TestParimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := Parimeter(rect)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12.0, 6.0}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, test := range areaTests {
		got := test.shape.Area()
		if got != test.want {
			t.Errorf("%#v got %g want %g", test.shape, got, test.want)
		}
	}
}
