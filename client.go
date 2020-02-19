package goaudit

import (
	"strings"
)

// OS type
type OS string

// OSAndroid const
const (
	OSAndroid OS = "Android"
	OSIPad       = "iPadOS"
	OSIPhone     = "iPhoneOS"
	OSLinux      = "Linux"
	OSMac        = "macOS"
	OSWindows    = "Windows"
	OSUnknown    = "Unknown"
)

// Browser type
type Browser string

// BrowserChrome const
const (
	BrowserChrome  Browser = "Chrome"
	BrowserSafari          = "Safari"
	BrowserEdge            = "Edge"
	BrowserIE              = "IE"
	BrowserUnknown         = "Unknown"
)

// Device type
type Device int

// DeviceIPad const
const (
	DeviceIPad Device = iota
	DeviceIPhone
	DeviceMac
	DeviceAndroid
	DeviceWindows
)

// ClientPlatform type
type ClientPlatform struct {
	OS         string
	OSVer      string
	Device     string
	Browser    string
	BrowserVer string
}

func (id *ClientPlatform) parseOS(platform []string) OS {
	switch platform[0] {
	case "Macintosh":
		return OSMac
	case "iPad":
		return OSIPad
	case "iPhone":
		return OSIPhone
	case "Linux":
		switch platform[1] {
		case "Android":
			return OSAndroid
		default:
			return OSLinux
		}
	case "Windows":
		return OSWindows
	}
	return OSUnknown
}

func (id *ClientPlatform) parseBrowser(moz13 string) Browser {
	if strings.Contains(moz13, "Chrome/") {
		return BrowserChrome
	}
	if strings.Contains(moz13, "Safari/") {
		return BrowserSafari
	}
	if strings.Contains(moz13, "CriOS/") {
		return BrowserChrome
	}
	if strings.Contains(moz13, "Edge/") {
		return BrowserEdge
	}
	if strings.Contains(moz13, "Trident/") {
		return BrowserIE
	}
	return BrowserUnknown
}

// NewClientPlatform init
func NewClientPlatform(userAgent string) ClientPlatform {
	id := ClientPlatform{}

	start := strings.Index(userAgent, "(")
	end := strings.Index(userAgent[start:], ")")
	platform := userAgent[start+1 : start+end]

	pcomp := strings.Split(platform, ";")

	id.OS = string(id.parseOS(pcomp))
	id.Browser = string(id.parseBrowser(userAgent[13:]))

	return id
}
