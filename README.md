
# go-configurable

Make things configurable.

### Usage

Struct embedding

```go
package main

import "fmt"
import "configurable"

type App struct {
  configurable.Settings
}

func main() {
  app := &App{}

  // defaults
  app.Defaults(map[string]interface{}{
    "debug": true,
    "title": "api",
    "port": 3000,
  })

  // get an option
  fmt.Printf("debug=%t\n", app.Get("debug"))
  fmt.Printf("port=%d\n", app.Get("port"))
  fmt.Printf("title=%s\n", app.Get("title"))

  // set an option
  app.Set("port", 3001)
  fmt.Printf("port=%d\n", app.Get("port"))

  // toogle an option
  app.Disable("debug");
  fmt.Printf("debug=%t\n", app.Get("debug"))
  app.Enable("debug");
  fmt.Printf("debug=%t\n", app.Get("debug"))
}
```

Constructor

```go
package main

import "fmt"
import "configurable"

func main() {
  cfg := configurable.New(map[string]interface{}{
    "env": "development"
  })

  fmt.Printf("env=%t\n", cfg.Get("env"))
  cfg.Set("env", "production")
  fmt.Printf("env=%t\n", cfg.Get("env"))
}
```

### TODO

  - Ask people who know Go what I am doing wrong! :-)

### License

MIT
