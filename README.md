## permutations

Go implementation of https://en.wikipedia.org/wiki/Heap%27s_algorithm

Usage:

```go
input := []string{"a", "b", "c"}
for v := range permutations.Strings(input) {
	fmt.Println(v)
}
```

Output:

```
["a","b","c"]
["b","a","c"]
["c","a","b"]
["a","c","b"]
["b","c","a"]
["c","b","a"]
```
