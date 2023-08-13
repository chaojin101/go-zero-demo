# ping-pong demo

tech stack: 
- go-zero

## Create api file

```shell
mkdir ping-pong
cd ping-pong
go mod init ping-pong
```

add a new file `ping-pong.api`

`ping-pong.api`:
```text
type (
	req  struct{}
	resp {
		Message string `json:"message"`
	}
)

service ping-pong-api {
	@handler ping
	get /ping (req) returns (resp)
}
```

## Generate go file

```shell
goctl api go --api ./*.api --dir . --style goZero
go mod tidy
```

## Write ping logic

in `ping-pong-http/internal/logic/pingLogic.go`, add logic in `Ping` function

```go
func (l *PingLogic) Ping(req *types.Req) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line

	return &types.Resp{
		Message: "pong",
	}, nil
}
```

## Run

```shell
go run ping-pong.go
```