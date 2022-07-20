package hooks

import (
	"errors"
	"reflect"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mitchellh/mapstructure"
)

// StringToAddressHookFunc is a mapstructure hook function that converts a string to a common.Address.
func StringToAddressHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data any,
	) (any, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}

		if t != reflect.TypeOf(common.Address{}) {
			return data, nil
		}

		if !common.IsHexAddress(data.(string)) {
			return data, errors.New("invalid address")
		}

		// Convert it by parsing
		return common.HexToAddress(data.(string)), nil
	}
}

// StringToDurationHookFunc is a mapstructure hook function that converts a string to a common.Address.
func StringToDurationHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data any,
	) (any, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}

		if t != reflect.TypeOf(time.Duration(0)) {
			return data, nil
		}

		// Convert it by parsing
		return time.ParseDuration(data.(string))
	}
}

// StringToLipglossColorHookFunc is a mapstructure hook function that converts a string to a lipgloss.Color.
func StringToLipglossColorHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data any,
	) (any, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}

		if t != reflect.TypeOf(lipgloss.Color("")) {
			return data, nil
		}

		// convert it by parsing
		return lipgloss.Color(data.(string)), nil
	}
}
