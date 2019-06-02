package window

import "syscall"

// Move the window with the given handle to the position given by x and y
func Move(hwnd syscall.Handle, x int, y int) {
	procMoveWindow.Call(uintptr(hwnd), uintptr(x), uintptr(y), 100, 100)
}
