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
	p1 := Point{0.0, 0.0}
	p2 := Point{2.0, 0.0}
	p3 := Point{2.0, 2.0}

	t := NewTriangle(p1, p2, p3)

	angles, _ := t.GetAngles()
	for _, a := range angles {
		fmt.Printf("%.2f rad ≈ %.2f degree\n", a, a*180/math.Pi)
	}
}
```

```
0.79 rad ≈ 45.00 degree
0.79 rad ≈ 45.00 degree
1.57 rad ≈ 90.00 degree
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