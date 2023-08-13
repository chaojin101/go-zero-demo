# bookstore-cache demo

tech stack
- go-zero
- mysql
- redis

## Create api file

```shell
mkdir bookstore-cache
cd bookstore-cache
go mod init bookstore-cache
```

add a new file `bookstore-cache/bookstore.api`
```text
type (
	createReq {
		Name string `json:"name"`
	}
	createResp {
		Message string `json:"message"`
	}
)

type Book {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type (
	readBookByIdReq {
		Id int64 `path:"id"`
	}
	readBookByIdResp {
		Book
	}
)

type (
	readBookByNameReq {
		Name string `form:"name"`
	}
	readBookByNameResp {
		Book
	}
)

@server (
	prefix: /v1/books
)
service bookstore-api {
	@handler createBook
	post / (createReq) returns (createResp)
	@handler readBookById
	get /:id (readBookByIdReq) returns (readBookByIdResp)
	@handler readBookByName
	get / (readBookByNameReq) returns (readBookByNameResp)
}
```

## Create model file

the difference between `bookstore` and `bookstore-cache` is that `bookstore-cache` has `-cache` flag in `bookstore-cache/model/scripts/gen.sh`

add a new file `bookstore-cache/model/scripts/ddl.sql`
```sql
CREATE TABLE IF NOT EXISTS `book` (
    `id` integer PRIMARY KEY AUTO_INCREMENT,
    `name` varchar(255) NOT NULL
);
```

use this file `ddl.sql` to create table in mysql

add a new file `bookstore-cache/model/scripts/gen.sh`, set your own mysql config
```shell
#!/usr/bin/env bash

# https://github.com/Mikaelemmmm/go-zero-looklook/blob/main/deploy/script/mysql/genModel.sh

# set your own mysql config
host=172.30.112.1
port=3306
username=root
passwd=password
dbname=bookstore
#

table=$1
outDir=./model

echo "开始创建库：$dbname 的表：$table"
goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${table}"  -dir="${outDir}" --style=goZero -cache
```

## Generate go file

add a new file `bookstore-cache/Makefile`
```makefile
api:
	goctl api go -api *.api -dir . -style goZero
gen_model:
	sh ./model/scripts/gen.sh ${table}
```

```shell
make api
make gen_model table=book
go mod tidy
```

## Add configs

append db config to `bookstore/etc/bookstore-api.yaml`, using your own mysql and redis config
```yaml
# database
DB:
  DataSource: root:password@tcp(172.30.112.1:3306)/bookstore

# Cache
Cache:
  - Host: 127.0.0.1:6379
    Pass:
```

modify `Config struct` in `bookstore/internal/config/config.go` to
```go
type Config struct {
	rest.RestConf
	DB struct {
		DataSource string
	}
	Cache cache.CacheConf
}
```

modify `ServiceContext` struct and `NewServiceContext` function in `bookstore/internal/svc/serviceContext.go` to
```go
type ServiceContext struct {
	Config    config.Config
	BookModel model.BookModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		BookModel: model.NewBookModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
```

## Write logic

### CreateBook

modify `CreateBook` function in `bookstore/internal/logic/createBookLogic.go` to
```go
func (l *CreateBookLogic) CreateBook(req *types.CreateReq) (resp *types.CreateResp, err error) {
	_, err = l.svcCtx.BookModel.Insert(l.ctx, &model.Book{
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateResp{
		Message: "create " + req.Name,
	}, nil
}
```

### ReadBookById

modify `ReadBookById` function in `bookstore/internal/logic/readBookByIdLogic.go` to
```go
func (l *ReadBookByIdLogic) ReadBookById(req *types.ReadBookByIdReq) (resp *types.ReadBookByIdResp, err error) {
	book, err := l.svcCtx.BookModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &types.ReadBookByIdResp{
		Book: types.Book{
			Id:   book.Id,
			Name: book.Name,
		},
	}, nil
}
```

### ReadBookByName

modify `ReadBookByName` function in `bookstore/internal/logic/readBookByNameLogic.go` to
```go
func (l *ReadBookByNameLogic) ReadBookByName(req *types.ReadBookByNameReq) (resp *types.ReadBookByNameResp, err error) {
	book, err := l.svcCtx.BookModel.FindOneByName(l.ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &types.ReadBookByNameResp{
		Book: types.Book{
			Id:   book.Id,
			Name: book.Name,
		},
	}, nil
}
```

add a `FindOneByName` function in `bookstore/model/bookModel.go`
```go
func (m *defaultBookModel) FindOneByName(ctx context.Context, name string) (*Book, error) {
	query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", bookRows, m.table)
	var resp Book
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, name)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
```

add a method `FindOneByName` in `bookstore/model/bookModel.go` to `BookModel interface`
```go
type {
    BookModel interface {
        bookModel
        FindOneByName(ctx context.Context, name string) (*Book, error)
    }
    //...
}
```

## Run

```shell
go run bookstore.go
```