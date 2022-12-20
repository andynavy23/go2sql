# go2sql
Generate sql query string in golang

## Usage

### Initialize your module
```go
go mod init go2sql-demo
```

### Get the go2sql module
```go
go get github.com/andynavy23/go2sql@v0.1.0
```

#### main.go
```go
package main

import (
    "fmt"

    "github.com/andynavy23/go2sql"
)

func main() {
    sqlQuery := NewSqlQuery()
    sqlQuery.Select([]string{
		`col_a`,
		`col_b`,
	})
    sqlQuery.From(`table`)
    fmt.Println(sqlQuery.ToSql())
}
```

#### Console output
```console
SELECT
  col_a,
  col_b
FROM
  table
```

## Testing
```go
go test
```

## License
The MIT License (MIT)

## Donate
ERC20 wallet address:  
0x390634EfA0f3a4B083572a19B9D17063027feFb6
  
[Paypal Me](https://www.paypal.me/fu0224)

