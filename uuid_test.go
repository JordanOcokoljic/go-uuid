package uuid

import (
	"bytes"
	"strings"
	"testing"
)

func TestParseTypeDecider(t *testing.T) {
	expected := UUID{
		0x3d, 0x83, 0x38, 0xc3,
		0x71, 0x06,
		0x4d, 0x89,
		0xa8, 0x13,
		0x9f, 0xd9, 0xb7, 0x95, 0x5b, 0x55,
	}

	str := "3d8338c3-7106-4d89-a813-9fd9b7955b55"

	uuid, err := Parse(str)
	if err != nil {
		t.Fatalf("An error occured: %e", err)
	}

	if !Equal(expected, uuid) {
		t.Errorf("String failed, expected '%s', got '%s'", expected, uuid)
	}

	b := []byte{
		0x3d, 0x83, 0x38, 0xc3,
		0x71, 0x06,
		0x4d, 0x89,
		0xa8, 0x13,
		0x9f, 0xd9, 0xb7, 0x95, 0x5b, 0x55,
	}

	uuid, err = Parse(b)
	if err != nil {
		t.Fatalf("An error occured: %e", err)
	}

	if !Equal(expected, uuid) {
		t.Errorf("Byte failed, expected '%s', got '%s'", expected, uuid)
	}
}

func TestUUIDString(t *testing.T) {
	expected := "3d8338c3-7106-4d89-a813-9fd9b7955b55"
	sample := UUID{
		0x3d, 0x83, 0x38, 0xc3,
		0x71, 0x06,
		0x4d, 0x89,
		0xa8, 0x13,
		0x9f, 0xd9, 0xb7, 0x95, 0x5b, 0x55,
	}

	if sample.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, sample)
	}
}

func TestParseCanonicalWithValidString(t *testing.T) {
	sample := "3d8338c3-7106-4d89-a813-9fd9b7955b55"
	expected := UUID{
		0x3d, 0x83, 0x38, 0xc3,
		0x71, 0x06,
		0x4d, 0x89,
		0xa8, 0x13,
		0x9f, 0xd9, 0xb7, 0x95, 0x5b, 0x55,
	}

	uuid, err := parseCanonical(sample)

	if err != nil {
		t.Fatalf("An error occured: %e", err)
	}

	if uuid.String() != expected.String() {
		t.Errorf("Expected %s, got %s", expected, uuid)
	}
}

func TestParseCanonicalWithInvalidLengthString(t *testing.T) {
	sample := "3d8338c3-7106-4d89-a813-9fd9b7955b550"
	expected := "uuid: invalid parse length"

	_, err := parseCanonical(sample)

	if err == nil {
		t.Fatal("No error was returned")
	}

	if !strings.HasPrefix(err.Error(), expected) {
		t.Errorf("Expected '%s', got '%s'", expected, err.Error())
	}
}

func TestParseCanonicalWithMalformedString(t *testing.T) {
	sample := "3d8338c37710664d89fa813a9fd9b7955b55"
	expected := "uuid: invalid parse format"

	_, err := parseCanonical(sample)

	if err == nil {
		t.Fatal("No error was returned")
	}

	if !strings.HasPrefix(err.Error(), expected) {
		t.Errorf("Expected '%s', got '%s'", expected, err.Error())
	}
}

func TestParseBytesWithValidSlice(t *testing.T) {
	expected := "3d8338c3-7106-4d89-a813-9fd9b7955b55"
	b := []byte{
		0x3d, 0x83, 0x38, 0xc3,
		0x71, 0x06,
		0x4d, 0x89,
		0xa8, 0x13,
		0x9f, 0xd9, 0xb7, 0x95, 0x5b, 0x55,
	}

	uuid, _ := parseBytes(b)

	if uuid.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, uuid)
	}
}

func TestParseBytesWithInvalidLengthSlice(t *testing.T) {
	expected := "uuid: invalid parse length"
	b := []byte{
		0x3d, 0x83, 0x38, 0xc3,
		0x71, 0x06,
		0x4d, 0x89,
		0xa8, 0x13, 0xaa,
		0x9f, 0xd9, 0xb7, 0x95, 0x5b, 0x55,
	}

	_, err := parseBytes(b)
	if !strings.HasPrefix(err.Error(), expected) {
		t.Errorf("Expected '%s', got '%s'", expected, err.Error())
	}
}

