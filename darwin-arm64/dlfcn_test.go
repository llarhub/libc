package libc

import (
	"testing"

	"github.com/goplus/lib/c"
)

func TestDlfcn(t *testing.T) {
	handle := Dlopen(nil, RTLD_NOW)
	if handle == nil {
		t.Fatal("dlopen failed")
	}
	symbol := Dlsym(handle, c.Str("malloc"))
	if symbol == nil {
		t.Fatal("dlsym malloc failed")
	}

	var info DlInfo
	if Dladdr(symbol, &info) == 0 {
		t.Fatal("dladdr failed")
	}
	_ = DlopenPreflight(c.Str("/usr/lib/libSystem.B.dylib"))
	if Dlsym(handle, c.Str("__llcppg_missing_symbol__")) != nil || Dlerror() == nil {
		t.Fatal("missing symbol did not report an error")
	}
	if Dlclose(handle) != 0 {
		t.Fatal("dlclose failed")
	}
}
