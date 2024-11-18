package shapes

// Shape interface defines methods that all shapes must implement.
type Shape interface {
	// Area calculates and returns the area of the shape.
	Area() float64

	// Perimeter calculates and returns the perimeter of the shape.
	Perimeter() float64
}

// Rectangle represents a rectangle shape with width and height.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the area of a rectangle.
func (r Rectangle) Area() float64 {
	// TODO: Implement this method.
	return 0
}

// Perimeter calculates the perimeter of a rectangle.
func (r Rectangle) Perimeter() float64 {
	// TODO: Implement this method.
	return 0
}

// Circle represents a circle shape with a radius.
type Circle struct {
	Radius float64
}

// Area calculates the area of a circle.
func (c Circle) Area() float64 {
	// TODO: Implement this method.
	return 0
}

// Perimeter calculates the perimeter (circumference) of a circle.
func (c Circle) Perimeter() float64 {
	// TODO: Implement this method.
	return 0
}

// Triangle represents a triangle shape with three sides.
type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

// Area calculates the area of a triangle.
func (t Triangle) Area() float64 {
	// TODO: Implement this method.
	return 0
}

// Perimeter calculates the perimeter of a triangle.
func (t Triangle) Perimeter() float64 {
	// TODO: Implement this method.
	return 0
}

// CalculateTotalArea takes a slice of shapes and returns their combined area.
func CalculateTotalArea(shapes []Shape) float64 {
	// TODO: Implement this function.
	return 0
}

// AveragePerimeter calculates the average perimeter of a slice of shapes.
func AveragePerimeter(shapes []Shape) float64 {
	// TODO: Implement this function.
	return 0
}
