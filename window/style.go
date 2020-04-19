// styles in this file referenced from here:
// https://docs.microsoft.com/en-us/windows/desktop/winmsg/window-styles

package window

import "syscall"

const (
	maximizeButton = 0x00010000 // WS_MAXIMIZEBOX
	minimizeButton = 0x00020000 // WS_MINIMIZEBOX
	titlebar       = 0x00C00000 // WS_CAPTION
)

func addStyle(hwnd syscall.Handle, ws int) {
	gwlStyle := -16
	gwlStylePtr := uintptr(gwlStyle)

	windowStyle, _, _ := procGetWindowLongW.Call(uintptr(hwnd), gwlStylePtr)
	windowStyle = windowStyle | uintptr(ws)
	procSetWindowLongW.Call(uintptr(hwnd), gwlStylePtr, windowStyle)
}

func removeStyle(hwnd syscall.Handle, ws int) {
	gwlStyle := -16
	gwlStylePtr := uintptr(gwlStyle)

	windowStyle, _, _ := procGetWindowLongW.Call(uintptr(hwnd), gwlStylePtr)
	windowStyle = windowStyle &^ uintptr(ws)
	procSetWindowLongW.Call(uintptr(hwnd), gwlStylePtr, windowStyle)
}

// DisableMaximizeButton disables the maximize button on the window with the hwnd handle
// It can be re-enabled by calling EnableMaximizeButton
func DisableMaximizeButton(hwnd syscall.Handle) {
	removeStyle(hwnd, maximizeButton)
}

// DisableMinimizeButton disables the minimize button on the window with the hwnd handle
// It can be re-enabled by calling EnableMinimizeButton
func DisableMinimizeButton(hwnd syscall.Handle) {
	removeStyle(hwnd, minimizeButton)
}

// DisableTitlebar disables the titlebar on the window with the hwnd handle
// It can be re-enabled by calling EnableTitlebar
func DisableTitlebar(hwnd syscall.Handle) {
	removeStyle(hwnd, titlebar)
}

// EnableMaximizeButton enables the maximize button on the window with the hwnd handle
func EnableMaximizeButton(hwnd syscall.Handle) {
	addStyle(hwnd, maximizeButton)
}

// EnableMinimizeButton enables the minimize button on the window with the hwnd handle
func EnableMinimizeButton(hwnd syscall.Handle) {
	addStyle(hwnd, minimizeButton)
}

// EnableTitlebar enables the titlebar on the window with the hwnd handle
func EnableTitlebar(hwnd syscall.Handle) {
	addStyle(hwnd, titlebar)
}
