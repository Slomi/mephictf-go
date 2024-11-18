package shapes

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

const delta = 0.000001

func TestRectangleArea(t *testing.T) {
	tests := []struct {
		name     string
		rect     Rectangle
		expected float64
	}{
		{
			name:     "positive_dimensions",
			rect:     Rectangle{Width: 4, Height: 5},
			expected: 20,
		},
		{
			name:     "zero_dimensions",
			rect:     Rectangle{Width: 0, Height: 0},
			expected: 0,
		},
		{
			name:     "large_dimensions",
			rect:     Rectangle{Width: 100, Height: 200},
			expected: 20000,
		},
		{
			name:     "unequal_dimensions",
			rect:     Rectangle{Width: 2, Height: 8},
			expected: 16,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.rect.Area()
			require.InDelta(t, tc.expected, got, delta)
		})
	}
}

func TestRectanglePerimeter(t *testing.T) {
	tests := []struct {
		name     string
		rect     Rectangle
		expected float64
	}{
		{
			name:     "positive_dimensions",
			rect:     Rectangle{Width: 4, Height: 5},
			expected: 18,
		},
		{
			name:     "zero_dimensions",
			rect:     Rectangle{Width: 0, Height: 0},
			expected: 0,
		},
		{
			name:     "square_shape",
			rect:     Rectangle{Width: 5, Height: 5},
			expected: 20,
		},
		{
			name:     "long_rectangle",
			rect:     Rectangle{Width: 1, Height: 10},
			expected: 22,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.rect.Perimeter()
			require.InDelta(t, tc.expected, got, delta)
		})
	}
}

func TestCircleArea(t *testing.T) {
	tests := []struct {
		name     string
		circle   Circle
		expected float64
	}{
		{
			name:     "positive_radius",
			circle:   Circle{Radius: 5},
			expected: math.Pi * 25,
		},
		{
			name:     "zero_radius",
			circle:   Circle{Radius: 0},
			expected: 0,
		},
		{
			name:     "unit_circle",
			circle:   Circle{Radius: 1},
			expected: math.Pi,
		},
		{
			name:     "large_radius",
			circle:   Circle{Radius: 100},
			expected: math.Pi * 10000,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.circle.Area()
			require.InDelta(t, tc.expected, got, delta)
		})
	}
}

func TestCirclePerimeter(t *testing.T) {
	tests := []struct {
		name     string
		circle   Circle
		expected float64
	}{
		{
			name:     "positive_radius",
			circle:   Circle{Radius: 5},
			expected: 2 * math.Pi * 5,
		},
		{
			name:     "zero_radius",
			circle:   Circle{Radius: 0},
			expected: 0,
		},
		{
			name:     "unit_circle",
			circle:   Circle{Radius: 1},
			expected: 2 * math.Pi,
		},
		{
			name:     "small_radius",
			circle:   Circle{Radius: 0.5},
			expected: math.Pi,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.circle.Perimeter()
			require.InDelta(t, tc.expected, got, delta)
		})
	}
}

func TestTriangleArea(t *testing.T) {
	tests := []struct {
		name     string
		triangle Triangle
		expected float64
	}{
		{
			name:     "right_triangle",
			triangle: Triangle{SideA: 3, SideB: 4, SideC: 5},
			expected: 6,
		},
		{
			name:     "zero_sides",
			triangle: Triangle{SideA: 0, SideB: 0, SideC: 0},
			expected: 0,
		},
		{
			name:     "equilateral_triangle",
			triangle: Triangle{SideA: 5, SideB: 5, SideC: 5},
			expected: 10.825317547305483,
		},
		{
			name:     "isosceles_triangle",
			triangle: Triangle{SideA: 5, SideB: 5, SideC: 8},
			expected: 12,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.triangle.Area()
			require.InDelta(t, tc.expected, got, delta)
		})
	}
}

func TestTrianglePerimeter(t *testing.T) {
	tests := []struct {
		name     string
		triangle Triangle
		expected float64
	}{
		{
			name:     "right_triangle",
			triangle: Triangle{SideA: 3, SideB: 4, SideC: 5},
			expected: 12,
		},
		{
			name:     "zero_sides",
			triangle: Triangle{SideA: 0, SideB: 0, SideC: 0},
			expected: 0,
		},
		{
			name:     "equilateral_triangle",
			triangle: Triangle{SideA: 5, SideB: 5, SideC: 5},
			expected: 15,
		},
		{
			name:     "scalene_triangle",
			triangle: Triangle{SideA: 2, SideB: 3, SideC: 4},
			expected: 9,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.triangle.Perimeter()
			require.InDelta(t, tc.expected, got, delta)
		})
	}
}

func TestCalculateTotalArea(t *testing.T) {
	shapes := []Shape{
		Rectangle{Width: 4, Height: 3},
		Circle{Radius: 7},
		Triangle{SideA: 3, SideB: 4, SideC: 5},
	}

	got := CalculateTotalArea(shapes)
	expected := 12.0 + (math.Pi * 49) + 6.0

	require.InDelta(t, expected, got, delta)
}

func TestAveragePerimeter(t *testing.T) {
	shapes := []Shape{
		Rectangle{Width: 2, Height: 3},
		Circle{Radius: 7},
		Triangle{SideA: 3, SideB: 4, SideC: 5},
	}

	got := AveragePerimeter(shapes)
	expected := (10.0 + (2 * math.Pi * 7) + 12.0) / 3

	require.InDelta(t, expected, got, delta)
}
