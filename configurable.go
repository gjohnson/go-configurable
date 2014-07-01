package configurable

import "sync"

// Settings
type Settings struct {
  options map[string]interface{}
  sync.RWMutex
}

// Returns an new instance of Settings
func New(options map[string]interface{}) (*Settings) {
  cfg := &Settings{}
  cfg.Defaults(options)
  return cfg
}

// Initialize defaults
func (cfg *Settings) Defaults(options map[string]interface{}) {
  cfg.options = make(map[string]interface{})
  for option, value := range options {
    setOption(cfg, option, value)
  }
}

// Sets the option to the value
func (cfg *Settings) Set(option string, value interface{}) {
  setOption(cfg, option, value)
}

// Gets the option
func (cfg *Settings) Get(option string) (interface{}) {
  return getOption(cfg, option)
}

// Sets the option to false
func (cfg *Settings) Disable(option string) {
  setOption(cfg, option, false)
}

// Sets the option to true
func (cfg *Settings) Enable(option string) {
  setOption(cfg, option, true)
}

// Returns wether the option is truthy
func (cfg *Settings) Enabled(option string) (bool) {
  return true == getOption(cfg, option).(bool)
}

// Returns wether the option is falsey
func (cfg *Settings) Disabled(option string) (bool) {
  return false == getOption(cfg, option).(bool)
}

// Internal option getter that takes care of locking
func getOption(cfg *Settings, option string) (interface{}) {
  cfg.Lock()
  defer cfg.Unlock()
  return cfg.options[option]
}

// Internal option setter that takes care of locking
func setOption(cfg *Settings, option string, value interface{}) {
  cfg.Lock()
  cfg.options[option] = value
  cfg.Unlock()
}
