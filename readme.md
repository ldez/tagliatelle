# Tagliatelle

[![Sponsor](https://img.shields.io/badge/Sponsor%20me-%E2%9D%A4%EF%B8%8F-pink)](https://github.com/sponsors/ldez)
[![Build Status](https://github.com/ldez/tagliatelle/workflows/Main/badge.svg?branch=master)](https://github.com/ldez/tagliatelle/actions)

A linter that handles struct tags.

Supported string casing:

- `camel`
- `pascal`
- `kebab`
- `snake`
- `goCamel` Respects Go's common initialisms (e.g. HttpResponse -> HTTPResponse).
- `goPascal` Respects Go's common initialisms (e.g. HttpResponse -> HTTPResponse).
- `goKebab` Respects Go's common initialisms (e.g. HttpResponse -> HTTPResponse).
- `goSnake` Respects Go's common initialisms (e.g. HttpResponse -> HTTPResponse).
- `upper`
- `lower`

## Examples

```go
// json and camel case
type Foo struct {
    ID     string `json:"ID"` // must be "id"
    UserID string `json:"UserID"`// must be "userId"
    Name   string `json:"name"`
    Value  string `json:"val,omitempty"`// must be "value"
}
```
