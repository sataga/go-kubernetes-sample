# go-kubernetes-sample

## mockを作り直したい

```bash
go generate ./...
goimports -w .
```

## テスト流したい

```bash
go test -v ./...

or 

go test -v github.com/sataga/go-kubernetes-sample/domain/updateconfigmap
```
