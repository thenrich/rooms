package conf

import (
	"net/http"
	"net/url"
	"fmt"
	"github.com/thenrich/rooms/util"
	"encoding/base64"
	"sort"
)

type TwilioConfig struct {
	// Twilio API Key
	Key []byte

	// Base HTTP URL
	BaseUrl string

	// Number of digits for conference room
	ConferenceRoomNumDigits int64

	// Timeout in seconds to enter room number
	ConferenceRoomTimeout int64
}

// Create a new Twilio config
// key: Twilio API key
// baseUrl: URL of the room API -- this is usually a base AppEngine URL like
// https://xyz-appsot.com/
// roomNumDigits: Number of digits for conference rooms
// roomTimeout: Number of seconds to enter room number
func NewTwilioConfig(key []byte, baseUrl string, roomNumDigits int64, roomTimeout int64) *TwilioConfig {
	return &TwilioConfig{
		key,
		baseUrl,
		roomNumDigits,
		roomTimeout,
	}
}

// Return sorted keys
func sortedPostKeys(v url.Values) []string {
	var keys []string
	for k, _ := range v {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}

// Verify Twilio Authorization header
func VerifyTwilioRequest(r *http.Request, key []byte) bool {
	if err := r.ParseForm(); err != nil {
		return false
	}

	// Extract full URL
	//url := "https://" + r.Host + r.URL.String()
	url := r.URL.String()

	if r.Method == "POST" {
		// Iterate sorted form keys and append the key and value to the
		// url
		for _, v := range sortedPostKeys(r.Form) {
			url += fmt.Sprintf("%s%s", v, r.FormValue(v))
		}
	}

	// Decode signature
	decoded, err := base64.StdEncoding.DecodeString(r.Header.Get("X-Twilio-Signature"))
	if err != nil {
		return false
	}

	// Should now have full URL + keyvaluekeyvaluekeyvalue...
	return utils.VerifyHMACSignature(decoded, []byte(url), key)

}
