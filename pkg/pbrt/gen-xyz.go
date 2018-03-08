// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package pbrt

import (
	"fmt"
	"math"
)

type XYInt64 struct {
	X, Y int64
}

func (xy *XYInt64) GetIndex(i int) int64 {
	switch i {
	case 0:
		return xy.X
	default:
		return xy.Y
	}
}

func (xy *XYInt64) SetIndex(i int, v int64) {
	switch i {
	case 0:
		xy.X = v
	default:
		xy.Y = v
	}
}

func (xy *XYInt64) Set(x, y int64) *XYInt64 {
	xy.X = x
	xy.Y = y
	return xy
}

func (xy *XYInt64) LengthSquared() int64 {
	return xy.X*xy.X + xy.Y*xy.Y
}

func (xy *XYInt64) Length() int64 {
	return int64(math.Sqrt(float64(xy.LengthSquared())))
}

func (xy *XYInt64) String() string {
	return fmt.Sprintf("XYint64(%f.3, %f.3)", xy.X, xy.Y)
}

func (xy *XYInt64) Add(other *XYInt64) *XYInt64 {
	return &XYInt64{xy.X + other.X, xy.Y + other.Y}
}

func (xy *XYInt64) AddAssign(other *XYInt64) {
	xy.X += other.X
	xy.Y += other.Y
}

func (xy *XYInt64) AddConst(other int64) *XYInt64 {
	return &XYInt64{xy.X + other, xy.Y + other}
}

func (xy *XYInt64) Sub(other *XYInt64) *XYInt64 {
	return &XYInt64{xy.X - other.X, xy.Y - other.Y}
}

func (xy *XYInt64) SubAssign(other *XYInt64) {
	xy.X -= other.X
	xy.Y -= other.Y
}

func (xy *XYInt64) Mul(other *XYInt64) *XYInt64 {
	return &XYInt64{xy.X * other.X, xy.Y * other.Y}
}

func (xy *XYInt64) MulAssign(other *XYInt64) {
	xy.X *= other.X
	xy.Y *= other.Y
}

func (xy *XYInt64) MulScalar(other int64) *XYInt64 {
	return &XYInt64{xy.X * other, xy.Y * other}
}

func (xy *XYInt64) Div(other *XYInt64) *XYInt64 {
	return &XYInt64{xy.X / other.X, xy.Y / other.Y}
}

func (xy *XYInt64) DivAssign(other *XYInt64) {
	xy.X /= other.X
	xy.Y /= other.Y
}

func (xy *XYInt64) Equals(other *XYInt64) bool {
	return xy.X == other.X && xy.Y == other.Y
}

func (xy *XYInt64) NotEquals(other *XYInt64) bool {
	return xy.X != other.X || xy.Y != other.Y
}

func (xy *XYInt64) Floor() *XYInt64 {
	return &XYInt64{
		X: int64(math.Floor(float64(xy.X))),
		Y: int64(math.Floor(float64(xy.Y))),
	}
}

func (xy *XYInt64) Ceil() *XYInt64 {
	return &XYInt64{
		X: int64(math.Ceil(float64(xy.X))),
		Y: int64(math.Ceil(float64(xy.Y))),
	}
}

type XYZInt64 struct {
	X, Y, Z int64
}

func (xyz *XYZInt64) Index(i int) int64 {
	switch i {
	case 0:
		return xyz.X
	case 1:
		return xyz.Y
	default:
		return xyz.Z
	}
}

func (xyz *XYZInt64) SetIndex(i int, v int64) {
	switch i {
	case 0:
		xyz.X = v
	case 1:
		xyz.Y = v
	default:
		xyz.Z = v
	}
}

func (xyz *XYZInt64) Set(x, y, z int64) *XYZInt64 {
	xyz.X = x
	xyz.Y = y
	xyz.Z = z
	return xyz
}

func (xyz *XYZInt64) LengthSquared() int64 {
	return xyz.X*xyz.X + xyz.Y*xyz.Y + xyz.Z*xyz.Z
}

func (xyz *XYZInt64) Length() int64 {
	return int64(math.Sqrt(float64(xyz.LengthSquared())))
}

func (xyz *XYZInt64) String() string {
	return fmt.Sprintf("XYZint64(%f.3, %f.3, %f.3)", xyz.X, xyz.Y, xyz.Z)
}

