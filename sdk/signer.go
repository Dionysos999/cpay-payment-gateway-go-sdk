package sdk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"

	"github.com/shopspring/decimal"
)

func GenSign(p map[string]interface{}, salt string) string {
	// Create a copy of the map to avoid modifying the original
	params := make(map[string]interface{}, len(p))
	for k, v := range p {
		if k != "sign" {
			params[k] = v
		}
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
		tmp := fmt.Sprintf("%v", s)
		switch s.(type) {
		case int64:
		case int:
		case float64:
			ss, _ := decimal.NewFromString(tmp)
			tmp = ss.String()
		}

		if len(tmp) > 0 {
			qs += ks[i] + "=" + tmp + "&"
		}
	}
	qs = qs + fmt.Sprintf("key=%s", salt)
	h := hmac.New(sha256.New, []byte(salt))
	h.Write([]byte(qs))
	return hex.EncodeToString(h.Sum(nil))
}
