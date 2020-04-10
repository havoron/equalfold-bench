# equalfold-bench

 Benchmarks for strings.EqualFold.

## Hardware

- Intel Core i5-8250U CPU @ 1.60GHz, RAM 2400 MHz
- Intel Core i7-3770K CPU @ 3.50GHz, RAM 1800 MHz
- Intel Core i5-3317U CPU @ 1.70GHz, RAM SODIMM DDR3 Synchronous 1600 MHz (0,6 ns)
- Intel Celeron CPU N3350 @ 1.10GHz

## Test Cases

- [x] EqualFold ASCII.
- [x] EqualFold cyrillic.
- [x] == (equals operator)
- [x] == (equals operator) inside a function.
- [x] Call strings.EqualFold without if test.
- [x] Single ToLower.
- [x] Double ToLower.
- [x] ToLower ASCII.
- [x] ToLower cyrillic.

## TODO

- [ ] Add comparison table.
- [x] Read [How to write benchmarks in Go](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)
- [ ] Watch [justforfunc #28: Basic Benchmarks](https://www.youtube.com/watch?v=2AulMm-hsdI)
- [ ] Read [Analyzing the performance of Go functions with benchmarks](https://medium.com/justforfunc/analyzing-the-performance-of-go-functions-with-benchmarks-60b8162e61c6)
- [ ] Read [Practical Go Benchmarks](https://stackimpact.com/blog/practical-golang-benchmarks/)

## Questions

- When to use `strings.EqualFold` instead of double `strings.ToLower`?
- When `strings.ToLower` doesn't equal `strings.EqualFold`?
- How benchmark detects allocations?

# Benchmarks

## Intel Celeron CPU N3350 @ 1.10GHz

```
Running tool: /usr/local/go/bin/go test -benchmem -run=^$ github.com/havoron/equalfold-bench -bench .

goos: linux
goarch: amd64
pkg: github.com/havoron/equalfold-bench
BenchmarkConcatBuffer_WriteString-2       	    5424	    212070 ns/op	   37296 B/op	       9 allocs/op
BenchmarkConcatBuffer_Write-2             	    5654	    242331 ns/op	   37296 B/op	       9 allocs/op
BenchmarkConcatString-2                   	      46	  26755066 ns/op	53163994 B/op	    9999 allocs/op
BenchmarkConcatBuilder-2                  	   17005	     69165 ns/op	   48504 B/op	      17 allocs/op
BenchmarkConcatBuilder_Reset-2            	   16508	     68924 ns/op	   48504 B/op	      17 allocs/op
BenchmarkConcatBuilder_Single-2           	134815324	         9.39 ns/op	       4 B/op	       0 allocs/op
BenchmarkConcatString_Single-2            	  484178	      3991 ns/op	    5289 B/op	       0 allocs/op
BenchmarkConcatBuffer_Single-2            	56014182	        25.3 ns/op	       3 B/op	       0 allocs/op
BenchmarkEqualFold-2                      	14558738	        76.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkEqualFold_WithoutIf-2            	14264120	        81.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkEqualFold_Cyrillic-2             	 1536663	       765 ns/op	       0 B/op	       0 allocs/op
BenchmarkNotEqualFold-2                   	30824091	        42.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkEqual-2                          	1000000000	         0.707 ns/op	       0 B/op	       0 allocs/op
BenchmarkEqual_Var-2                      	1000000000	         0.732 ns/op	       0 B/op	       0 allocs/op
BenchmarkEqual_Cyrillic-2                 	1000000000	         0.744 ns/op	       0 B/op	       0 allocs/op
BenchmarkNotEqual-2                       	1000000000	         0.686 ns/op	       0 B/op	       0 allocs/op
BenchmarkEqual_InsideFunc-2               	122106164	         9.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkEqual_SingleToLower-2            	11044893	       129 ns/op	       8 B/op	       1 allocs/op
BenchmarkEqual_DoubleToLower-2            	 5028532	       246 ns/op	      16 B/op	       2 allocs/op
BenchmarkEqual_DoubleToLower_Cyrillic-2   	  718818	      1886 ns/op	      32 B/op	       2 allocs/op
BenchmarkFib1-2                           	249203096	         5.31 ns/op	       0 B/op	       0 allocs/op
BenchmarkFib2-2                           	79173658	        16.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkFib3-2                           	37090099	        32.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkFib10-2                          	  900574	      1310 ns/op	       0 B/op	       0 allocs/op
BenchmarkFib20-2                          	    5313	    205896 ns/op	       0 B/op	       0 allocs/op
BenchmarkFib40-2                          	       1	2818285289 ns/op	       0 B/op	       0 allocs/op
BenchmarkFibComplete-2                    	  988584	      1207 ns/op	       0 B/op	       0 allocs/op
BenchmarkToLowerASCII_AlreadyLower-2      	65721991	        20.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkToLowerASCII-2                   	16263948	        86.6 ns/op	       3 B/op	       1 allocs/op
BenchmarkToLowerCyrillic_AlreadyLower-2   	 3380022	       332 ns/op	       0 B/op	       0 allocs/op
BenchmarkToLowerCyrillic_Equal-2          	  730219	      1759 ns/op	      32 B/op	       2 allocs/op
PASS
ok  	github.com/havoron/equalfold-bench	48.723s
Success: Benchmarks passed.
```