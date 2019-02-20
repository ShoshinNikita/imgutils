# imgutils

**imgutils** is a Go package with useful functions for image processing

## Example

```go
  package main

  import (
      "image"
      "image/jpeg"
      "os"

      "github.com/ShoshinNikita/imgutils"
  )

  func main() {
      var img1, img2, img3 image.Image

      // Load images
      // ...

      res := imgutils.Concatenate(img1, img2, imgutils.ConcatHorizontalMode)
      res = imgutils.Concatenate(res, img3, imgutils.ConcatVerticalMode)

      f, _ = os.Create("result.jpeg")
      jpeg.Encode(f, res, nil)
  }
```

**Result: (resized)**

<img src="./example/result.jpeg" alt="drawing" width="300"/>


[Full program](example/example.go)

[Full image](./example/result.jpeg)