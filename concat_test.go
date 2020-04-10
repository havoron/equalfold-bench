/*
	Practical Go Benchmarks
	https://stackimpact.com/blog/practical-golang-benchmarks/

	Hardware: Intel(R) Core(TM) i5-3317U CPU @ 1.70GHz

	100 one character concatenations

	BenchmarkConcatBuffer_WriteString-4   	  527738	      2283 ns/op	     208 B/op	       2 allocs/op
	BenchmarkConcatBuffer_Write-4         	  469114	      2248 ns/op	     208 B/op	       2 allocs/op
	BenchmarkConcatString-4               	   80935	     16450 ns/op	    5728 B/op	      99 allocs/op
	BenchmarkConcatBuilder-4              	 1300896	       968 ns/op	     248 B/op	       5 allocs/op
	BenchmarkConcatBuilder_Reset-4        	 1261710	       942 ns/op	     248 B/op	       5 allocs/op
	BenchmarkConcatBuilder_Single-4       	100000000	        10.7 ns/op	       2 B/op	       0 allocs/op
	BenchmarkConcatString_Single-4        	 7171830	       163 ns/op	      57 B/op	       0 allocs/op
	BenchmarkConcatBuffer_Single-4        	43374642	        25.3 ns/op	       2 B/op	       0 allocs/op

	10000 one character concatenations

	BenchmarkConcatBuffer_WriteString-4   	    5284	    232283 ns/op	   37296 B/op	       9 allocs/op
	BenchmarkConcatBuffer_Write-4         	    4981	    231886 ns/op	   37296 B/op	       9 allocs/op
	BenchmarkConcatString-4               	      40	  30619861 ns/op	53163998 B/op	    9999 allocs/op
	BenchmarkConcatBuilder-4              	   14174	     84585 ns/op	   48504 B/op	      17 allocs/op
	BenchmarkConcatBuilder_Reset-4        	   14468	     79649 ns/op	   48504 B/op	      17 allocs/op
	BenchmarkConcatBuilder_Single-4       	142615550	         8.00 ns/op	       4 B/op	       0 allocs/op
	BenchmarkConcatString_Single-4        	  475950	      2891 ns/op	    5289 B/op	       0 allocs/op
	BenchmarkConcatBuffer_Single-4        	44059114	        23.9 ns/op	       3 B/op	       0 allocs/op
*/

package equalfold

import (
	"bytes"
	"strings"
	"testing"
)

var strLen int = 10000

func BenchmarkConcatBuffer_WriteString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var buffer bytes.Buffer
		for i := 0; i < strLen; i++ {
			buffer.WriteString("x")
		}
	}
}

func BenchmarkConcatBuffer_Write(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var buffer bytes.Buffer
		for i := 0; i < strLen; i++ {
			buffer.Write([]byte("x"))
		}
	}
}

func BenchmarkConcatString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		str := ""
		for i := 0; i < strLen; i++ {
			str += "x"
		}
	}
}

func BenchmarkConcatBuilder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var builder strings.Builder
		for i := 0; i < strLen; i++ {
			builder.WriteString("x")
		}
	}
}

func BenchmarkConcatBuilder_Reset(b *testing.B) {
	var builder strings.Builder
	for n := 0; n < b.N; n++ {
		builder.Reset()
		for i := 0; i < strLen; i++ {
			builder.WriteString("x")
		}
	}
}

func BenchmarkConcatBuilder_Single(b *testing.B) {
	var builder strings.Builder

	i := 0

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		builder.WriteString("x")

		i++
		if i >= strLen {
			i = 0
			builder = strings.Builder{}
		}
	}
}

func BenchmarkConcatString_Single(b *testing.B) {
	var str string

	i := 0

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		str += "x"

		i++
		if i >= strLen {
			i = 0
			str = ""
		}
	}
}

func BenchmarkConcatBuffer_Single(b *testing.B) {
	var buffer bytes.Buffer

	i := 0

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		buffer.WriteString("x")

		i++
		if i >= strLen {
			i = 0
			buffer = bytes.Buffer{}
		}
	}
}
