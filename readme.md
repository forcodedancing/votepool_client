Run following commands before run tests or main.

```shell
export CGO_CFLAGS="-O -D__BLST_PORTABLE__"
export CGO_CFLAGS_ALLOW="-O -D__BLST_PORTABLE__"
```

```shell
go test tm_test.go  -v
```

```shell
go run main.go
```