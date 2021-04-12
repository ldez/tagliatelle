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

| Source         | Snake Case       | Kebab Case       | Pascal Case    | Camel Case     | Go Pascal Case | Go Camel Case  | Go Snake Case    | Go KebabCase     |
|----------------|------------------|------------------|----------------|----------------|----------------|----------------|------------------|------------------|
| GooID          | goo_id           | goo-id           | GooId          | gooId          | gooID          | GooID          | goo_ID           | goo-ID           |
| HTTPStatusCode | http_status_code | http-status-code | HttpStatusCode | httpStatusCode | httpStatusCode | HTTPStatusCode | HTTP_status_code | HTTP-status-code |
| FooBAR         | foo_bar          | foo-bar          | FooBar         | fooBar         | fooBar         | FooBar         | foo_bar          | foo-bar          |
| URL            | url              | url              | Url            | url            | url            | URL            | URL              | URL              |
| ID             | id               | id               | Id             | id             | id             | ID             | ID               | ID               |
| hostIP         | host_ip          | host-ip          | HostIp         | hostIp         | hostIP         | HostIP         | host_IP          | host-IP          |
| JSON           | json             | json             | Json           | json           | json           | JSON           | JSON             | JSON             |
| JSONName       | json_name        | json-name        | JsonName       | jsonName       | jsonName       | JSONName       | JSON_name        | JSON-name        |
| NameJSON       | name_json        | name-json        | NameJson       | nameJson       | nameJSON       | NameJSON       | name_JSON        | name-JSON        |
| UneTête        | une_tête         | une-tête         | UneTête        | uneTête        | uneTête        | UneTête        | une_tête         | une-tête         |

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
