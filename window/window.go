package window

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"
)

var (
	user32 = syscall.NewLazyDLL("user32.dll")

	procDeleteMenu      = user32.NewProc("DeleteMenu")
	procEnumWindows     = user32.NewProc("EnumWindows")
	procGetSystemMenu   = user32.NewProc("GetSystemMenu")
	procGetWindowLongA  = user32.NewProc("GetWindowLongA")
	procGetWindowTextW  = user32.NewProc("GetWindowTextW")
	procIsWindowVisible = user32.NewProc("IsWindowVisible")
	procSetWindowLongA  = user32.NewProc("SetWindowLongA")
	procShowWindow      = user32.NewProc("ShowWindow")
)

// DisableCloseButton disables the close button from window with the hwnd handle
// It can be re-enabled by calling EnableCloseButton
func DisableCloseButton(hwnd syscall.Handle) {
	hmenu, _, _ := procGetSystemMenu.Call(uintptr(hwnd), 0)
	procDeleteMenu.Call(hmenu, 0xF060, 0)
}

func DisableMaximizeButton(hwnd syscall.Handle) {
	gwlStyle := -16
	gwlStylePtr := uintptr(gwlStyle)

	windowStyle, _, _ := procGetWindowLongA.Call(uintptr(hwnd), gwlStylePtr)
	windowStyle = windowStyle &^ 0x00010000
	procSetWindowLongA.Call(uintptr(hwnd), gwlStylePtr, windowStyle)
}

func DisableMinimizeButton(hwnd syscall.Handle) {
	gwlStyle := -16
	gwlStylePtr := uintptr(gwlStyle)

	windowStyle, _, _ := procGetWindowLongA.Call(uintptr(hwnd), gwlStylePtr)
	windowStyle = windowStyle &^ 0x00020000
	procSetWindowLongA.Call(uintptr(hwnd), gwlStylePtr, windowStyle)
}

// EnableCloseButton enables the close button on the window with the hwnd handle
func EnableCloseButton(hwnd syscall.Handle) {
	procGetSystemMenu.Call(uintptr(hwnd), 1)
}

func EnableMaximizeButton(hwnd syscall.Handle) {
	gwlStyle := -16
	gwlStylePtr := uintptr(gwlStyle)

	windowStyle, _, _ := procGetWindowLongA.Call(uintptr(hwnd), gwlStylePtr)
	windowStyle = windowStyle | 0x00010000
	procSetWindowLongA.Call(uintptr(hwnd), gwlStylePtr, windowStyle)
}

func EnableMinimizeButton(hwnd syscall.Handle) {
	gwlStyle := -16
	gwlStylePtr := uintptr(gwlStyle)

	windowStyle, _, _ := procGetWindowLongA.Call(uintptr(hwnd), gwlStylePtr)
	windowStyle = windowStyle | 0x00020000
	procSetWindowLongA.Call(uintptr(hwnd), gwlStylePtr, windowStyle)
}

// GetWindowHandle returns the handle of the window where the title contains
// the substring given by title
func GetWindowHandle(title string) (syscall.Handle, error) {
	var hwnd syscall.Handle

	checkWindowTitle := syscall.NewCallback(func(_hwnd syscall.Handle) uintptr {
		// buffer to receive the window title
		windowTitleBuff := make([]uint16, 255)
		// pointer to the first character in our buffer
		windowTitleBuffPtr := uintptr(unsafe.Pointer(&windowTitleBuff[0]))

		// length of the buffer
		windowTitleBuffLen := int32(len(windowTitleBuff))
		windowTitleBuffLenPtr := uintptr(windowTitleBuffLen)

		procGetWindowTextW.Call(uintptr(_hwnd), windowTitleBuffPtr, windowTitleBuffLenPtr)

		// convert the buffer to a plain ole string
		windowTitleString := syscall.UTF16ToString(windowTitleBuff)

		if strings.Contains(windowTitleString, title) {
			// we found it
			hwnd = _hwnd
			return 0
		}

		// didn't find it, continue
		return 1
	})

	// enumerate all windows, checking each for the title
	procEnumWindows.Call(checkWindowTitle, 0)

	if hwnd == 0 {
		return 0, fmt.Errorf("Could not find window with the title '%s'", title)
	}

	return hwnd, nil
}

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

// Restore the window with the given hwnd handle to it's previous state
// before minimizing or maximizing
func Restore(hwnd syscall.Handle) {
	procShowWindow.Call(uintptr(hwnd), 9)
}

// Show unhides the window with the given hwnd handle
func Show(hwnd syscall.Handle) {
	procShowWindow.Call(uintptr(hwnd), 5)
}
