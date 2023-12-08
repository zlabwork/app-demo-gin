## code

```golang
package libs

import (
	"bytes"
	"fmt"
	"math/big"
)

const base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

func Base58Encode(input []byte) string {
	var result []byte

	// 将字节数组转换为大整数
	bigInt := new(big.Int).SetBytes(input)

	// 获取Base58的长度
	base := big.NewInt(58)
	zero := big.NewInt(0)
	mod := &big.Int{}

	for bigInt.Cmp(zero) > 0 {
		// 通过取模运算获得Base58字符
		bigInt.DivMod(bigInt, base, mod)
		result = append(result, base58Alphabet[mod.Int64()])
	}

	// 反转结果
	reverseBytes(result)

	// 处理前导0
	for _, b := range input {
		if b == 0x00 {
			result = append([]byte{base58Alphabet[0]}, result...)
		} else {
			break
		}
	}

	return string(result)
}

func Base58Decode(input string) ([]byte, error) {
	x := big.NewInt(0)
	base := big.NewInt(58)
	base58AlphabetBytes := []byte(base58Alphabet)

	// 计算输入字符串对应的大数值
	for _, c := range []byte(input) {
		pos := bytes.IndexByte(base58AlphabetBytes, c)
		if pos == -1 {
			return nil, fmt.Errorf("invalid input character: %c", c)
		}
		x.Mul(x, base)
		x.Add(x, big.NewInt(int64(pos)))
	}

	// 将大数值转换为字节数组
	result := x.Bytes()

	// 处理前导零
	leadingZeros := 0
	for leadingZeros < len(input) && input[leadingZeros] == base58Alphabet[0] {
		leadingZeros++
	}
	result = append(bytes.Repeat([]byte{0}, leadingZeros), result...)

	return result, nil
}

// 反转字节数组
func reverseBytes(bytes []byte) {
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
}
```
