/**
 * @Author: zhangdi
 * @File: read_test
 * @Version: 1.0.0
 * @Date: 2022/8/27 15:13
 * @Software : GoLand
 */
package configs

import (
	"testing"
)

func TestRead(t *testing.T) {
	println(Conf.Wechat.Appid)
}

func TestJson(t *testing.T) {
	jsonStr := `{"access_token":"60_uH3yfTp6-8iVA4az3Zh6bGTxo9jyA1IMwGKcrOQMFcnw6pep4SX5entqXEfQ5xyg4A2GENw-xVtUYieDwZDWVVSS0tQuKbsTG0kThs7mgNuBfJHuWfFuc26_rvXVrcAtZsp3__4Dy8QY6QVpTCYiADAJSO","expires_in":7200}`
	m := make(map[string]string)
	json.Unmarshal([]byte(jsonStr), m)
}
