package invite

import (
	"errors"
	"strings"
)

const CHARSET = "97FEMpQdLjq2ca3yGU5ZrHB84bDznYkWeRSgKoXmJh6itCuNvATsPxwVf"

//var CHARSET = []rune{
//	'9', '7', 'F', 'E', 'M', 'p', 'Q', 'd', 'L', 'j', 'q', '2', 'c', 'a', '3', 'y', 'G', 'U', '5',
//	'Z', 'r', 'H', 'B', '8', '4', 'b', 'D', 'z', 'n', 'Y', 'k', 'W', 'e', 'R', 'S', 'g', 'K', 'o',
//	'X', 'm', 'J', 'h', '6', 'i', 't', 'C', 'u', 'N', 'v', 'A', 'T', 's', 'P', 'x', 'w', 'V', 'f',
//}

var base = uint64(len(CHARSET))

type Generator struct {
	length       int
	coprime      int
	decodeFactor uint16
	maxSupport   uint64
}

func NewGenerator(length uint8) (*Generator, error) {
	return &Generator{
		length:       int(length),
		coprime:      int(minCoprime(uint64(length))),
		decodeFactor: uint16(base) * uint16(length),
		maxSupport:   pow(base, uint64(length)) - 1,
	}, nil
}

func (g *Generator) MaxSupportID() uint64 {
	return g.maxSupport
}

// Encode 通过id获取指定邀请码（进制法+扩散+混淆）
func (g *Generator) Encode(id uint64) (string, error) {
	if id > g.maxSupport {
		return "", errors.New("id out of range")
	}

	idx := make([]uint16, g.length)

	// 扩散
	for i := 0; i < g.length; i++ {
		idx[i] = uint16(id % base)
		idx[i] = (idx[i] + uint16(i)*idx[0]) % uint16(base)
		id /= base
	}

	// 混淆
	var buf strings.Builder
	buf.Grow(g.length)
	for i := 0; i < g.length; i++ {
		n := i * g.coprime % g.length
		buf.WriteByte(CHARSET[idx[n]])
	}

	return buf.String(), nil
}

// Decode 通过邀请码反推id
func (g *Generator) Decode(code string) uint64 {
	var idx = make([]uint16, g.length)
	for i, c := range code {
		idx[i*g.coprime%g.length] = uint16(strings.IndexRune(CHARSET, c)) // 反推下标数组
	}

	var id uint64
	for i := g.length - 1; i >= 0; i-- {
		id *= base
		idx[i] = (idx[i] + g.decodeFactor - idx[0]*uint16(i)) % uint16(base)
		id += uint64(idx[i])
	}

	return id
}

// 求uint64类型n的最小互质数
func minCoprime(n uint64) uint64 {
	// 如果n是1，那么最小互质数是2
	if n == 1 {
		return 2
	}
	// 从2开始遍历，找到第一个和n互质的数
	for i := uint64(2); i < n; i++ {
		// 如果i和n的最大公约数是1，那么i就是最小互质数
		if isCoprime(i, n) {
			return i
		}
	}
	// 如果没有找到，那么返回n+1，因为n+1一定和n互质
	return n + 1
}

// 判断两个数是否互质
func isCoprime(n, m uint64) bool {
	// 求最大公因数
	return gcd(n, m) == 1
}

// 辗转相除法求最大公因数
func gcd(n, m uint64) uint64 {
	if m == 0 {
		return n
	}
	return gcd(m, n%m)
}

// 求n的m次方
func pow(n, m uint64) uint64 {
	sum := n
	for i := uint64(1); i < m; i++ {
		sum *= n
	}
	return sum
}
