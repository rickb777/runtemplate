// A derived string-based type compatible with marshalling and database APIs.
//
// Generated from types/stringy.tpl with Type=Category
// options: SortableSlice:<no value>
// by runtemplate v3.5.0
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package testtypes

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

// Category is a specialised kind of string.
type Category string

// Ptr returns the address of a Category.
func (category Category) Ptr() *Category {
	return &category
}

// String converts to a string and implements fmt.Stringer.
func (category Category) String() string {
	return string(category)
}

// TrimSpace removes surrounding whitespace.
func (category Category) TrimSpace() Category {
	return Category(strings.TrimSpace(category.String()))
}

// ToLower converts the value to lowercase.
func (category Category) ToLower() Category {
	return Category(strings.ToLower(string(category)))
}

// ToUpper converts the value to uppercase.
func (category Category) ToUpper() Category {
	return Category(strings.ToUpper(string(category)))
}

//-------------------------------------------------------------------------------------------------

// Scan parses some value. It implements sql.Scanner,
// https://golang.org/pkg/database/sql/#Scanner
func (category *Category) Scan(value interface{}) error {
	if value == nil {
		*category = Category("")
		return nil
	}

	switch value.(type) {
	case string:
		*category = Category(value.(string))
	case []byte:
		*category = Category(string(value.([]byte)))
	case nil:
	default:
		return errors.New(fmt.Sprintf("Category.Scan(%#v)", value))
	}
	return nil
}

// Value converts the value to a string. It implements driver.Valuer,
// https://golang.org/pkg/database/sql/driver/#Valuer
func (category Category) Value() (driver.Value, error) {
	return string(category), nil
}

//-------------------------------------------------------------------------------------------------

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
// https://golang.org/pkg/encoding/#TextMarshaler
func (category Category) MarshalText() (text []byte, err error) {
	return []byte(category.String()), nil
}

// UnmarshalText converts transmitted values to ordinary values.
// https://golang.org/pkg/encoding/#TextUnmarshaler
func (category *Category) UnmarshalText(text []byte) error {
	return category.Scan(text)
}

//-------------------------------------------------------------------------------------------------
