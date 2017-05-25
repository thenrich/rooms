package conf

import (
	"testing"
	"net/http"
	_ "google.golang.org/appengine/aetest"
	"github.com/stretchr/testify/assert"
	"encoding/base64"
	"crypto/hmac"
	"crypto/sha1"
	"bytes"
)

func TestVerifyTwilioRequest(t *testing.T) {
	// Create sample request
	var body bytes.Buffer
	body.WriteString("key1=value1&key2=value2")
	req, err := http.NewRequest("POST", "https://example.org/path", &body)
	assert.Nil(t, err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Set form values
	err = req.ParseForm()
	assert.Nil(t, err)

	// Append form values to URL and generate signature
	key := []byte("twiliokey")
	h := hmac.New(sha1.New, key)
	h.Write([]byte("https://example.org/pathkey1value1key2value2"))
	sig := base64.StdEncoding.EncodeToString(h.Sum(nil))

	req.Header.Set("X-Twilio-Signature", string(sig))

	assert.True(t, VerifyTwilioRequest(req, key))
}
