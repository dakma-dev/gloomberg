package degendb

import (
	"fmt"

	"github.com/benleb/gloomberg/internal/style"
)

type Trait struct {
	Name        string      `json:"trait_type"   mapstructure:"trait_type"`
	Value       interface{} `json:"value"        mapstructure:"value"`
	DisplayType string      `json:"display_type" mapstructure:"display_type"`
	MaxValue    interface{} `json:"max_value"    mapstructure:"max_value"`
	Order       interface{} `json:"order"        mapstructure:"order"`
	TraitCount  float64     `json:"trait_count"  mapstructure:"trait_count"`
}

func (t *Trait) String() string {
	return fmt.Sprintf("%s: %s", t.Name, fmt.Sprint(t.Value))
}

func (t *Trait) StringBold() string {
	return fmt.Sprintf("%s: %s", t.Name, style.BoldAlmostWhite(fmt.Sprint(t.Value)))
}
