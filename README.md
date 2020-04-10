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

## Intel Core i7-3770K CPU @ 3.50GHz, RAM 1800 MHz

```
Running tool: C:\Go\bin\go.exe test -benchmem -run=^$ github.com/havoron/equalfold-bench -bench . -v -count=1

goos: windows
goarch: amd64
pkg: github.com/havoron/equalfold-bench
BenchmarkConcatBuffer_WriteString
BenchmarkConcatBuffer_WriteString-8       	   12456	    101821 ns/op	   37296 B/op	       9 allocs/op
BenchmarkConcatBuffer_Write
BenchmarkConcatBuffer_Write-8             	   12534	     99151 ns/op	   37296 B/op	       9 allocs/op
BenchmarkConcatString
BenchmarkConcatString-8                   	     121	  10009482 ns/op	53164062 B/op	   10000 allocs/op
BenchmarkConcatBuilder
BenchmarkConcatBuilder-8                  	   40568	     30027 ns/op	   48504 B/op	      17 allocs/op
BenchmarkConcatBuilder_Reset
BenchmarkConcatBuilder_Reset-8            	   39895	     31385 ns/op	   48504 B/op	      17 allocs/op
BenchmarkConcatBuilder_Single
BenchmarkConcatBuilder_Single-8           	342121093	         3.27 ns/op	       4 B/op	       0 allocs/op
BenchmarkConcatString_Single
BenchmarkConcatString_Single-8            	 1136538	      1343 ns/op	    5305 B/op	       1 allocs/op
BenchmarkConcatBuffer_Single
BenchmarkConcatBuffer_Single-8            	100000000	        11.6 ns/op	       3 B/op	       0 allocs/op
BenchmarkEqualFold
BenchmarkEqualFold-8                      	36389880	        37.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkEqualFold_WithoutIf
BenchmarkEqualFold_WithoutIf-8            	42885783	        41.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkEqualFold_Cyrillic
BenchmarkEqualFold_Cyrillic-8             	 4921489	       274 ns/op	       0 B/op	       0 allocs/op
BenchmarkNotEqualFold
BenchmarkNotEqualFold-8                   	63216786	        16.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkEqual
BenchmarkEqual-8                          	1000000000	         0.287 ns/op	       0 B/op	       0 allocs/op
BenchmarkEqual_Var
BenchmarkEqual_Var-8                      	1000000000	         0.287 ns/op	       0 B/op	       0 allocs/op
BenchmarkEqual_Cyrillic
BenchmarkEqual_Cyrillic-8                 	1000000000	         0.291 ns/op	       0 B/op	       0 allocs/op
BenchmarkNotEqual
BenchmarkNotEqual-8                       	1000000000	         0.284 ns/op	       0 B/op	       0 allocs/op
BenchmarkEqual_InsideFunc
BenchmarkEqual_InsideFunc-8               	234082852	         5.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkEqual_SingleToLower
BenchmarkEqual_SingleToLower-8            	25018293	        48.5 ns/op	       8 B/op	       1 allocs/op
BenchmarkEqual_DoubleToLower
BenchmarkEqual_DoubleToLower-8            	11889568	        97.8 ns/op	      16 B/op	       2 allocs/op
BenchmarkEqual_DoubleToLower_Cyrillic
BenchmarkEqual_DoubleToLower_Cyrillic-8   	 2358512	       516 ns/op	      32 B/op	       2 allocs/op
BenchmarkFib1
BenchmarkFib1-8                           	482269753	         2.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkFib2
BenchmarkFib2-8                           	187545450	         6.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkFib3
BenchmarkFib3-8                           	100000000	        10.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkFib10
BenchmarkFib10-8                          	 2994614	       387 ns/op	       0 B/op	       0 allocs/op
BenchmarkFib20
BenchmarkFib20-8                          	   24861	     47956 ns/op	       0 B/op	       0 allocs/op
BenchmarkFib40
BenchmarkFib40-8                          	       2	 749054900 ns/op	       0 B/op	       0 allocs/op
BenchmarkFibComplete
BenchmarkFibComplete-8                    	 3002101	       397 ns/op	       0 B/op	       0 allocs/op
BenchmarkToLowerASCII_AlreadyLower
BenchmarkToLowerASCII_AlreadyLower-8      	121666033	        10.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkToLowerASCII
BenchmarkToLowerASCII-8                   	35318943	        35.0 ns/op	       3 B/op	       1 allocs/op
BenchmarkToLowerCyrillic_AlreadyLower
BenchmarkToLowerCyrillic_AlreadyLower-8   	12379773	        96.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkToLowerCyrillic_Equal
BenchmarkToLowerCyrillic_Equal-8          	 2415140	       509 ns/op	      32 B/op	       2 allocs/op
PASS
ok  	github.com/havoron/equalfold-bench	45.921s
Success: Benchmarks passed.
```