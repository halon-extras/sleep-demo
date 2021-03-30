package main

// #cgo CFLAGS: -I/opt/halon/include
// #cgo LDFLAGS: -Wl,--unresolved-symbols=ignore-all
// #include <HalonMTA.h>
import "C"
import "unsafe"
import "time"

func main() {}

//export Halon_version
func Halon_version() C.int {
	return C.HALONMTA_PLUGIN_VERSION;
}

func sleepTask(hhc *C.HalonHSLContext, ret *C.HalonHSLValue, sec float64) {
	time.Sleep(time.Duration(sec * 1000.0) * time.Millisecond);
	C.HalonMTA_hsl_schedule(hhc);
}

//export sleep
func sleep(hhc *C.HalonHSLContext, arg *C.HalonHSLArguments, ret *C.HalonHSLValue) {
	var a C.double = 0;
	var x = C.HalonMTA_hsl_argument_get(arg, 0);
	if (x != nil) {
		C.HalonMTA_hsl_value_get(x, C.HALONMTA_HSL_TYPE_NUMBER, unsafe.Pointer(&a), nil);
	}
	go sleepTask(hhc, ret, float64(a));
	C.HalonMTA_hsl_suspend_return(hhc);
}

//export Halon_hsl_register
func Halon_hsl_register(hhrc *C.HalonHSLRegisterContext) C.bool {
	C.HalonMTA_hsl_register_function(hhrc, C.CString("sleep"), nil);
	return true;
}