func TestNew(t *testing.T) {
	uuid, _ := New()

	if Equal(uuid, UUID{}) {
		t.Error("Expected generated UUID to not be nil")
	}
}

func TestUUIDByts(t *testing.T) {
	sample := UUID{
		0x3d, 0x83, 0x38, 0xc3,
		0x71, 0x06,
		0x4d, 0x89,
		0xa8, 0x13,
		0x9f, 0xd9, 0xb7, 0x95, 0x5b, 0x55,
	}

	expected := []byte{
		0x3d, 0x83, 0x38, 0xc3,
		0x71, 0x06,
		0x4d, 0x89,
		0xa8, 0x13,
		0x9f, 0xd9, 0xb7, 0x95, 0x5b, 0x55,
	}

	result := sample.Bytes()

	if !bytes.Equal(expected, result) {
		t.Errorf("Expected %#v, got %#v", expected, result)
	}
}

func TestEqualWithSameUUIDs(t *testing.T) {
	u1 := UUID{
		0x3d, 0x83, 0x38, 0xc3,
		0x71, 0x06,
		0x4d, 0x89,
		0xa8, 0x13,
		0x9f, 0xd9, 0xb7, 0x95, 0x5b, 0x55,
	}

	u2 := UUID{
		0x3d, 0x83, 0x38, 0xc3,
		0x71, 0x06,
		0x4d, 0x89,
		0xa8, 0x13,
		0x9f, 0xd9, 0xb7, 0x95, 0x5b, 0x55,
	}

	result := Equal(u1, u2)
	if !result {
		t.Error("Expected 'true', got 'false'")
	}
}

func TestEqualWithDifferentUUIDs(t *testing.T) {
	u1 := UUID{
		0x3d, 0x83, 0x38, 0xc3,
		0x71, 0x06,
		0x4d, 0x89,
		0xa8, 0x13,
		0x9f, 0xd9, 0xb7, 0x95, 0x5b, 0x55,
	}

	u2 := UUID{
		0x3e, 0x83, 0x38, 0xc3,
		0x71, 0x06,
		0x4d, 0x89,
		0xa8, 0x13,
		0x9f, 0xd9, 0xb7, 0x95, 0x5b, 0x55,
	}

	result := Equal(u1, u2)
	if result {
		t.Error("Expected 'false', got 'true'")
	}
}

func TestUUIDSQLScanTypeDecider(t *testing.T) {
	expected := UUID{
		0x3d, 0x83, 0x38, 0xc3,
		0x71, 0x06,
		0x4d, 0x89,
		0xa8, 0x13,
		0x9f, 0xd9, 0xb7, 0x95, 0x5b, 0x55,
	}

	str := "3d8338c3-7106-4d89-a813-9fd9b7955b55"

	strUUID := UUID{}
	err := strUUID.Scan(str)
	if err != nil {
		t.Fatalf("An error occured: %e", err)
	}

	if !Equal(expected, strUUID) {
		t.Errorf("Expected '%s', got '%s'", expected, strUUID)
	}

	b := []byte{
		0x3d, 0x83, 0x38, 0xc3,
		0x71, 0x06,
		0x4d, 0x89,
		0xa8, 0x13,
		0x9f, 0xd9, 0xb7, 0x95, 0x5b, 0x55,
	}

	bUUID := UUID{}
	err = bUUID.Scan(b)
	if err != nil {
		t.Fatalf("An error occured: %e", err)
	}

	if !Equal(expected, bUUID) {
		t.Errorf("Expected '%s', got '%s'", expected, bUUID)
	}

	nilExpected := UUID{}
	nilUUID := UUID{}
	err = nilUUID.Scan(nil)
	if err != nil {
		t.Fatalf("An error occured: %e", err)
	}

	if !Equal(nilExpected, nilUUID) {
		t.Errorf("Expected '%s', got '%s'", nilExpected, nilUUID)
	}
}

func TestUUIDSQLValue(t *testing.T) {
	uuid := UUID{
		0x3d, 0x83, 0x38, 0xc3,
		0x71, 0x06,
		0x4d, 0x89,
		0xa8, 0x13,
		0x9f, 0xd9, 0xb7, 0x95, 0x5b, 0x55,
	}

	expected := "3d8338c3-7106-4d89-a813-9fd9b7955b55"
	actual, _ := uuid.Value()

	if expected != actual {
		t.Errorf("Expected '%s', got '%s'", expected, actual)
	}
}

