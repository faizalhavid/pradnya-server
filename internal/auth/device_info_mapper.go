package auth

import (
	"crypto/sha256"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/mssola/user_agent"
)

func GetIPAddress(r *http.Request) string {
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		rips := strings.Split(forwarded, ",")
		return strings.TrimSpace(rips[0]) // get first IP
	}

	// cek header X-Real-IP
	realIP := r.Header.Get("X-Real-IP")
	if realIP != "" {
		return realIP
	}

	// falback remote address
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

func GetDeviceFingerprint(r *http.Request) string {
	// get from header frotn-end
	clientFingerprint := r.Header.Get("X-Device-Fingerprint")
	if clientFingerprint != "" {
		return clientFingerprint
	}

	// fallback generate fingerprint from user-agent and ip address
	ua := r.UserAgent()
	ip := GetIPAddress(r)
	lang := r.Header.Get("Accept-Language")

	rawString := fmt.Sprintf("%s|%s|%s", ua, ip, lang)
	hash := sha256.Sum256([]byte(rawString))
	return fmt.Sprintf("%x", hash)
}

func BuildDeviceInfo(userAgent string, ipAddress string, userID string) DeviceInfo {
	ua := user_agent.New(userAgent)

	osInfo := ua.OSInfo().Name + " " + ua.OSInfo().Version
	browserName, browserVer := ua.Browser()

	deviceType := "Desktop"
	if ua.Mobile() {
		deviceType = "Mobile"
	}
	deviceName := fmt.Sprintf("%s (%s %s)", deviceType, browserName, browserVer)

	// Fingerprint sederhana (kombinasi UA + IP)
	rawString := fmt.Sprintf("%s|%s", userAgent, ipAddress)
	hash := sha256.Sum256([]byte(rawString))

	return DeviceInfo{
		DeviceName:  deviceName,
		DeviceOS:    osInfo,
		LastLoginAt: time.Now(),
		IpAddress:   ipAddress,
		UserAgent:   userAgent,
		Fingerprint: fmt.Sprintf("%x", hash),
		UserId:      userID,
	}
}
