package window

import (
	"syscall"
	"unsafe"
)

// Hide the window given by hwnd from the user
func Hide(hwnd syscall.Handle) {
	procShowWindow.Call(uintptr(hwnd), 0)
}

// IsVisible returns true if the window is visible, false if it is not
// Hidden windows can be made visible by calling Show
func IsVisible(hwnd syscall.Handle) bool {
	ret, _, _ := procIsWindowVisible.Call(uintptr(hwnd))

	if ret == 1 {
		return true
	}

	return false
}

// Maximize the window with the given hwnd handle
func Maximize(hwnd syscall.Handle) {
	procShowWindow.Call(uintptr(hwnd), 3)
}

// Minimize the window with the given hwnd handle
func Minimize(hwnd syscall.Handle) {
	procShowWindow.Call(uintptr(hwnd), 6)
}

// RemoveMenu removes the menu from the window with the given hwnd window
func RemoveMenu(hwnd syscall.Handle) {
	procSetMenu.Call(uintptr(hwnd), uintptr(unsafe.Pointer(nil)))
}

// Restore the window with the given hwnd handle to it's previous state
// before minimizing or maximizing
func Restore(hwnd syscall.Handle) {
	procShowWindow.Call(uintptr(hwnd), 9)
}

// Show unhides the window with the given hwnd handle
func Show(hwnd syscall.Handle) {
	procShowWindow.Call(uintptr(hwnd), 5)
}
