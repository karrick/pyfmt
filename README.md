# pyfmt

python style formatter

# Installation

```bash
go install gitli.corp.linkedin.com/secsre/pyfmt
```

# Usage

```Go
dict := make(map[string]string)
dict["name"] = "John"
dict["age"] = "99
foo := pyfmt.Sprintf("Hello {name}, are you really {age} years old?", dict)
```

# Limitations

 * does not allow embedded curly braces
 * does not allow escaped { or }