func (xyz *XYZInt64) Add(other *XYZInt64) *XYZInt64 {
	return &XYZInt64{
		xyz.X + other.X,
		xyz.Y + other.Y,
		xyz.Z + other.Z,
	}
}

func (xyz *XYZInt64) AddAssign(other *XYZInt64) {
	xyz.X += other.X
	xyz.Y += other.Y
	xyz.Z += other.Z
}

func (xyz *XYZInt64) AddConst(other int64) *XYZInt64 {
	return &XYZInt64{xyz.X + other, xyz.Y + other, xyz.Z + other}
}

func (xyz *XYZInt64) Sub(other *XYZInt64) *XYZInt64 {
	return &XYZInt64{
		xyz.X - other.X,
		xyz.Y - other.Y,
		xyz.Z - other.Z,
	}
}

func (xyz *XYZInt64) SubAssign(other *XYZInt64) {
	xyz.X += other.X
	xyz.Y += other.Y
	xyz.Z += other.Z
}

func (xyz *XYZInt64) Mul(other *XYZInt64) *XYZInt64 {
	return &XYZInt64{
		xyz.X * other.X,
		xyz.Y * other.Y,
		xyz.Z * other.Z,
	}
}

func (xyz *XYZInt64) MulAssign(other *XYZInt64) {
	xyz.X += other.X
	xyz.Y += other.Y
	xyz.Z += other.Z
}

func (xyz *XYZInt64) MulScalar(other int64) *XYZInt64 {
	return &XYZInt64{
		xyz.X * other,
		xyz.Y * other,
		xyz.Z * other,
	}
}

func (xyz *XYZInt64) Div(other *XYZInt64) *XYZInt64 {
	return &XYZInt64{
		xyz.X / other.X,
		xyz.Y / other.Y,
		xyz.Z / other.Z,
	}
}

func (xyz *XYZInt64) DivAssign(other *XYZInt64) {
	xyz.X += other.X
	xyz.Y += other.Y
	xyz.Z += other.Z
}

func (xyz *XYZInt64) DivScalar(other int64) *XYZInt64 {
	return &XYZInt64{
		xyz.X / other,
		xyz.Y / other,
		xyz.Z / other,
	}
}

func (xyz *XYZInt64) Abs() *XYZInt64 {
	return &XYZInt64{
		int64(math.Abs(float64(xyz.X))),
		int64(math.Abs(float64(xyz.Y))),
		int64(math.Abs(float64(xyz.Z))),
	}
}

func (xyz *XYZInt64) Dot(other *XYZInt64) int64 {
	return xyz.X*other.X + xyz.Y*other.Y + xyz.Z*other.Z
}

func (xyz *XYZInt64) AbsDot(other *XYZInt64) int64 {
	return int64(math.Abs(float64(xyz.Dot(other))))
}

func (xyz *XYZInt64) Distance(other *XYZInt64) int64 {
	return int64(math.Sqrt(float64(xyz.Dot(other))))
}

func (xyz *XYZInt64) Cross(other *XYZInt64) *XYZInt64 {
	return &XYZInt64{(xyz.Y * other.Z) - (xyz.Z * other.Y), (xyz.Z * other.X) - (xyz.X * other.Z), (xyz.X * other.Y) - (xyz.Y * other.X)}
}

func (xyz *XYZInt64) Normalize() {
	nor2 := xyz.LengthSquared()
	if nor2 > 0 {
		invNor := int64(1) / int64(math.Sqrt(float64(nor2)))
		xyz.X *= invNor
		xyz.Y *= invNor
		xyz.Z *= invNor
	}
}

func (xyz XYZInt64) Normalized() *XYZInt64 {
	nor2 := xyz.LengthSquared()
	if nor2 > 0 {
		invNor := int64(1) / int64(math.Sqrt(float64(nor2)))
		xyz.X *= invNor
		xyz.Y *= invNor
		xyz.Z *= invNor
	}
	return &xyz
}

func (xyz *XYZInt64) Equals(other *XYZInt64) bool {
	return xyz.X == other.X && xyz.Y == other.Y && xyz.Z == other.Z
}

func (xyz *XYZInt64) NotEquals(other *XYZInt64) bool {
	return xyz.X != other.X || xyz.Y != other.Y || xyz.Z != other.Z
}

type XYFloat64 struct {
	X, Y float64
}

