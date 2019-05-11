package window

import "syscall"

// DisableCloseButton disables the close button on the window with the hwnd handle
// It can be re-enabled by calling EnableCloseButton
func DisableCloseButton(hwnd syscall.Handle) {
	hmenu, _, _ := procGetSystemMenu.Call(uintptr(hwnd), 0)
	procDeleteMenu.Call(hmenu, 0xF060, 0)
}

// EnableCloseButton enables the close button on the window with the hwnd handle
func EnableCloseButton(hwnd syscall.Handle) {
	procGetSystemMenu.Call(uintptr(hwnd), 1)
}
