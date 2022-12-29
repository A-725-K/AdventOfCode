package ds

import (
  "math"
  "strings"
  "strconv"
)

type Comparable interface {
	int | int8 | int16 | int32 | int64 |
	uint | uint8 | uint16 | uint32 | uint64 |
	uintptr |
	float32 | float64
}

func Abs[T Comparable](x T) T {
  if x >= 0 {
    return x
  }
  return -x
}

func Max[T Comparable](x, y T) T {
  return T(math.Max(float64(x), float64(y)))
}

func Min[T Comparable](x, y T) T {
  return T(math.Min(float64(x), float64(y)))
}

func Sign[T Comparable](x T) T {
  if x == 0 {
    return 0
  }
  return x / Abs(x)
}

func Divmod(numerator, denominator int64) (quotient, remainder int64) {
	quotient = numerator / denominator
	remainder = numerator % denominator
	return
}

func Mod(x, mod int64) int64 {
  res := x % mod
  if res < 0 && mod > 0 {
    res += mod
  }
  return res
}

func ToKey(x, y int) string {
  return strconv.Itoa(x) + "," + strconv.Itoa(y)
}

func FromKey(k string) (int, int) {
  vs := strings.Split(k, ",")
  v0, err := strconv.Atoi(vs[0])
  if err != nil {
    panic("Cannot convert key to values")
  }
  v1, err := strconv.Atoi(vs[1])
  if err != nil {
    panic("Cannot convert key to values")
  }
  return v0, v1
}

func ToKey64(x, y int64) string {
  return strconv.FormatInt(x, 10) + "," + strconv.FormatInt(y, 10)
}

func FromKey64(k string) (int64, int64) {
  vs := strings.Split(k, ",")
  v0, err := strconv.ParseInt(vs[0], 10, 64)
  if err != nil {
    panic("Cannot convert key to values")
  }
  v1, err := strconv.ParseInt(vs[1], 10, 64)
  if err != nil {
    panic("Cannot convert key to values")
  }
  return v0, v1
}

// Gcd: Greatest Common Divisor via Euclidean algorithm
func Gcd(a, b int) int {
  for b != 0 {
    t := b
    b = a % b
    a = t
  }
  return a
}

// Lcm: find Least Common Multiple via Gcd
func Lcm(a, b int, integers ...int) int {
  result := a * b / Gcd(a, b)

  for i := 0; i < len(integers); i++ {
    result = Lcm(result, integers[i])
  }

  return result
}
