package shapes

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

// table driven test
func TestArea(t *testing.T) {
	areaTests := map[string]struct {
		shape Shape
		want  float64
	}{
		"Rectangle": {shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		"Circle":    {shape: Circle{Radius: 10}, want: 314.1592653589793},
		"Triangle":  {shape: Triangle{Base: 10, Height: 3}, want: 15},
	}
	for name, tt := range areaTests {
		t.Run(name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		})
	}
}
