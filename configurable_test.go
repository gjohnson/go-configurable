package configurable

import "testing"

type App struct {
  Settings
}

var options = map[string]interface{}{
  "host": "localhost",
  "port": 3000,
  "debug": false,
  "silent": true,
}

func TestNewDefaults(t *testing.T) {
  config := New(options)

  if "localhost" != config.Get("host") {
    t.Error("expected default host")
  }
  if 3000 != config.Get("port") {
    t.Error("expected default port")
  }
}

func TestNewSet(t *testing.T) {
  config := New(options)
  config.Set("port", 8080)

  if 8080 != config.Get("port") {
    t.Error("expected new value to be assigned")
  }
}

func TestNewEnable(t *testing.T) {
  config := New(options)
  config.Enable("debug")

  if true != config.Get("debug") {
    t.Error("expected enable to set value to true")
  }
}

func TestNewDisable(t *testing.T) {
  config := New(options)
  config.Disable("silent")

  if false != config.Get("silent") {
    t.Error("expected disable to set value to false")
  }
}

func TestNewDisabled(t *testing.T) {
  config := New(options)

  if true != config.Disabled("debug") {
    t.Error("expected to falsey")
  }
}

func TestNewEnabled(t *testing.T) {
  config := New(options)

  if false != config.Disabled("silent") {
    t.Error("expected to be truthy")
  }
}
