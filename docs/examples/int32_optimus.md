## Optimus Lib
```go
// NewOptimus opt := NewOptimus(1580030173, 59260789, 1163945558, 31)
func NewOptimus(prime, inverse, random int, bitSize int) *Optimus {
	// https://github.com/jenssegers/optimus
	max := int(math.Pow(2, float64(bitSize)) - 1)
	return &Optimus{
		prime:   prime,
		inverse: inverse,
		xor:     random,
		max:     max,
	}
}

type Optimus struct {
	prime   int
	inverse int
	xor     int
	max     int
}

func (op *Optimus) Encode(n int) int {
	return ((n * op.prime) & op.max) ^ op.xor
}

func (op *Optimus) Decode(n int) int {

	return ((n ^ op.xor) * op.inverse) & op.max
}
```

## Usage
```go
opt := NewOptimus(1580030173, 59260789, 1163945558, 31)
for i := 0; i < 20; i++ {
	n := opt.Encode(i)
	k := opt.Decode(n)
	fmt.Println(k, "=>", n)
}
```


## Ref
[https://github.com/jenssegers/optimus](https://github.com/jenssegers/optimus)
