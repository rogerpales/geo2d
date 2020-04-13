package geo2d

import (
	"errors"
	"math"
)

var errLineIntersect = errors.New("lines do not intersect")
var errPolygonMinVex = errors.New("polygon has a minimum of 3 (different) vertices")
var errTriangleNuVex = errors.New("triangle must have exactly 3 (different) vertices")

// Point reprsents a 2D point
type Point struct {
	X, Y float64
}

// Coord returns point coordinates x, y
func (p Point) Coord() (x, y float64) {
	return p.X, p.Y
}

// Rotate point a given angle from given center
func (p *Point) Rotate(center Point, angle float64) {
	s := math.Sin(angle)
	c := math.Cos(angle)

	p.X -= center.X
	p.Y -= center.Y

	x := p.X*c - p.Y*s
	y := p.X*s + p.Y*c

	p.X = x + center.X
	p.Y = y + center.Y
}

// Translate point a given Segment
func (p *Point) Translate(v Vector) {
	p.X += v.X
	p.Y += v.Y
}

// Vector reprsents a 2D point
type Vector struct {
	X, Y float64
}

// NewVector returns a new Segment from origin (P1 0,0) to
// given coord x and y (P2 x,y)
func NewVector(x, y float64) Vector {
	return Vector{X: x, Y: y}
}

// Magnitude is the distance between the two ends of a Segment (v)
func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Segment represents a 2D Segment
type Segment struct {
	P1, P2 Point
}

// Intersect returns wether or not two given Segments intersect
func (s Segment) Intersect(u Segment) bool {
	vl := NewLine(s.P1, s.P2)
	ul := NewLine(u.P1, u.P2)

	iP, err := vl.Intersection(ul)
	if err != nil {
		return false
	}

	return (iP.X >= s.P1.X && iP.X <= s.P2.X ||
		iP.X >= s.P2.X && iP.X <= s.P1.X) &&
		(iP.Y >= s.P1.Y && iP.Y <= s.P2.Y ||
			iP.Y >= s.P2.Y && iP.Y <= s.P1.Y)
}

// Length is returns Segment (v) magnitude
func (s Segment) Length() float64 {
	return s.Magnitude()
}

// Magnitude is the distance between the two ends of a Segment (v)
func (s Segment) Magnitude() float64 {
	deltaX := s.P2.X - s.P1.X
	deltaY := s.P2.Y - s.P1.Y
	return Vector{X: deltaX, Y: deltaY}.Magnitude()
}

// Path represents a 2D path (connected points)
type Path struct {
	Vertices []Point
}

// GetVertices returns a slice of points
func (p Path) GetVertices() (vs []Point) {
	return p.Vertices
}

// Translate all path vertices a given vector
func (p *Path) Translate(v Vector) {
	for i, vx := range p.Vertices {
		vx.Translate(v)
		p.Vertices[i] = vx
	}
}

// GetSides returns a slice of Segment (polygon sides)
func (p Path) GetSides() (sides []Segment) {
	for i, v := range p.Vertices {
		nextIndex := i + 1
		if nextIndex == len(p.Vertices) {
			continue
		}
		sides = append(sides, Segment{v, p.Vertices[nextIndex]})
	}
	return
}

// Line represents a 2D line
type Line struct {
	Slope float64
	Yint  float64
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
		return Point{}, errLineIntersect
	}
	x := (l2.Yint - l.Yint) / (l.Slope - l2.Slope)
	y := l.GetY(x)
	return Point{x, y}, nil
}

// Polygon represents a 2D polygon
type Polygon struct {
	Path
}

// NewRegularPolygon returns a new polygon with vNum number of
// vertices and with given vertex and center points
func NewRegularPolygon(center, vertex Point, vNum int) (p Polygon, err error) {
	raidusVect := Segment{vertex, center}
	radius := raidusVect.Magnitude()

	p, err = NewRegularPolygonWithRadius(center, radius, vNum)
	if err != nil {
		return
	}

	t := NewTriangle(center, p.Vertices[0], vertex)
	var angles [3]float64
	angles, err = t.GetAngles()
	if err == nil {
		p.Rotate(center, angles[0])
		// otherwise vertex index 0 is already given vertex
	}

	return
}

