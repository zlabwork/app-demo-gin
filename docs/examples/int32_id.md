## 混淆 uint32 id

```go
func main() {

    for i := 0; i < 10; i++ {
        fmt.Printf("%d\n", NewDistributedId(i))
    }

}

// n is 32 bits, uint32
func NewDistributedId(n int) int {
    s0 := n & 255 << 24
    s1 := n & (255 << 8) << 8
    s2 := n & (255 << 16) >> 8
    s3 := n & (255 << 24) >> 24
    return s0 | s1 | s2 | s3
}
```

## 生成唯一ID
```go
// m 与 p 是互质数
type SeqId struct {
	p int
	m int
	r int
}

func NewSeqId(prime, m, random int) *SeqId {
	return &SeqId{
		p: prime,
		m: m,
		r: random,
	}
}

func (sid *SeqId) Encode(n int) int {
	return (n*sid.p + sid.r) % sid.m
}
```

```go
seq := NewSeqId(1580030173, 9999, 123456)
for i := 0; i < 16; i++ {
	v := seq.Encode(i)
	fmt.Println(v)
}
```
