package equalfold

import (
	"strings"
	"testing"
)

/*
	BenchmarkEqualFold-8   	60042930	        20.5 ns/op	       0 B/op	       0 allocs/op
	BenchmarkEqualFold-8   	52315848	        20.7 ns/op	       0 B/op	       0 allocs/op
	BenchmarkEqualFold-8   	60037822	        20.3 ns/op	       0 B/op	       0 allocs/op
*/
func BenchmarkEqualFold(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if r := strings.EqualFold("abcDEF", "ABCdef"); !r {
			b.Errorf("got %v but want %v", r, true)
		}
	}
}

/*
	BenchmarkEqualFold_WithoutIf-8   	59944850	        20.5 ns/op	       0 B/op	       0 allocs/op
	BenchmarkEqualFold_WithoutIf-8   	59950540	        20.1 ns/op	       0 B/op	       0 allocs/op
	BenchmarkEqualFold_WithoutIf-8   	57111309	        20.5 ns/op	       0 B/op	       0 allocs/op
*/
func BenchmarkEqualFold_WithoutIf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.EqualFold("abcDEF", "ABCdef")
	}
}

/*
	BenchmarkEqualFold_Cyrillic-8   	 5179902	       226 ns/op	       0 B/op	       0 allocs/op
	BenchmarkEqualFold_Cyrillic-8   	 5225479	       221 ns/op	       0 B/op	       0 allocs/op
	BenchmarkEqualFold_Cyrillic-8   	 5380099	       221 ns/op	       0 B/op	       0 allocs/op
*/
func BenchmarkEqualFold_Cyrillic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if r := strings.EqualFold("абвГДЕ", "АБВгде"); !r {
			b.Errorf("got %v but want %v", r, true)
		}
	}
}

/*
	BenchmarkNotEqualFold-8   	85716734	        13.5 ns/op	       0 B/op	       0 allocs/op
	BenchmarkNotEqualFold-8   	85585906	        13.0 ns/op	       0 B/op	       0 allocs/op
	BenchmarkNotEqualFold-8   	85712448	        13.4 ns/op	       0 B/op	       0 allocs/op
*/
func BenchmarkNotEqualFold(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if r := strings.EqualFold("abcXYZ", "ABCdef"); r {
			b.Errorf("got %v but want %v", r, false)
		}
	}
}

/*
	BenchmarkEqual-8   	1000000000	         0.301 ns/op	       0 B/op	       0 allocs/op
	BenchmarkEqual-8   	1000000000	         0.304 ns/op	       0 B/op	       0 allocs/op
	BenchmarkEqual-8   	1000000000	         0.303 ns/op	       0 B/op	       0 allocs/op
*/
func BenchmarkEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if "abc_def" != "abc_def" {
			b.Errorf("got %v but want %v", false, true)
		}
	}
}

/*
	BenchmarkEqual_Var-8   	1000000000	         0.302 ns/op	       0 B/op	       0 allocs/op
	BenchmarkEqual_Var-8   	1000000000	         0.313 ns/op	       0 B/op	       0 allocs/op
	BenchmarkEqual_Var-8   	1000000000	         0.317 ns/op	       0 B/op	       0 allocs/op
*/
func BenchmarkEqual_Var(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if r := "abc_def" != "abc_def"; r {
			b.Errorf("got %v but want %v", false, true)
		}
	}
}

/*
	BenchmarkEqual_Cyrillic-8   	1000000000	         0.303 ns/op	       0 B/op	       0 allocs/op
	BenchmarkEqual_Cyrillic-8   	1000000000	         0.307 ns/op	       0 B/op	       0 allocs/op
	BenchmarkEqual_Cyrillic-8   	1000000000	         0.318 ns/op	       0 B/op	       0 allocs/op
*/
func BenchmarkEqual_Cyrillic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if "абвгде" != "абвгде" {
			b.Errorf("got %v but want %v", true, false)
		}
	}
}

/*
	BenchmarkNotEqual-8   	1000000000	         0.605 ns/op	       0 B/op	       0 allocs/op
	BenchmarkNotEqual-8   	1000000000	         0.326 ns/op	       0 B/op	       0 allocs/op
	BenchmarkNotEqual-8   	1000000000	         0.302 ns/op	       0 B/op	       0 allocs/op
*/
func BenchmarkNotEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if "abcxyz" == "abcXYZ" {
			b.Errorf("got %v but want %v", true, false)
		}
	}
}

/*
	BenchmarkEqual_InsideFunc-8   	332268780	         3.37 ns/op	       0 B/op	       0 allocs/op
	BenchmarkEqual_InsideFunc-8   	360622434	         3.32 ns/op	       0 B/op	       0 allocs/op
	BenchmarkEqual_InsideFunc-8   	361701588	         3.34 ns/op	       0 B/op	       0 allocs/op
*/
func BenchmarkEqual_InsideFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if equal("abcxyz", "abcXYZ") {
			b.Errorf("got %v but want %v", true, false)
		}
	}
}

func equal(a, b string) bool {
	return a == b
}

/*
	BenchmarkEqual_SingleToLower-8   	29968906	        40.3 ns/op	       8 B/op	       1 allocs/op
	BenchmarkEqual_SingleToLower-8   	30660660	        40.1 ns/op	       8 B/op	       1 allocs/op
	BenchmarkEqual_SingleToLower-8   	29317417	        40.9 ns/op	       8 B/op	       1 allocs/op
*/
func BenchmarkEqual_SingleToLower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if "abcxyz" != strings.ToLower("abcXYZ") {
			b.Errorf("got %v but want %v", true, false)
		}
	}
}

/*
	BenchmarkEqual_DoubleToLower-8   	21396324	        53.5 ns/op	       8 B/op	       1 allocs/op
	BenchmarkEqual_DoubleToLower-8   	23138907	        52.4 ns/op	       8 B/op	       1 allocs/op
	BenchmarkEqual_DoubleToLower-8   	22685638	        54.2 ns/op	       8 B/op	       1 allocs/op
*/
func BenchmarkEqual_DoubleToLower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if strings.ToLower("abcxyz") != strings.ToLower("abcXYZ") {
			b.Errorf("got %v but want %v", true, false)
		}
	}
}
