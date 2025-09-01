# GoToSocial SDK

> [!WARNING]
> This package is deprecated and will no longer receive updates. Please migrate to [codeberg.org/gotosocial-sdk/golang](https://codeberg.org/gotosocial-sdk/golang) or other alternatives.

[![Go Reference](https://pkg.go.dev/badge/github.com/owu-one/gotosocial-sdk.svg)](https://pkg.go.dev/github.com/owu-one/gotosocial-sdk)

GoToSocial Swagger API Client (Auto-Generated)

## Generate

1. Install [Go Swagger](https://goswagger.io/install.html)
```bash
go get -u github.com/go-swagger/go-swagger/cmd/swagger
```

2. Generate
```bash
rm -rf client models
go generate ./...
```

3. Apply Patch
workaround for https://github.com/go-swagger/go-swagger/issues/2997.
```bash
patch -u -p1 -i patches/filter-context.diff
```

## Credits

- [`VyrCossont/slurp`](https://github.com/VyrCossont/slurp). The original project that this repo is extracted from.
- [`superseriousbusiness/gotosocial`](https://github.com/superseriousbusiness/gotosocial).
