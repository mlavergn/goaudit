package goaudit

import (
	"encoding/json"
	"net/http"
	"time"
)

// Version export
const Version = "0.2.0"

// ClientAudit type
type ClientAudit struct {
	Timestamp int64  `json:"timestamp"`
	Host      string `json:"host"`
	URL       string `json:"url"`
	IPAddr    string `json:"ipAddr"`
	UserAgent string `json:"userAgent"`
	Referer   string `json:"referer"`
	Proxy     string `json:"forwarded"`
	XHR       string `json:"xhr"`
	CORS      string `json:"cors"`
}

// NewClientAudit init
func NewClientAudit(req *http.Request) ClientAudit {
	id := ClientAudit{
		Timestamp: time.Now().UTC().Unix(),
		Host:      req.Header.Get("Host"),
		URL:       req.RequestURI,
		IPAddr:    req.RemoteAddr,
		UserAgent: req.UserAgent(),
		Referer:   req.Referer(),
		Proxy:     req.Header.Get("Forwarded"),
		XHR:       req.Header.Get("X-Requested-With"),
		CORS:      req.Header.Get("Origin"),
	}

	return id
}

// JSON helper returns byte encoded JSON data
func (id *ClientAudit) JSON() []byte {
	result, _ := json.Marshal(id)
	return result
}

// String helper returns JSON data as a string
func (id *ClientAudit) String() string {
	return string(id.JSON())
}

// Map helper returns data as a map
func (id *ClientAudit) Map() map[string]interface{} {
	raw := id.JSON()
	var result map[string]interface{}
	json.Unmarshal(raw, &result)
	return result
}
