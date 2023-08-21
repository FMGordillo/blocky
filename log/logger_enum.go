// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package log

import (
	"fmt"
	"strings"
)

const (
	// FormatTypeText is a FormatType of type Text.
	// logging as text
	FormatTypeText FormatType = iota
	// FormatTypeJson is a FormatType of type Json.
	// JSON format
	FormatTypeJson
)

var ErrInvalidFormatType = fmt.Errorf("not a valid FormatType, try [%s]", strings.Join(_FormatTypeNames, ", "))

const _FormatTypeName = "textjson"

var _FormatTypeNames = []string{
	_FormatTypeName[0:4],
	_FormatTypeName[4:8],
}

// FormatTypeNames returns a list of possible string values of FormatType.
func FormatTypeNames() []string {
	tmp := make([]string, len(_FormatTypeNames))
	copy(tmp, _FormatTypeNames)
	return tmp
}

var _FormatTypeMap = map[FormatType]string{
	FormatTypeText: _FormatTypeName[0:4],
	FormatTypeJson: _FormatTypeName[4:8],
}

// String implements the Stringer interface.
func (x FormatType) String() string {
	if str, ok := _FormatTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("FormatType(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x FormatType) IsValid() bool {
	_, ok := _FormatTypeMap[x]
	return ok
}

var _FormatTypeValue = map[string]FormatType{
	_FormatTypeName[0:4]: FormatTypeText,
	_FormatTypeName[4:8]: FormatTypeJson,
}

// ParseFormatType attempts to convert a string to a FormatType.
func ParseFormatType(name string) (FormatType, error) {
	if x, ok := _FormatTypeValue[name]; ok {
		return x, nil
	}
	return FormatType(0), fmt.Errorf("%s is %w", name, ErrInvalidFormatType)
}

// MarshalText implements the text marshaller method.
func (x FormatType) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *FormatType) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseFormatType(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// LevelInfo is a Level of type Info.
	LevelInfo Level = iota
	// LevelTrace is a Level of type Trace.
	LevelTrace
	// LevelDebug is a Level of type Debug.
	LevelDebug
	// LevelWarn is a Level of type Warn.
	LevelWarn
	// LevelError is a Level of type Error.
	LevelError
	// LevelFatal is a Level of type Fatal.
	LevelFatal
)

var ErrInvalidLevel = fmt.Errorf("not a valid Level, try [%s]", strings.Join(_LevelNames, ", "))

const _LevelName = "infotracedebugwarnerrorfatal"

var _LevelNames = []string{
	_LevelName[0:4],
	_LevelName[4:9],
	_LevelName[9:14],
	_LevelName[14:18],
	_LevelName[18:23],
	_LevelName[23:28],
}

// LevelNames returns a list of possible string values of Level.
func LevelNames() []string {
	tmp := make([]string, len(_LevelNames))
	copy(tmp, _LevelNames)
	return tmp
}

var _LevelMap = map[Level]string{
	LevelInfo:  _LevelName[0:4],
	LevelTrace: _LevelName[4:9],
	LevelDebug: _LevelName[9:14],
	LevelWarn:  _LevelName[14:18],
	LevelError: _LevelName[18:23],
	LevelFatal: _LevelName[23:28],
}

// String implements the Stringer interface.
func (x Level) String() string {
	if str, ok := _LevelMap[x]; ok {
		return str
	}
	return fmt.Sprintf("Level(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x Level) IsValid() bool {
	_, ok := _LevelMap[x]
	return ok
}

var _LevelValue = map[string]Level{
	_LevelName[0:4]:   LevelInfo,
	_LevelName[4:9]:   LevelTrace,
	_LevelName[9:14]:  LevelDebug,
	_LevelName[14:18]: LevelWarn,
	_LevelName[18:23]: LevelError,
	_LevelName[23:28]: LevelFatal,
}

// ParseLevel attempts to convert a string to a Level.
func ParseLevel(name string) (Level, error) {
	if x, ok := _LevelValue[name]; ok {
		return x, nil
	}
	return Level(0), fmt.Errorf("%s is %w", name, ErrInvalidLevel)
}

// MarshalText implements the text marshaller method.
func (x Level) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *Level) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseLevel(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
