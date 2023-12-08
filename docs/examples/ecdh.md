## Elliptic Curve Diffie-Hellman
```go
priKey1, _ := ecdh.P256().GenerateKey(rand.Reader)
pubKey1 := priKey1.PublicKey()

priKey2, _ := ecdh.P256().GenerateKey(rand.Reader)
pubKey2 := priKey2.PublicKey()

share1, _ := priKey1.ECDH(pubKey2)
share2, _ := priKey2.ECDH(pubKey1)
```

output
```text
2023/07/14 11:22:15 [254 23 146 85 59 58 87 253 119 16 130 248 134 143 228 236 101 11 35 166 33 220 9 237 60 87 50 234 22 220 120 22]
2023/07/14 11:22:15 [254 23 146 85 59 58 87 253 119 16 130 248 134 143 228 236 101 11 35 166 33 220 9 237 60 87 50 234 22 220 120 22]
```

## docs 
[https://pkg.go.dev/crypto/ecdh](https://pkg.go.dev/crypto/ecdh)
