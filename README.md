# Usage

```
import gt "github.com/kyai/google-translate-tk"
```

Get `tkk`, This value usually does not change, so you can cache it.

```
tkk, _ := gt.GetTKK()
```

Get `tk`, Give the text to be translated and `tkk`.

```
tk := gt.GetTK("hello", tkk)
```
