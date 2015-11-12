**ex. install `ffmpeg` first, then create a main.go**
```go
package main

import (
	"flag"
	"fmt"

	"github.com/goodjob1114/vmp3"
)

func main() {
	flag.Parse()
	root := flag.Arg(0)
	fmt.Println("rootDir::", root)
	vmp3.GoMp3(root)
	fmt.Println("Job::finished")
}

```
**run**

$ go run main.go /foo/bar/somefolder

it will convert video to mp3 all in folder
