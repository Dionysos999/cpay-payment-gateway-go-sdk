package sdk

import (
	"strings"
	"testing"
)

// A valid 32-character salt for testing
const validSalt = "12345678901234567890123456789012"

func TestGenSign_ValidSignature(t *testing.T) {
	params := map[string]interface{}{
		"merchant_id": int64(12345),
		"amount":      100.50,
		"currency":    "USD",
		"order_id":    "test-order-123",
	}

	sign, err := GenSign(params, validSalt)
	if err != nil {
		t.Fatalf("GenSign failed: %v", err)
	}

	if len(sign) == 0 {
		t.Fatal("Expected non-empty signature")
	}

	// HMAC-SHA256 produces 64 hex characters
	if len(sign) != 64 {
		t.Errorf("Expected 64 character signature, got %d", len(sign))
	}
}

func TestGenSign_SignParameterIgnored(t *testing.T) {
	params := map[string]interface{}{
		"merchant_id": int64(12345),
		"amount":      100.50,
		"sign":        "should-be-ignored",
	}

	sign1, err1 := GenSign(params, validSalt)
	if err1 != nil {
		t.Fatalf("GenSign failed: %v", err1)
	}

	// Remove the sign parameter
	delete(params, "sign")
	sign2, err2 := GenSign(params, validSalt)
	if err2 != nil {
		t.Fatalf("GenSign failed: %v", err2)
	}

	// Signatures should be identical
	if sign1 != sign2 {
		t.Error("Sign parameter should be ignored, but signatures differ")
	}
}

func TestGenSign_EmptySalt(t *testing.T) {
	params := map[string]interface{}{
		"merchant_id": int64(12345),
	}

	_, err := GenSign(params, "")
	if err == nil {
		t.Fatal("Expected error for empty salt")
	}
	if !strings.Contains(err.Error(), "cannot be empty") {
		t.Errorf("Expected 'cannot be empty' error, got: %v", err)
	}
}

func TestGenSign_WeakSalt(t *testing.T) {
	params := map[string]interface{}{
		"merchant_id": int64(12345),
	}

	weakSalts := []string{
		"short",                     // 5 chars
		"1234567890123456789012345", // 25 chars
		"abc123",                    // 6 chars
	}

	for _, salt := range weakSalts {
		_, err := GenSign(params, salt)
		if err == nil {
			t.Errorf("Expected error for weak salt '%s'", salt)
		}
		if !strings.Contains(err.Error(), "too weak") {
			t.Errorf("Expected 'too weak' error for salt '%s', got: %v", salt, err)
		}
	}
}

func TestGenSign_NilMap(t *testing.T) {
	_, err := GenSign(nil, validSalt)
	if err == nil {
		t.Fatal("Expected error for nil map")
	}
	if !strings.Contains(err.Error(), "cannot be nil") {
		t.Errorf("Expected 'cannot be nil' error, got: %v", err)
	}
}

func TestGenSign_EmptyMap(t *testing.T) {
	params := map[string]interface{}{}

	_, err := GenSign(params, validSalt)
	if err == nil {
		t.Fatal("Expected error for empty map")
	}
	if !strings.Contains(err.Error(), "cannot be empty") {
		t.Errorf("Expected 'cannot be empty' error, got: %v", err)
	}
}

func TestGenSign_TooManyParameters(t *testing.T) {
	params := make(map[string]interface{})
	for i := 0; i <= 1000; i++ {
		params[("key_" + string(rune(i)))] = i
	}

	_, err := GenSign(params, validSalt)
	if err == nil {
		t.Fatal("Expected error for too many parameters")
	}
	if !strings.Contains(err.Error(), "too many") {
		t.Errorf("Expected 'too many' error, got: %v", err)
	}
}

func TestGenSign_WhitespaceOnlyValue(t *testing.T) {
	params := map[string]interface{}{
		"merchant_id": int64(12345),
		"amount":      "   ", // whitespace only
		"currency":    "USD",
	}

	sign, err := GenSign(params, validSalt)
	if err != nil {
		t.Fatalf("GenSign failed: %v", err)
	}

	// Whitespace-only values should be trimmed and not included
	if strings.Contains(sign, "amount") {
		t.Error("Whitespace-only value should not be included in signature")
	}
}

func TestGenSign_InvalidParameterKey(t *testing.T) {
	invalidKeys := []string{
		"key with spaces",   // spaces
		"key\twith\ttabs",   // tabs
		"key\nwith\nnewlin", // newlines
		"key@with!special#", // special chars
	}

	for _, key := range invalidKeys {
		params := map[string]interface{}{
			key: "value",
		}

		_, err := GenSign(params, validSalt)
		if err == nil {
			t.Errorf("Expected error for invalid key '%s'", key)
		}
		if !strings.Contains(err.Error(), "invalid parameter key") {
			t.Errorf("Expected 'invalid parameter key' error, got: %v", err)
		}
	}
}

