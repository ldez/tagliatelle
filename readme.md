# Tagliatelle

[![Sponsor](https://img.shields.io/badge/Sponsor%20me-%E2%9D%A4%EF%B8%8F-pink)](https://github.com/sponsors/ldez)
[![Build Status](https://github.com/ldez/tagliatelle/workflows/Main/badge.svg?branch=master)](https://github.com/ldez/tagliatelle/actions)

A linter that handles struct tags.

Supported string casing:

- `camel`
- `pascal`
- `kebab`
- `snake`
- `goCamel` Respects [Go's common initialisms](https://github.com/golang/lint/blob/83fdc39ff7b56453e3793356bcff3070b9b96445/lint.go#L770-L809) (e.g. HttpResponse -> HTTPResponse).
- `goPascal` Respects [Go's common initialisms](https://github.com/golang/lint/blob/83fdc39ff7b56453e3793356bcff3070b9b96445/lint.go#L770-L809) (e.g. HttpResponse -> HTTPResponse).
- `goKebab` Respects [Go's common initialisms](https://github.com/golang/lint/blob/83fdc39ff7b56453e3793356bcff3070b9b96445/lint.go#L770-L809) (e.g. HttpResponse -> HTTPResponse).
- `goSnake` Respects [Go's common initialisms](https://github.com/golang/lint/blob/83fdc39ff7b56453e3793356bcff3070b9b96445/lint.go#L770-L809) (e.g. HttpResponse -> HTTPResponse).
- `upper`
- `lower`

| Source         | Camel Case     | Go Camel Case  |
|----------------|----------------|----------------|
| GooID          | gooId          | gooID          |
| HTTPStatusCode | httpStatusCode | httpStatusCode |
| FooBAR         | fooBar         | fooBar         |
| URL            | url            | url            |
| ID             | id             | id             |
| hostIP         | hostIp         | hostIP         |
| JSON           | json           | json           |
| JSONName       | jsonName       | jsonName       |
| NameJSON       | nameJson       | nameJSON       |
| UneTête        | uneTête        | uneTête        |

| Source         | Pascal Case    | Go Pascal Case |
|----------------|----------------|----------------|
| GooID          | GooId          | GooID          |
| HTTPStatusCode | HttpStatusCode | HTTPStatusCode |
| FooBAR         | FooBar         | FooBar         |
| URL            | Url            | URL            |
| ID             | Id             | ID             |
| hostIP         | HostIp         | HostIP         |
| JSON           | Json           | JSON           |
| JSONName       | JsonName       | JSONName       |
| NameJSON       | NameJson       | NameJSON       |
| UneTête        | UneTête        | UneTête        |

| Source         | Snake Case       | Go Snake Case    |
|----------------|------------------|------------------|
| GooID          | goo_id           | goo_ID           |
| HTTPStatusCode | http_status_code | HTTP_status_code |
| FooBAR         | foo_bar          | foo_bar          |
| URL            | url              | URL              |
| ID             | id               | ID               |
| hostIP         | host_ip          | host_IP          |
| JSON           | json             | JSON             |
| JSONName       | json_name        | JSON_name        |
| NameJSON       | name_json        | name_JSON        |
| UneTête        | une_tête         | une_tête         |

| Source         | Kebab Case       | Go KebabCase     |
|----------------|------------------|------------------|
| GooID          | goo-id           | goo-ID           |
| HTTPStatusCode | http-status-code | HTTP-status-code |
| FooBAR         | foo-bar          | foo-bar          |
| URL            | url              | URL              |
| ID             | id               | ID               |
| hostIP         | host-ip          | host-IP          |
| JSON           | json             | JSON             |
| JSONName       | json-name        | JSON-name        |
| NameJSON       | name-json        | name-JSON        |
| UneTête        | une-tête         | une-tête         |

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

## What this tool is about

This tool is about validating tags according to rules you define.

The tool also allows to fix tags according to the rules you defined.

## What this tool is not

This tool is not intended to validate the fact a tag in valid or not.

To do that, you can use `go vet`, or use golangci "go vet" linter.

More information about it here: <https://golangci-lint.run/usage/linters/#govet>

## How to use the tool

### Install and run it from the binary

```shell
go install github.com/ldez/tagliatelle/cmd/tagliatelle@latest
```

then launch it manually:

```text
tagliatelle: Checks the struct tags.

Usage: tagliatelle [-flag] [package]

Flags:
  -V print version and exit
  -c int
     display offending line with this many lines of context (default -1)
  -cpuprofile string
     write CPU profile to this file
  -debug string
     debug flags, any subset of "fpstv"
  -fix
     apply all suggested fixes
  -flags
     print analyzer flags in JSON
  -json
     emit JSON output
  -memprofile string
     write memory profile to this file
  -trace string
     write trace log to this file
```

Deprecated flags:

The following flags were available in previous versions of the tools, they are now deprecated

```text
Flags:
  -all
     no effect (deprecated)
  -source
     no effect (deprecated)
  -tags string
     no effect (deprecated)
  -v no effect (deprecated)
```

### As a golangci linter

Define the rules, you want via your golangci configuration file

```yaml
linters-settings:
  tagliatelle:
    # Check the struck tag name case.
    case:
      # Use the struct field name to check the name of the struct tag.
      # Default: false
      use-field-name: true
      rules:
        # Any struct tag type can be used.
        # Support string case: `camel`, `pascal`, `kebab`, `snake`, `goCamel`, `goPascal`, `goKebab`, `goSnake`, `upper`, `lower`
        json: camel
        yaml: camel
        xml: camel
```

More information here <https://golangci-lint.run/usage/linters/#tagliatelle>

Then either enable tagliatelle by passing it on demand.

```shell
golangci run --enable tagliatelle
```

or enable it directly via golangci configuration file, by adding this

```yaml
linters:
    enable:
      - tagliatelle
```

Please note, these examples are provided in yaml format.

GolangCI also supports other formats, please refer to [golangci.run](https://golangci-lint.run/)
for a more complete documentation.

## Rules

Here are the default rules for the well known and used tags, when using tagliatelle as a binary or golangci

- json: "camel"
- yaml: "camel"
- xml:  "camel"
- bson: "camel"
- avro: "snake"

### Custom Rules

The tool is not limited to the tags used in example, you can use it to validate any tag.

You can add your own tag, for example "whatever" and tells the tool you want to use "kebab"

This option is only available via golangci linter.

```yaml
linters-settings:
  tagliatelle:
    # Check the struck tag name case.
    case:
      # Use the struct field name to check the name of the struct tag.
      # Default: false
      use-field-name: true
      rules:
        # Any struct tag type can be used.
        # Support string case: `camel`, `pascal`, `kebab`, `snake`, `goCamel`, `goPascal`, `goKebab`, `goSnake`, `upper`, `lower`
        json:     camel
        yaml:     camel
        xml:      camel
        whatever: kebab
```

Source

```go
// json and camel case
type Foo struct {
    ID     string `json:"ID"` // must be "id"
    UserID string `json:"UserID" whatever:"userid"`// json must be "userId", but whatever must be "user-id"
    Name   string `json:"name"`
    Value  string `json:"val,omitempty"`// must be "value"
}
```

Expected

```go
type Foo struct {
    ID     string `json:"id"`
    UserID string `json:"userId" whatever:"user-id"`
    Name   string `json:"name"`
    Value  string `json:"value,omitempty"`
}
```

## Maintainers

- Ludovic Fernandez [ldez](https://github.com/ldez)
