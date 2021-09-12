# builderpool

A simple strings.Builder pool using sync.

## SYNOPSIS

``` go
import "github.com/webx-top/poolx/builderpool"

var pool = builderpool.New()

func main() {
    builder := pool.Get()
    defer pool.Release(builder)

    // ...
}
```
