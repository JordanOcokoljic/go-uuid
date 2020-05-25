package uuid

import (
	"bytes"
	"crypto/rand"
	"database/sql/driver"
	"encoding/hex"
	"fmt"
	"io"
)

// A UUID is a 16 byte Universally Unqiue Identifier as defined in RFC4122.
type UUID [16]byte

// String implements the Stringer interface and returns the UUID in the form
// of: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
func (uuid UUID) String() string {
	buf := encodeHex(uuid)
	return string(buf[:])
}

// Bytes returns the underlying bytes the UUID represents.
func (uuid UUID) Bytes() []byte {
	return uuid[:]
}

// Scan implements the sql.Scanner interface so that UUIDs can be read from
// databases so support SQL operations.
func (uuid *UUID) Scan(src interface{}) error {
	switch src := src.(type) {
	case nil:
		return nil
	case []byte:
		return scanBytes(uuid, src)
	case string:
		return scanString(uuid, src)
	default:
		return fmt.Errorf("uuid: unable to scan type %T", src)
	}
}

// scanBytes scans the incoming byte collection from the sql.Scanner into
// a UUID.
func scanBytes(uuid *UUID, src []byte) error {
	if len(src) == 16 {
		return scan(uuid, src)
	}

	return scanString(uuid, string(src))
}

// scanString scans the incoming text value from the sql.Scanner into
/// a UUID.
func scanString(uuid *UUID, src string) error {
	if src == "" {
		return nil
	}

	return scan(uuid, src)
}

// scan will parse an incoming value in a UUID, and assign into the provided
// pointer.
func scan(uuid *UUID, src interface{}) error {
	u, err := Parse(src)
	if err != nil {
		return err
	}

	*uuid = u
	return nil
}

// Value implements the sql.Valuer so that UUIDs can be written to databases
// to support SQL operations.
func (uuid UUID) Value() (driver.Value, error) {
	return uuid.String(), nil
}

// MarshalText implements the encoding.TextMarshaler interface.
func (uuid UUID) MarshalText() ([]byte, error) {
	buf := encodeHex(uuid)
	return buf[:], nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (uuid *UUID) UnmarshalText(text []byte) error {
	parsed, err := Parse(text)
	if err != nil {
		return err
	}

	*uuid = parsed
	return nil
}

// Parse decodes string s into a UUID or returns an error. It currently only
// supports UUIDs in the format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
func Parse(src interface{}) (UUID, error) {
	switch src := src.(type) {
	case string:
		return parseCanonical(src)
	case []byte:
		return parseBytes(src)
	default:
		err := fmt.Errorf("uuid: could not parse uuid from type %T", src)
		return UUID{}, err
	}
}

// parseCanonical will parse a UUID in the form of:
// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
func parseCanonical(s string) (UUID, error) {
	if len(s) != 36 {
		return UUID{}, fmt.Errorf("uuid: invalid parse length %s", s)
	}

	if s[8] != '-' || s[13] != '-' || s[18] != '-' || s[23] != '-' {
		return UUID{}, fmt.Errorf("uuid: invalid parse format %s", s)
	}

	u := UUID{}
	src := []byte(s)[:]
	dst := u[:]

	for i, group := range []int{8, 4, 4, 4, 12} {
		if i > 0 {
			src = src[1:]
		}

		_, err := hex.Decode(dst[:group/2], src[:group])
		if err != nil {
			return UUID{}, err
		}

		src = src[group:]
		dst = dst[group/2:]
	}

	return u, nil
}

// ParseBytes will parse a collection of bytes into a UUID.
func parseBytes(b []byte) (UUID, error) {
	uuid := UUID{}

	if len(b) != 16 {
		err := fmt.Errorf("uuid: invalid parse length (%d) on %#v", len(b), b)
		return UUID{}, err
	}

	copy(uuid[:], b)
	return uuid, nil
}

// New returns a random V4 UUID.
func New() (UUID, error) {
	u := UUID{}
	_, err := io.ReadFull(rand.Reader, u[:])
	if err != nil {
		return UUID{}, err
	}

	u[6] = (u[6] & 0x0f) | 0x40
	u[8] = (u[8] & 0x3f) | 0x80
	return u, nil
}

// Must behaves the same as New, but catches any errors that occur and
// panics. This simplifies declaration and initlization.
func Must() UUID {
	u, err := New()
	if err != nil {
		panic(err)
	}

	return u
}

// Equal returns true if the two UUIDs are equal.
func Equal(u1, u2 UUID) bool {
	return bytes.Equal(u1.Bytes(), u2.Bytes())
}

// encodeHex returns a hex encoded byte slice representing the UUID.
func encodeHex(uuid UUID) []byte {
	buf := make([]byte, 36)
	hex.Encode(buf, uuid[:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], uuid[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], uuid[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], uuid[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:], uuid[10:])
	return buf
}
