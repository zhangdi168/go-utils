// Package errorhandler Package common /**
package tools

// GetErrorRecover
// @Description: 捕获异常
func GetErrorRecover() {
	if err := recover(); err != nil {
		println(err)
		logWrite(err)
	}
}
