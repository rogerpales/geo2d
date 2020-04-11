package geo2d

import (
	"errors"
	"fmt"
	"math"
)

// Point reprsents a 2D point
type Point struct {
	X, Y float64
}

// Vector represents a 2D vector
type Vector struct {
	P1, P2 Point
}

// Line represents a 2D line
type Line struct {
	Slope float64
	Yint  float64
}

// Polygon represents a 2D polygon
type Polygon struct {
	Vertices []Point
}

// Coord returns point coordinates x, y
func (p Point) Coord() (x, y float64) {
	return p.X, p.Y
}

// Rotate point a given angle from given center
func (p *Point) Rotate(center Point, angleRad float64) {
	x := p.X*math.Cos(angleRad) - p.Y*math.Sin(angleRad)
	y := p.Y*math.Cos(angleRad) - p.X*math.Sin(angleRad)

	p.X = x
	p.Y = y
}

// Triangle represents a triangle polygon
type Triangle struct {
	Vertices [3]Point
}

// NewTriangle returns a triangle
func NewTriangle(v1, v2, v3 Point) Triangle {
	return Triangle{Vertices: [3]Point{v1, v2, v3}}
}

var errTrinagleValid = errors.New("Triangle polygon must have 3 vertices")

// GetAngles returns the 3 angles ordered as follows:
//		result[0] is the angle opposite to side from t.Vertices[0] to t.Vertices[1]
//		result[1] is the angle opposite to side from t.Vertices[1] to t.Vertices[2]
// 		result[2] is the angle opposite to side from t.Vertices[2] to t.Vertices[0]
func (t Triangle) GetAngles() (angles [3]float64, err error) {
	for i, v := range t.Vertices {
		for i2, v2 := range t.Vertices {
			if i == i2 {
				continue
			}
			if v == v2 {
				fmt.Println(v)
				fmt.Println(v2)
				err = errTrinagleValid
				return
			}
		}
	}

	a := Vector{t.Vertices[0], t.Vertices[1]}.Magnitude()
	b := Vector{t.Vertices[1], t.Vertices[2]}.Magnitude()
	c := Vector{t.Vertices[2], t.Vertices[0]}.Magnitude()

	angles[0] = math.Acos((math.Pow(b, 2) + math.Pow(c, 2) - math.Pow(a, 2)) / (2 * b * c))
	angles[1] = math.Acos((math.Pow(a, 2) + math.Pow(c, 2) - math.Pow(b, 2)) / (2 * a * c))
	angles[2] = math.Acos((math.Pow(a, 2) + math.Pow(b, 2) - math.Pow(c, 2)) / (2 * a * b))
	return
}

// NewLine returns the line intersecting points a, b
func NewLine(a, b Point) Line {
	slope := (b.Y - a.Y) / (b.X - a.X)
	yAtXo := a.Y - slope*a.X
	return Line{slope, yAtXo}
}

// GetY returns y value on line (l) given x
func (l Line) GetY(x float64) float64 {
	return l.Slope*x + l.Yint
}

// Intersection returns the intersection Point of line (l) with another line (l2)
func (l Line) Intersection(l2 Line) (Point, error) {
	if l.Slope == l2.Slope {
		return Point{}, errors.New("The lines do not intersect")
	}
	x := (l2.Yint - l.Yint) / (l.Slope - l2.Slope)
	y := l.GetY(x)
	return Point{x, y}, nil
}

// Intersect returns wether or not two given vectors intersect
func (v Vector) Intersect(u Vector) bool {
	vl := NewLine(v.P1, v.P2)
	ul := NewLine(u.P1, u.P2)

	iP, err := vl.Intersection(ul)
	if err != nil {
		return false
	}

	return (iP.X >= v.P1.X && iP.X <= v.P2.X ||
		iP.X >= v.P2.X && iP.X <= v.P1.X) &&
		(iP.Y >= v.P1.Y && iP.Y <= v.P2.Y ||
			iP.Y >= v.P2.Y && iP.Y <= v.P1.Y)
}

// Length is returns vector (v) magnitude
func (v Vector) Length() float64 {
	return v.Magnitude()
}

// Magnitude is the distance between the two ends of a vector (v)
func (v Vector) Magnitude() float64 {
	deltaX := v.P2.X - v.P1.X
	deltaY := v.P2.Y - v.P1.Y

	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

// NewRegularPolygon returns a new polygon with vNum number of
// vertices and with given vertex and center points
func NewRegularPolygon(vertex, center Point, vNum int) (p Polygon, err error) {
	raidusVect := Vector{vertex, center}
	radius := raidusVect.Magnitude()

	p, err = newRegularPolygonWithCenter(center, radius, vNum)
	if err != nil {
		return
	}

	offsetVect := Vector{p.Vertices[0], vertex}
	for i := range p.Vertices {
		p.Vertices[i].X = p.Vertices[i].X + (offsetVect.P1.X - offsetVect.P2.X)
		p.Vertices[i].Y = p.Vertices[i].Y + (offsetVect.P1.Y - offsetVect.P2.Y)
	}

	return
}

func newRegularPolygonWithCenter(center Point, radius float64, vNum int) (p Polygon, err error) {
	if vNum < 3 {
		err = errors.New("Polygon has a minimum of 3 vertices")
		return
	}

	for i := 0; i < vNum; i++ {
		rate := float64(i) / float64(vNum)

		x := float64(radius*math.Cos(2*math.Pi*rate)) + center.X
		y := float64(radius*math.Sin(2*math.Pi*rate)) + center.Y

		p.Vertices = append(p.Vertices, Point{X: x, Y: y})
	}

	return
}

// Rotate polygon (vertices) a given angle from given center
func (p *Polygon) Rotate(center Point, angleRad float64) {
	for _, v := range p.Vertices {
		v.Rotate(center, angleRad)
	}
}
