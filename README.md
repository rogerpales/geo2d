### geo2d

Golang util library with basic 2D geo structs and functions.

Example:

```go
import(
	"fmt"
	"math"

	geo "github.com/rogerpales/geo2d"
)

import (
	"fmt"
	"math"

	geo "github.com/rogerpales/geo2d"
)

func main() {
	s1 := geo.Segment{P1: geo.Point{0.0, 0.0}, P2: geo.Point{2.0, 2.0}}
	s2 := geo.Segment{P1: geo.Point{0.0, 2.0}, P2: geo.Point{2.0, 0.0}}
	s3 := geo.Segment{P1: geo.Point{0.0, 5.0}, P2: geo.Point{2.0, 3.0}}
	fmt.Printf("s1 and s2 intersect: %v\n", s1.Intersect(s2))
	fmt.Printf("s1 and s3 intersect: %v\n", s1.Intersect(s3))

	p1 := geo.Point{0.0, 0.0}
	p2 := geo.Point{2.0, 0.0}
	p3 := geo.Point{2.0, 2.0}

	t := geo.NewTriangle(p1, p2, p3)

	angles, _ := t.GetAngles()
	for i, a := range angles {
		nextIndex := (i + 1) % len(angles)
		desc := fmt.Sprintf("opposite to side %v->%v", t.Vertices[i], t.Vertices[nextIndex])
		fmt.Printf("%s: %.2f rad ≈ %.2f°\n", desc, a, a*180/math.Pi)
	}

	t2 := geo.NewTriangle(p1, p2, p3) // copy

	v := geo.NewVector(3, 4)
	t.Translate(v)

	desc := "triangle vx transl %.2f,%.2f -> %.2f,%.2f\n"
	for i, p := range t.GetVertices() {
		fmt.Printf(desc, t2.Vertices[i].X, t2.Vertices[i].Y, p.X, p.Y)
	}

	t = geo.NewTriangle(p1, p2, p3)

	rotateCenter := p3
	rotateAngle := math.Pi

	t.Rotate(rotateCenter, rotateAngle)

	desc = "triangle vx rotate %.2f,%.2f -> %.2f,%.2f\n"
	for i, p := range t.GetVertices() {
		fmt.Printf(desc, t2.Vertices[i].X, t2.Vertices[i].Y, p.X, p.Y)
	}
}
```

```
v1 and v2 intersect: true
v1 and v3 intersect: false
opposite to side {0 0}->{2 0}: 0.79 rad ≈ 45.00°
opposite to side {2 0}->{2 2}: 0.79 rad ≈ 45.00°
opposite to side {2 2}->{0 0}: 1.57 rad ≈ 90.00°
```

### Some types

* Point
* Vector
* Line
* Polygon
* Triangle

### Some functions

* Rotate
* Intersection
* Magnitude
* New regular polygon