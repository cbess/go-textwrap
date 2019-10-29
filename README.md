# Go Text Wrapping

It is a lightweight and simple port of the [textwrap](https://docs.python.org/3.6/library/textwrap.html) Python module.

[Soli Deo gloria](https://perfectGod.com)

Features:

- Lightweight. Uses only the Go built-in library (no external deps).
- Simple. Easy to understand code.
- Counts number of words and characters.
- Splits (or groups) text based on character limit (width) per group (each wrap).
- `100%` test coverage.
- [GoDoc](https://godoc.org/github.com/cbess/go-textwrap) documentation

### Note

It was only designed to work with English text; it separates words by whitespace.

## Example

```go
import (
    "fmt"
    "textwrap"
)

origText := "Jesus is God. He Saves by grace through faith alone."
result, err := textwrap.WordWrap(origText, 10, -1)
if err != nil {
    // handle error
}

// check if max word count exceeded
if result.IsValid() {
    // print text groups
    for idx, text := range result.TextGroups {
        fmt.Println("[", idx+1, "]", text)
    }
}
```

## License

MIT
