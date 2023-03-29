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
	delete(p, "sign")

	counter := len(p)
	ks := make([]string, 0, counter)

	for k := range p {
		ks = append(ks, k)
	}
	sort.Strings(ks)

	qs := ""
	for i := 0; i < counter; i++ {
		var s = p[ks[i]]
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
