# equalfold-bench

 Benchmarks for strings.EqualFold.

## Hardware

Laptop with Intel Core i5-8250U and dual-channel RAM with 2400 MHz.

## Test Cases

- [x] EqualFold ASCII.
- [x] EqualFold cyrillic.
- [x] == (equals operator)
- [x] == (equals operator) inside a function.
- [x] Call strings.EqualFold without if test.
- [x] Single ToLower.
- [x] Double ToLower.
- [ ] ToLower ASCII.
- [ ] ToLower cyrillic.

## TODO

- [ ] Add comparison table.

## Questions

- When to use `strings.EqualFold` instead of double `strings.ToLower`?
- When `strings.ToLower` doesn't equal `strings.EqualFold`?
- How benchmark detects allocations?
