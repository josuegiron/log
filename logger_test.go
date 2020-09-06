package log

import (
	"testing"
)

func TestLogger_Output(t *testing.T) {
	func() {
		std.SetNewFormat(NewFormaterConfig(LstdDevFlags, LstdProdFlags|Linfo, "test"))
		std.DevelopmentMode()
		stackTrace := GetStackTrace(2)
		std.Output(1, infoPriority, "Hello world!", nil, stackTrace)
	}()
}