func (xy *XYFloat64) GetIndex(i int) float64 {
	switch i {
	case 0:
		return xy.X
	default:
		return xy.Y
	}
}

func (xy *XYFloat64) SetIndex(i int, v float64) {
	switch i {
	case 0:
		xy.X = v
	default:
		xy.Y = v
	}
}

func (xy *XYFloat64) Set(x, y float64) *XYFloat64 {
	xy.X = x
	xy.Y = y
	return xy
}

func (xy *XYFloat64) LengthSquared() float64 {
	return xy.X*xy.X + xy.Y*xy.Y
}

func (xy *XYFloat64) Length() float64 {
	return float64(math.Sqrt(float64(xy.LengthSquared())))
}

func (xy *XYFloat64) String() string {
	return fmt.Sprintf("XYfloat64(%f.3, %f.3)", xy.X, xy.Y)
}

func (xy *XYFloat64) Add(other *XYFloat64) *XYFloat64 {
	return &XYFloat64{xy.X + other.X, xy.Y + other.Y}
}

func (xy *XYFloat64) AddAssign(other *XYFloat64) {
	xy.X += other.X
	xy.Y += other.Y
}

func (xy *XYFloat64) AddConst(other float64) *XYFloat64 {
	return &XYFloat64{xy.X + other, xy.Y + other}
}

func (xy *XYFloat64) Sub(other *XYFloat64) *XYFloat64 {
	return &XYFloat64{xy.X - other.X, xy.Y - other.Y}
}

func (xy *XYFloat64) SubAssign(other *XYFloat64) {
	xy.X -= other.X
	xy.Y -= other.Y
}

func (xy *XYFloat64) Mul(other *XYFloat64) *XYFloat64 {
	return &XYFloat64{xy.X * other.X, xy.Y * other.Y}
}

func (xy *XYFloat64) MulAssign(other *XYFloat64) {
	xy.X *= other.X
	xy.Y *= other.Y
}

func (xy *XYFloat64) MulScalar(other float64) *XYFloat64 {
	return &XYFloat64{xy.X * other, xy.Y * other}
}

func (xy *XYFloat64) Div(other *XYFloat64) *XYFloat64 {
	return &XYFloat64{xy.X / other.X, xy.Y / other.Y}
}

func (xy *XYFloat64) DivAssign(other *XYFloat64) {
	xy.X /= other.X
	xy.Y /= other.Y
}

func (xy *XYFloat64) Equals(other *XYFloat64) bool {
	return xy.X == other.X && xy.Y == other.Y
}

func (xy *XYFloat64) NotEquals(other *XYFloat64) bool {
	return xy.X != other.X || xy.Y != other.Y
}

func (xy *XYFloat64) Floor() *XYFloat64 {
	return &XYFloat64{
		X: float64(math.Floor(float64(xy.X))),
		Y: float64(math.Floor(float64(xy.Y))),
	}
}

func (xy *XYFloat64) Ceil() *XYFloat64 {
	return &XYFloat64{
		X: float64(math.Ceil(float64(xy.X))),
		Y: float64(math.Ceil(float64(xy.Y))),
	}
}

type XYZFloat64 struct {
	X, Y, Z float64
}

func (xyz *XYZFloat64) Index(i int) float64 {
	switch i {
	case 0:
		return xyz.X
	case 1:
		return xyz.Y
	default:
		return xyz.Z
	}
}

func (xyz *XYZFloat64) SetIndex(i int, v float64) {
	switch i {
	case 0:
		xyz.X = v
	case 1:
		xyz.Y = v
	default:
		xyz.Z = v
	}
}

func (xyz *XYZFloat64) Set(x, y, z float64) *XYZFloat64 {
	xyz.X = x
	xyz.Y = y
	xyz.Z = z
	return xyz
}

func (xyz *XYZFloat64) LengthSquared() float64 {
	return xyz.X*xyz.X + xyz.Y*xyz.Y + xyz.Z*xyz.Z
}

func (xyz *XYZFloat64) Length() float64 {
	return float64(math.Sqrt(float64(xyz.LengthSquared())))
}

func (xyz *XYZFloat64) String() string {
	return fmt.Sprintf("XYZfloat64(%f.3, %f.3, %f.3)", xyz.X, xyz.Y, xyz.Z)
}

