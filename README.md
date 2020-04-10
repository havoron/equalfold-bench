# equalfold-bench

 Benchmarks for strings.EqualFold.

## Hardware

- Intel Core i5-8250U CPU @ 1.60GHz, RAM 2400 MHz
- Intel Core i7-3770K CPU @ 3.50GHz, RAM 1800 MHz
- Intel Core i5-3317U CPU @ 1.70GHz, RAM SODIMM DDR3 Synchronous 1600 MHz (0,6 ns)

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
