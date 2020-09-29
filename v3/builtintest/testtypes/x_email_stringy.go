// A derived string-based type compatible with marshalling and database APIs.
//
// Generated from types/stringy.tpl with Type=Email
// options: SortableSlice:true
// by runtemplate v3.7.1
// See https://github.com/rickb777/runtemplate/blob/master/v3/BUILTIN.md

package testtypes

import (
	"errors"
	"sort"
	"strings"
	"database/sql/driver"
	"fmt"
)

// Email is a specialised kind of string.
type Email string

// Ptr returns the address of a Email.
func (email Email) Ptr() *Email {
	return &email
}

// String converts to a string and implements fmt.Stringer.
func (email Email) String() string {
	return string(email)
}

// TrimSpace removes surrounding whitespace.
func (email Email) TrimSpace() Email {
	return Email(strings.TrimSpace(email.String()))
}

// ToLower converts the value to lowercase.
func (email Email) ToLower() Email {
	return Email(strings.ToLower(string(email)))
}

// ToUpper converts the value to uppercase.
func (email Email) ToUpper() Email {
	return Email(strings.ToUpper(string(email)))
}

//-------------------------------------------------------------------------------------------------

// Scan parses some value. It implements sql.Scanner,
// https://golang.org/pkg/database/sql/#Scanner
func (email *Email) Scan(value interface{}) error {
	if value == nil {
		*email = Email("")
		return nil
	}

	switch value.(type) {
	case string:
		*email = Email(value.(string))
	case []byte:
		*email = Email(string(value.([]byte)))
	case nil:
	default:
		return errors.New(fmt.Sprintf("Email.Scan(%#v)", value))
	}
	return nil
}

// Value converts the value to a string. It implements driver.Valuer,
// https://golang.org/pkg/database/sql/driver/#Valuer
func (email Email) Value() (driver.Value, error) {
	return string(email), nil
}

//-------------------------------------------------------------------------------------------------

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
// https://golang.org/pkg/encoding/#TextMarshaler
func (email Email) MarshalText() (text []byte, err error) {
	return []byte(email.String()), nil
}

// UnmarshalText converts transmitted values to ordinary values.
// https://golang.org/pkg/encoding/#TextUnmarshaler
func (email *Email) UnmarshalText(text []byte) error {
	return email.Scan(text)
}

//-------------------------------------------------------------------------------------------------

// EmailSlice attaches the methods of sort.Interface to []Email, sorting in increasing order.
type EmailSlice []Email

func (p EmailSlice) Len() int           { return len(p) }
func (p EmailSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p EmailSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// SortedN is a convenience method.
func (p EmailSlice) Sorted() { sort.Sort(p) }