func (xyz *XYZFloat64) Add(other *XYZFloat64) *XYZFloat64 {
	return &XYZFloat64{
		xyz.X + other.X,
		xyz.Y + other.Y,
		xyz.Z + other.Z,
	}
}

func (xyz *XYZFloat64) AddAssign(other *XYZFloat64) {
	xyz.X += other.X
	xyz.Y += other.Y
	xyz.Z += other.Z
}

func (xyz *XYZFloat64) AddConst(other float64) *XYZFloat64 {
	return &XYZFloat64{xyz.X + other, xyz.Y + other, xyz.Z + other}
}

func (xyz *XYZFloat64) Sub(other *XYZFloat64) *XYZFloat64 {
	return &XYZFloat64{
		xyz.X - other.X,
		xyz.Y - other.Y,
		xyz.Z - other.Z,
	}
}

func (xyz *XYZFloat64) SubAssign(other *XYZFloat64) {
	xyz.X += other.X
	xyz.Y += other.Y
	xyz.Z += other.Z
}

func (xyz *XYZFloat64) Mul(other *XYZFloat64) *XYZFloat64 {
	return &XYZFloat64{
		xyz.X * other.X,
		xyz.Y * other.Y,
		xyz.Z * other.Z,
	}
}

func (xyz *XYZFloat64) MulAssign(other *XYZFloat64) {
	xyz.X += other.X
	xyz.Y += other.Y
	xyz.Z += other.Z
}

func (xyz *XYZFloat64) MulScalar(other float64) *XYZFloat64 {
	return &XYZFloat64{
		xyz.X * other,
		xyz.Y * other,
		xyz.Z * other,
	}
}

func (xyz *XYZFloat64) Div(other *XYZFloat64) *XYZFloat64 {
	return &XYZFloat64{
		xyz.X / other.X,
		xyz.Y / other.Y,
		xyz.Z / other.Z,
	}
}

func (xyz *XYZFloat64) DivAssign(other *XYZFloat64) {
	xyz.X += other.X
	xyz.Y += other.Y
	xyz.Z += other.Z
}

func (xyz *XYZFloat64) DivScalar(other float64) *XYZFloat64 {
	return &XYZFloat64{
		xyz.X / other,
		xyz.Y / other,
		xyz.Z / other,
	}
}

func (xyz *XYZFloat64) Abs() *XYZFloat64 {
	return &XYZFloat64{
		float64(math.Abs(float64(xyz.X))),
		float64(math.Abs(float64(xyz.Y))),
		float64(math.Abs(float64(xyz.Z))),
	}
}

func (xyz *XYZFloat64) Dot(other *XYZFloat64) float64 {
	return xyz.X*other.X + xyz.Y*other.Y + xyz.Z*other.Z
}

func (xyz *XYZFloat64) AbsDot(other *XYZFloat64) float64 {
	return float64(math.Abs(float64(xyz.Dot(other))))
}

func (xyz *XYZFloat64) Distance(other *XYZFloat64) float64 {
	return float64(math.Sqrt(float64(xyz.Dot(other))))
}

func (xyz *XYZFloat64) Cross(other *XYZFloat64) *XYZFloat64 {
	return &XYZFloat64{(xyz.Y * other.Z) - (xyz.Z * other.Y), (xyz.Z * other.X) - (xyz.X * other.Z), (xyz.X * other.Y) - (xyz.Y * other.X)}
}

func (xyz *XYZFloat64) Normalize() {
	nor2 := xyz.LengthSquared()
	if nor2 > 0 {
		invNor := float64(1) / float64(math.Sqrt(float64(nor2)))
		xyz.X *= invNor
		xyz.Y *= invNor
		xyz.Z *= invNor
	}
}

func (xyz XYZFloat64) Normalized() *XYZFloat64 {
	nor2 := xyz.LengthSquared()
	if nor2 > 0 {
		invNor := float64(1) / float64(math.Sqrt(float64(nor2)))
		xyz.X *= invNor
		xyz.Y *= invNor
		xyz.Z *= invNor
	}
	return &xyz
}

func (xyz *XYZFloat64) Equals(other *XYZFloat64) bool {
	return xyz.X == other.X && xyz.Y == other.Y && xyz.Z == other.Z
}

func (xyz *XYZFloat64) NotEquals(other *XYZFloat64) bool {
	return xyz.X != other.X || xyz.Y != other.Y || xyz.Z != other.Z
}
