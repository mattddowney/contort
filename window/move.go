package window

import "syscall"

// Move the window with the given handle to the position given by x and y
func Move(hwnd syscall.Handle, x int, y int) {
	procSetWindowPos.Call(uintptr(hwnd), 0, uintptr(x), uintptr(y), 0, 0, 1)
}