// NewRegularPolygonWithRadius returns a new polygon with vNum
// number of vertices and given center point
func NewRegularPolygonWithRadius(center Point, radius float64, vNum int) (p Polygon, err error) {
	if vNum < 3 {
		err = errPolygonMinVex
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
	for i, v := range p.Vertices {
		v.Rotate(center, angleRad)
		p.Vertices[i] = v
	}
}

// GetSides returns a slice of Segment (polygon sides)
func (p Polygon) GetSides() (sides []Segment) {
	for i, v := range p.Vertices {
		nextIndex := (i + 1) % len(p.Vertices)
		sides = append(sides, Segment{v, p.Vertices[nextIndex]})
	}
	return
}

// Triangle represents a triangle polygon
type Triangle struct {
	Polygon
}

// NewTriangle returns a triangle
func NewTriangle(v1, v2, v3 Point) (t Triangle) {
	t.Vertices = []Point{v1, v2, v3}
	return t
}

// GetAngles returns the 3 angles ordered as follows:
//		result[0] is the angle opposite to side from t.Vertices[0] to t.Vertices[1]
//		result[1] is the angle opposite to side from t.Vertices[1] to t.Vertices[2]
// 		result[2] is the angle opposite to side from t.Vertices[2] to t.Vertices[0]
func (t Triangle) GetAngles() (angles [3]float64, err error) {
	if len(t.Vertices) != 3 {
		err = errTriangleNuVex
		return
	}
	for i, v := range t.Vertices {
		for i2, v2 := range t.Vertices {
			if i == i2 {
				continue
			}
			if v == v2 {
				err = errTriangleNuVex
				return
			}
		}
	}

	a := Segment{t.Vertices[0], t.Vertices[1]}.Magnitude()
	b := Segment{t.Vertices[1], t.Vertices[2]}.Magnitude()
	c := Segment{t.Vertices[2], t.Vertices[0]}.Magnitude()

	angles[0] = math.Acos((math.Pow(b, 2) + math.Pow(c, 2) - math.Pow(a, 2)) / (2 * b * c))
	angles[1] = math.Acos((math.Pow(a, 2) + math.Pow(c, 2) - math.Pow(b, 2)) / (2 * a * c))
	angles[2] = math.Acos((math.Pow(a, 2) + math.Pow(b, 2) - math.Pow(c, 2)) / (2 * a * b))
	return
}

// Figure interface implements GetSides and GetVertices
type Figure interface {
	GetSides() []Segment
	GetVertices() []Point
}

// Intersect returns wether or not at least one side of two
// given figures intersect
func Intersect(f1, f2 Figure) bool {
	f2sides := f2.GetSides()
	for _, s1 := range f1.GetSides() {
		for _, s2 := range f2sides {
			if s1.Intersect(s2) {
				return true
			}
		}
	}
	return false
}

// LoHiX returns lowest and highest X axis values
func LoHiX(f Figure) (lo float64, hi float64) {
	for i, v := range f.GetVertices() {
		if i == 0 {
			lo = v.X
			hi = v.X
			continue
		}
		if v.Y < lo {
			lo = v.X
		}
		if v.Y > hi {
			hi = v.X
		}
	}
	return
}

// LoHiY returns lowest and highest Y axis values
func LoHiY(f Figure) (lo float64, hi float64) {
	for i, v := range f.GetVertices() {
		if i == 0 {
			lo = v.Y
			hi = v.Y
			continue
		}
		if v.Y < lo {
			lo = v.Y
		}
		if v.Y > hi {
			hi = v.Y
		}
	}
	return
}
