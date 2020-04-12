### geo2d

Golang util library with basic 2D geo structs and functions.

Example:

```go
import(
	"fmt"
	"math"

	geo "github.com/rogerpales/geo2d"
)

func main() {
	v1 := geo.Vector{P1: geo.Point{0.0, 0.0}, P2: geo.Point{2.0, 2.0}}
	v2 := geo.Vector{P1: geo.Point{0.0, 2.0}, P2: geo.Point{2.0, 0.0}}
	v3 := geo.Vector{P1: geo.Point{0.0, 5.0}, P2: geo.Point{2.0, 3.0}}
	fmt.Printf("v1 and v2 intersect: %v\n", v1.Intersect(v2))
	fmt.Printf("v1 and v3 intersect: %v\n", v1.Intersect(v3))

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