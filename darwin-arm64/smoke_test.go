package libc

import (
	"testing"

	"github.com/goplus/lib/c"
)

func TestGeneratedLibcSmoke(t *testing.T) {
	if got := Strlen(c.Str("hello")); got != 5 {
		t.Fatalf("Strlen = %d, want 5", got)
	}
	if got := Strnlen(c.Str("hello"), 3); got != 3 {
		t.Fatalf("Strnlen = %d, want 3", got)
	}
	if got := Sqrt(9); got != 3 {
		t.Fatalf("Sqrt = %v, want 3", got)
	}
	if Localeconv() == nil {
		t.Fatal("Localeconv returned nil")
	}
	if got := Time(nil); got <= 0 {
		t.Fatalf("Time = %d, want positive value", got)
	}
	if ENOENT != 2 || EX_OK != 0 || EX_USAGE != 64 {
		t.Fatal("unexpected libc constants")
	}
	handle := Dlopen(nil, RTLD_NOW)
	if handle == nil {
		t.Fatal("Dlopen returned nil")
	}
	if Dlsym(handle, c.Str("strlen")) == nil {
		t.Fatal("Dlsym could not find strlen")
	}
	if rc := Dlclose(handle); rc != 0 {
		t.Fatalf("Dlclose = %d", rc)
	}
}
