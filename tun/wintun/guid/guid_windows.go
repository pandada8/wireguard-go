/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package guid

import (
	"fmt"
	"syscall"

	"golang.org/x/sys/windows"
)

//sys	clsidFromString(lpsz *uint16, pclsid *windows.GUID) (err error) [failretval!=0] = ole32.CLSIDFromString

//
// FromString parses "{XXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX}" string to GUID.
//
func FromString(str string) (*windows.GUID, error) {
	strUTF16, err := syscall.UTF16PtrFromString(str)
	if err != nil {
		return nil, err
	}
	guid := &windows.GUID{}
	err = clsidFromString(strUTF16, guid)
	if err != nil {
		return nil, err
	}
	return guid, nil
}

//
// ToString function converts GUID to string
// "{XXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX}".
//
// The resulting string is uppercase.
//
func ToString(guid *windows.GUID) string {
	return fmt.Sprintf("{%08X-%04X-%04X-%04X-%012X}", guid.Data1, guid.Data2, guid.Data3, guid.Data4[:2], guid.Data4[2:])
}