func TestUUIDMarshalText(t *testing.T) {
	uuid := UUID{
		0x3d, 0x83, 0x38, 0xc3,
		0x71, 0x06,
		0x4d, 0x89,
		0xa8, 0x13,
		0x9f, 0xd9, 0xb7, 0x95, 0x5b, 0x55,
	}

	slice, _ := uuid.MarshalText()

	if string(slice[:]) != uuid.String() {
		t.Errorf("Expected '%s', got '%s'", uuid, string(slice[:]))
	}
}

func TestScanBytes(t *testing.T) {
	b := []byte{
		0x3d, 0x83, 0x38, 0xc3,
		0x71, 0x06,
		0x4d, 0x89,
		0xa8, 0x13,
		0x9f, 0xd9, 0xb7, 0x95, 0x5b, 0x55,
	}

	expected := UUID{
		0x3d, 0x83, 0x38, 0xc3,
		0x71, 0x06,
		0x4d, 0x89,
		0xa8, 0x13,
		0x9f, 0xd9, 0xb7, 0x95, 0x5b, 0x55,
	}

	u := UUID{}
	err := scanBytes(&u, b)
	if err != nil {
		t.Fatalf("An error occured: %e", err)
	}

	if !Equal(expected, u) {
		t.Errorf("Expected '%s', got '%s'", expected, u)
	}
}

func TestScanString(t *testing.T) {
	str := "3d8338c3-7106-4d89-a813-9fd9b7955b55"

	expected := UUID{
		0x3d, 0x83, 0x38, 0xc3,
		0x71, 0x06,
		0x4d, 0x89,
		0xa8, 0x13,
		0x9f, 0xd9, 0xb7, 0x95, 0x5b, 0x55,
	}

	u := UUID{}
	err := scanString(&u, str)
	if err != nil {
		t.Fatalf("An error occured: %e", err)
	}

	if !Equal(expected, u) {
		t.Errorf("Expected '%s', got '%s'", expected, u)
	}
}

func TestScan(t *testing.T) {
	expected := UUID{
		0x3d, 0x83, 0x38, 0xc3,
		0x71, 0x06,
		0x4d, 0x89,
		0xa8, 0x13,
		0x9f, 0xd9, 0xb7, 0x95, 0x5b, 0x55,
	}

	str := "3d8338c3-7106-4d89-a813-9fd9b7955b55"

	strUUID := UUID{}
	err := scan(&strUUID, str)
	if err != nil {
		t.Fatalf("An error occured: %e", err)
	}

	if !Equal(expected, strUUID) {
		t.Errorf("Expected '%s', got '%s'", expected, strUUID)
	}

	b := []byte{
		0x3d, 0x83, 0x38, 0xc3,
		0x71, 0x06,
		0x4d, 0x89,
		0xa8, 0x13,
		0x9f, 0xd9, 0xb7, 0x95, 0x5b, 0x55,
	}

	bUUID := UUID{}
	err = scan(&bUUID, b)
	if err != nil {
		t.Fatalf("An error occured: %e", err)
	}

	if !Equal(expected, bUUID) {
		t.Errorf("Expected '%s', got '%s'", expected, bUUID)
	}
}

func TestEncodeHex(t *testing.T) {
	u := UUID{
		0x3d, 0x83, 0x38, 0xc3,
		0x71, 0x06,
		0x4d, 0x89,
		0xa8, 0x13,
		0x9f, 0xd9, 0xb7, 0x95, 0x5b, 0x55,
	}

	hexs := encodeHex(u)
	for _, pos := range []int{8, 13, 18, 23} {
		if hexs[pos] != '-' {
			t.Errorf("'-' missing at position %d", pos)
		}
	}

	expected := []byte{
		0x33, 0x64, 0x38, 0x33, 0x33, 0x38, 0x63, 0x33, 0x2d,
		0x37, 0x31, 0x30, 0x36, 0x2d,
		0x34, 0x64, 0x38, 0x39, 0x2d,
		0x61, 0x38, 0x31, 0x33, 0x2d,
		0x39, 0x66, 0x64, 0x39, 0x62, 0x37, 0x39, 0x35, 0x35, 0x62, 0x35, 0x35,
	}

	if !bytes.Equal(expected, hexs) {
		t.Errorf("Expected %#v, got %#v", expected, hexs)
	}
}
