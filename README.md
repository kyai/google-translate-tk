# Usage

```go
import gt "github.com/kyai/google-translate-tk"
```

Get `tkk`, this value usually does not change, so you can cache it.

```go
tkk, _ := gt.GetTKK()
```

Get `tk`, give the text to be translated and `tkk`.

```go
tk := gt.GetTK("hello", tkk)
```
