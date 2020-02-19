# go js
JavaScript virtual machine inside Golang (using V8)

Example of usage:

```go
js, err = gojs.New(1) // use 1 thread
if err != nil {
    return nil, err
}

defer js.Dispose()

err := js.Compile("my.js", "2 + 2")
if err != nil {
    return nil, err
}

future, err := js.Run("my.js")
if err != nil {
    return nil, err
}

res := <-future
if res.Err != nil {
    return nil, res.Err
}

defer res.Res.Dispose()

val, err := res.ToInt()
if err != nil {
    return nil, err
}

println(val) // int64, 4
```
