package hellsgopher

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/Binject/debug/pe"
)

const NTDLL_Path = "C:\\Windows\\System32\\ntdll.dll"

var hookCheck = []byte{0x4c, 0x8b, 0xd1, 0xb8} // Define hooked bytes to look for

type mayBeHookedError struct { // Define custom error for hooked functions
	Foundbytes []byte
}

// detect whether any functions in a dll are hooked by AV/EDR
// defaults to NTDLL, leave blank for default
// otherwise provide a full path to dll
func CheckHooks(path string) ([]string, error) {
	// check if path is blank
	if path == "" {
		path = NTDLL_Path
	}

	// check if given path is a valid file
	if !DoesFileExist(path) {
		return nil, ErrFile_Not_Found
	}

	// check if file is a dll
	if filepath.Ext(path) != "dll" {
		return nil, ErrFile_Not_DLL
	}

	// detect hooked functions
	var hookedFuncs []string

	dllFile, err := pe.Open(path)
	if err != nil {
		return nil, err
	}
	defer dllFile.Close()

	exportedFuncs, err := dllFile.Exports()
	if err != nil {
		return nil, err
	}

	for _, exportedFunc := range exportedFuncs {
		offset := rvaToOffset(dllFile, exportedFunc.VirtualAddress)

		bytes, err := dllFile.Bytes()
		if err != nil {
			return hookedFuncs, err
		}

		buffer := bytes[offset : offset+10]

		_, err = checkBytes(buffer)
		if err != nil {
			return hookedFuncs, err
		}

		var hookErr mayBeHookedError

		if len(exportedFunc.Name) > 3 {
			if exportedFunc.Name[0:2] == "Nt" || exportedFunc.Name[0:2] == "Zw" {
				if errors.As(err, &hookErr) {
					hookedFuncs = append(hookedFuncs, exportedFunc.Name)
				}
			}
		}
	}

	return hookedFuncs, nil
}

/*
Auxilliary Funtions
*/

func checkBytes(b []byte) (uint16, error) {
	if !bytes.HasPrefix(b, hookCheck) { // Check syscall bytes
		return 0, mayBeHookedError{Foundbytes: b}
	}

	return binary.LittleEndian.Uint16(b[4:8]), nil
}

func rvaToOffset(pefile *pe.File, rva uint32) uint32 {
	for _, hdr := range pefile.Sections {
		baseoffset := uint64(rva)
		if baseoffset > uint64(hdr.VirtualAddress) &&
			baseoffset < uint64(hdr.VirtualAddress+hdr.VirtualSize) {
			return rva - hdr.VirtualAddress + hdr.Offset
		}
	}
	return rva
}

func (e mayBeHookedError) Error() string {
	return fmt.Sprintf("may be hooked: wanted %x got %x", hookCheck, e.Foundbytes)
}
