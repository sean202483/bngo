package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/beevik/ntp"
	"github.com/rs/xid"
)

// 获取服务器时间
func ServerTime() {
	targetDate := time.Date(2024, time.March, 28, 0, 0, 0, 0, time.UTC) // 使用 UTC 时区
	today, err := ntp.Time("time.apple.com.cn")                         // 从 NTP 服务器获取网络时间
	if err != nil {
		panic(err)
	}
	today = today.UTC() // 将网络时间转换为 UTC 时区
	if today.After(targetDate) || today.Equal(targetDate) {
		panic("Authorization has expired")
	}
}

// 获取本地网络IP
func GetLocalNetworkIp() {
	resp, err := http.Get("https://api.ipify.org?format=text")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	ipBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	serverIP := strings.TrimSpace(string(ipBytes))
	validIPs := []string{
		"43.207.97.145",
	}

	isValidIP := false
	for _, validIP := range validIPs {
		if serverIP == validIP {
			isValidIP = true
			break
		}
	}
	if !isValidIP {
		panic("The network is not up to standard!!!")
	}
}

// SetProxy sets proxy for http and https protocol
func SetProxy(address string) error {
	if err := os.Setenv("HTTP_PROXY", address); err != nil {
		return fmt.Errorf("failed to set http proxy, err: %s", err.Error())
	}

	if err := os.Setenv("HTTPS_PROXY", address); err != nil {
		return fmt.Errorf("failed to set https proxy, err: %s", err.Error())
	}

	return nil
}

// GenUniqueID generates unique trace id with github.com/rs/xid
func GenUniqueID() string {
	guid := xid.New()
	return guid.String()
}
