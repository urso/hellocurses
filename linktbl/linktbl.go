package linktbl

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
//
// typedef struct tableEntry {
// 	const char* name;
// 	void* ptr;
// } tableEntry;
//
// typedef struct linkTable {
// 	void* handle;
// 	tableEntry* symbols;
// } linkTable;
//
import "C"
import (
	"fmt"
	"unsafe"
)

func Load(paths []string, in unsafe.Pointer) error {
	tbl := (*C.linkTable)(in)
	handle, libPath := openLib(paths)
	if handle == nil {
		return fmt.Errorf("can not open any of %v", paths)
	}

	tbl.handle = handle
	err := withEntries(tbl, func(entry *C.tableEntry) error {
		ret := linkSymbol(handle, entry.name, entry.ptr)
		if ret < 0 {
			name := C.GoString(entry.name)
			return fmt.Errorf("unknown symbol '%s' in '%v'", name, libPath)
		}
		return nil
	})
	if err != nil {
		Unload(in)
		return err
	}

	return nil
}

func Unload(in unsafe.Pointer) {
	tbl := (*C.linkTable)(in)
	if tbl == nil || tbl.handle == nil {
		return
	}

	withEntries(tbl, func(entry *C.tableEntry) error {
		setPtr(entry.ptr, nil)
		return nil
	})

	C.dlclose(tbl.handle)
	tbl.handle = nil
}

func withEntries(tbl *C.linkTable, fn func(*C.tableEntry) error) error {
	entry := tbl.symbols
	for entry != nil {
		if entry.name == nil {
			break
		}

		err := fn(entry)
		if err != nil {
			return err
		}

		// advance to next entry
		sz := uintptr(C.sizeof_struct_tableEntry)
		entry = (*C.tableEntry)(unsafe.Pointer(uintptr(unsafe.Pointer(entry)) + sz))
	}

	return nil
}

func openLib(paths []string) (unsafe.Pointer, string) {
	for _, path := range paths {
		handle := C.dlopen(C.CString(path), C.int(C.RTLD_LAZY))
		if handle != nil {
			return handle, path
		}
	}
	return nil, ""
}

func linkSymbol(handle unsafe.Pointer, name *C.char, ptr unsafe.Pointer) C.int {
	sym := C.dlsym(handle, name)
	if sym == nil {
		return -1
	}

	setPtr(ptr, sym)
	return 0
}

func setPtr(ptr unsafe.Pointer, sym unsafe.Pointer) {
	to := (*unsafe.Pointer)(ptr)
	*to = sym
}
