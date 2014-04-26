# pyfmt

python style formatter

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
