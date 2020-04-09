/*
	Hardware:

	CPU: Intel(R) Core(TM) i7-3770K CPU @ 3.50GHz
	RAM: 1600 MHz dual-channel
*/

package equalfold

import (
	"strings"
	"testing"
)

/*
	BenchmarkToLowerASCII_AlreadyLower-8    100000000               10.3 ns/op
*/
func BenchmarkToLowerASCII_AlreadyLower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.ToLower("abc")
	}
}

/*
	BenchmarkToLowerASCII-8         30790152                36.1 ns/op
*/
func BenchmarkToLowerASCII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.ToLower("ABC")
	}
}

/*
	BenchmarkToLowerCyrillic_AlreadyLower-8         12129788                95.7 ns/
*/
func BenchmarkToLowerCyrillic_AlreadyLower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.ToLower("абв")
	}
}

/*
	BenchmarkToLowerCyrillic_Equal-8   	 2287318	       497 ns/op	      32 B/op	       2 allocs/op
	BenchmarkToLowerCyrillic_Equal-8   	 2203382	       521 ns/op	      32 B/op	       2 allocs/op
	BenchmarkToLowerCyrillic_Equal-8   	 2140543	       533 ns/op	      32 B/op	       2 allocs/op
*/
func BenchmarkToLowerCyrillic_Equal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.ToLower("абвГДЕ") == strings.ToLower("АБВгде")
	}
}
