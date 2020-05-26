## combinations

Provides combinations of inputs.

Supports a maximum number of 63 items to choose from.

Usage:

```go
n := 2
r := []string{"a", "b", "c"}
f := func(combination []string) (stop bool) {
	fmt.Println(combination)
	return false
}
combinations.OfStrings(n, r, f)
```

Output:

```
[a b]
[a c]
[b c]
```
