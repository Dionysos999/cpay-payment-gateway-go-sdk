package sdk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/shopspring/decimal"
)

// GenSign generates an HMAC-SHA256 signature for the given parameters.
// Returns the signature or an error if validation fails.
func GenSign(p map[string]interface{}, salt string) (string, error) {
	// Validate salt (SecretKey)
	if len(salt) == 0 {
		return "", errors.New("salt/secret key cannot be empty")
	}
	if len(salt) < 32 {
		return "", errors.New("salt/secret key is too weak (minimum 32 characters required)")
	}

	// Validate parameter map
	if p == nil {
		return "", errors.New("parameter map cannot be nil")
	}
	if len(p) == 0 {
		return "", errors.New("parameter map cannot be empty")
	}

	// Maximum parameter limit to prevent potential DoS
	if len(p) > 1000 {
		return "", errors.New("too many parameters (maximum 1000)")
	}

	// Create a copy of the map to avoid modifying the original
	params := make(map[string]interface{}, len(p))
	for k, v := range p {
		if k == "sign" {
			continue
		}
		// Validate parameter key
		if !isValidParameterKey(k) {
			return "", fmt.Errorf("invalid parameter key: %s", k)
		}
		params[k] = v
	}

	counter := len(params)
	ks := make([]string, 0, counter)

	for k := range params {
		ks = append(ks, k)
	}
	sort.Strings(ks)

	qs := ""
	for i := 0; i < counter; i++ {
		var s = params[ks[i]]

		// Safe type conversion with proper error handling
		tmp, err := valueToString(s, ks[i])
		if err != nil {
			return "", fmt.Errorf("failed to convert parameter %s: %w", ks[i], err)
		}

		// Trim whitespace and check if non-empty
		tmp = strings.TrimSpace(tmp)
		if len(tmp) > 0 {
			qs += ks[i] + "=" + tmp + "&"
		}
	}
	qs = qs + fmt.Sprintf("key=%s", salt)
	h := hmac.New(sha256.New, []byte(salt))
	h.Write([]byte(qs))
	return hex.EncodeToString(h.Sum(nil)), nil
}

// valueToString safely converts interface{} values to string representation.
func valueToString(v interface{}, key string) (string, error) {
	switch val := v.(type) {
	case string:
		return val, nil
	case int:
		return strconv.Itoa(val), nil
	case int64:
		return strconv.FormatInt(val, 10), nil
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32), nil
	case float64:
		// Use decimal for precise float conversion
		str := fmt.Sprintf("%v", val)
		dec, err := decimal.NewFromString(str)
		if err != nil {
			return "", fmt.Errorf("decimal conversion failed: %w", err)
		}
		return dec.String(), nil
	case bool:
		return strconv.FormatBool(val), nil
	default:
		return "", fmt.Errorf("unsupported parameter type: %T for key %s", v, key)
	}
}

// isValidParameterKey validates that a parameter key contains only safe characters.
func isValidParameterKey(key string) bool {
	if len(key) == 0 || len(key) > 256 {
		return false
	}
	for _, r := range key {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' && r != '-' && r != '.' {
			return false
		}
	}
	return true
}
