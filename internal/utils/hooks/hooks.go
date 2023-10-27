//nolint:forcetypeassert
package hooks

import (
	"errors"
	"math/big"
	"reflect"
	"strconv"
	"time"

	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mitchellh/mapstructure"
)

var ErrInvalidAddress = errors.New("invalid address")

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
			return data, ErrInvalidAddress
		}

		// Convert it by parsing
		return common.HexToAddress(data.(string)), nil
	}
}

// StringToAddressHookFunc is a mapstructure hook function that converts a string to a common.Address.
func StringToHashHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data any,
	) (any, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}

		if t != reflect.TypeOf(common.Hash{}) {
			return data, nil
		}

		// Convert it by parsing
		return common.HexToHash(data.(string)), nil
	}
}

// StringToAddressHookFunc is a mapstructure hook function that converts a string to a common.Address.
func StringToEventTypeHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data any,
	) (any, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}

		if t.Kind() != reflect.Interface { // TypeOf((degendb.GBEventType{})).K {
			return data, nil
		}

		// Convert it by parsing
		return degendb.GetEventType(data.(string)), nil
	}
}

// // StringToDurationHookFunc is a mapstructure hook function that converts a string to a common.Address.
// func StringToDurationHookFunc() mapstructure.DecodeHookFunc {
// 	return func(
// 		f reflect.Type,
// 		t reflect.Type,
// 		data any,
// 	) (any, error) {
// 		if f.Kind() != reflect.String {
// 			return data, nil
// 		}

// 		if t != reflect.TypeOf(time.Duration(0)) {
// 			return data, nil
// 		}

// 		// Convert it by parsing
// 		return time.ParseDuration(data.(string))
// 	}
// }

// StringToDurationHookFunc is a mapstructure hook function that converts a string to a common.Address.
func StringToUnixTimeHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data any,
	) (any, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}

		if t != reflect.TypeOf(time.Time{}) {
			return data, nil
		}

		// Convert it by parsing
		timestamp, err := strconv.ParseInt(data.(string), 10, 64)
		if err != nil {
			return nil, err
		}

		return time.Unix(timestamp, 0), nil
	}
}

// StringToInt64HookFunc is a mapstructure hook function that converts a string to a common.Address.
func StringToInt64HookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data any,
	) (any, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}

		if t != reflect.TypeOf(int64(0)) {
			return data, nil
		}

		// Convert it by parsing
		return strconv.ParseInt(data.(string), 10, 64)
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

// StringToBigIntHookFunc is a mapstructure hook function that converts a string to a *big.Int.
func StringToBigIntHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data any,
	) (any, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}

		if t != reflect.TypeOf(*big.NewInt(0)) {
			return data, nil
		}

		// convert data string to big.Int
		value, ok := big.NewInt(0).SetString(data.(string), 10)
		if !ok {
			log.Errorf("invalid big.Int string: %+v", data)

			return nil, errors.New("invalid big.Int string")
		}

		return value, nil

		// if balance, err := strconv.ParseInt(data.(string), 10, 64); err == nil {
		// 	return big.NewInt(balance), nil
		// } else {
		// 	return data, err
		// }
	}
}