func TestGenSign_UnsupportedType(t *testing.T) {
	params := map[string]interface{}{
		"valid_param": "string",
		"invalid":     []string{"array", "type"},
	}

	_, err := GenSign(params, validSalt)
	if err == nil {
		t.Fatal("Expected error for unsupported type")
	}
	if !strings.Contains(err.Error(), "unsupported parameter type") {
		t.Errorf("Expected 'unsupported parameter type' error, got: %v", err)
	}
}

func TestGenSign_ValidParameterKey(t *testing.T) {
	validKeys := []string{
		"normal_key",
		"key-with-dashes",
		"key.with.dots",
		"key_with_underscores",
		"UpperCaseKey",
		"key123",
		"123key",
	}

	for _, key := range validKeys {
		if !isValidParameterKey(key) {
			t.Errorf("Key '%s' should be valid", key)
		}
	}
}

func TestGenSign_SortedKeys(t *testing.T) {
	params := map[string]interface{}{
		"z_param": "last",
		"a_param": "first",
		"m_param": "middle",
	}

	sign, err := GenSign(params, validSalt)
	if err != nil {
		t.Fatalf("GenSign failed: %v", err)
	}

	if len(sign) == 0 {
		t.Fatal("Expected non-empty signature")
	}
}

func TestGenSign_DifferentTypes(t *testing.T) {
	params := map[string]interface{}{
		"string_val":  "hello",
		"int_val":     42,
		"int64_val":   int64(1234567890),
		"float32_val": float32(3.14),
		"float64_val": 99.99,
		"bool_val":    true,
	}

	sign, err := GenSign(params, validSalt)
	if err != nil {
		t.Fatalf("GenSign failed: %v", err)
	}

	if len(sign) == 0 {
		t.Fatal("Expected non-empty signature")
	}
}

func TestGenSign_Deterministic(t *testing.T) {
	params := map[string]interface{}{
		"merchant_id": int64(12345),
		"amount":      100.50,
		"currency":    "USD",
	}

	sign1, err1 := GenSign(params, validSalt)
	if err1 != nil {
		t.Fatalf("GenSign failed: %v", err1)
	}

	sign2, err2 := GenSign(params, validSalt)
	if err2 != nil {
		t.Fatalf("GenSign failed: %v", err2)
	}

	if sign1 != sign2 {
		t.Error("Signature should be deterministic")
	}
}

func TestIsValidParameterKey(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want bool
	}{
		{"empty string", "", false},
		{"too long key", strings.Repeat("a", 257), false},
		{"normal key", "valid_key", true},
		{"with dashes", "valid-key", true},
		{"with dots", "valid.key", true},
		{"with spaces", "invalid key", false},
		{"with special chars", "invalid@key", false},
		{"with newlines", "invalid\nkey", false},
		{"mixed case", "ValidKey123", true},
		{"only letters", "abc", true},
		{"only digits", "123", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidParameterKey(tt.key)
			if got != tt.want {
				t.Errorf("isValidParameterKey(%q) = %v, want %v", tt.key, got, tt.want)
			}
		})
	}
}

func TestValueToString(t *testing.T) {
	tests := []struct {
		name    string
		value   interface{}
		want    string
		wantErr bool
	}{
		{"string", "hello", "hello", false},
		{"int", 42, "42", false},
		{"int64", int64(1234567890), "1234567890", false},
		{"float32", float32(3.14), "3.14", false},
		{"float64", 99.99, "99.99", false},
		{"bool true", true, "true", false},
		{"bool false", false, "false", false},
		{"unsupported", []int{1, 2, 3}, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := valueToString(tt.value, "test_key")
			if (err != nil) != tt.wantErr {
				t.Errorf("valueToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("valueToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenSign_SignatureUniqueness(t *testing.T) {
	params1 := map[string]interface{}{
		"merchant_id": int64(12345),
		"amount":      100.50,
	}

	params2 := map[string]interface{}{
		"merchant_id": int64(12345),
		"amount":      100.51, // Different amount
	}

	sign1, err1 := GenSign(params1, validSalt)
	if err1 != nil {
		t.Fatalf("GenSign failed: %v", err1)
	}

	sign2, err2 := GenSign(params2, validSalt)
	if err2 != nil {
		t.Fatalf("GenSign failed: %v", err2)
	}

	if sign1 == sign2 {
		t.Error("Different parameters should produce different signatures")
	}
}

func TestGenSign_FloatPrecision(t *testing.T) {
	params := map[string]interface{}{
		"amount": 100.123456789012345, // High precision float
	}

	sign, err := GenSign(params, validSalt)
	if err != nil {
		t.Fatalf("GenSign failed: %v", err)
	}

	// The decimal conversion should handle high precision without error
	if len(sign) != 64 {
		t.Errorf("Expected 64 character signature, got %d", len(sign))
	}
}
