# pyfmt

python style formatter for the go programming language

Why would you need such a thing? You probably don't. But I did once,
so I created this in case anyone else one day needed it.

# Installation

```bash
go install github.com/karrick/pyfmt
```

# Usage

## `Sprintf` accepts a `map[string]string`:

```Go
dict := make(map[string]string)
dict["name"] = "John"
dict["age"] = "99"
actual := pyfmt.Sprintf("Hello {name}, are you really {age} years old?", dict)
```

## `Sprintf` accepts multiple arguments:

```Go
actual := Sprintf("-->{0}*{1}<--", 3.5, "silly", "ignored")
```
