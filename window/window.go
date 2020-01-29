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
	procGetWindowLongW  = user32.NewProc("GetWindowLongW")
	procGetWindowTextW  = user32.NewProc("GetWindowTextW")
	procIsWindowVisible = user32.NewProc("IsWindowVisible")
	procSetMenu         = user32.NewProc("SetMenu")
	procSetWindowLongW  = user32.NewProc("SetWindowLongW")
	procSetWindowPos    = user32.NewProc("SetWindowPos")
	procShowWindow      = user32.NewProc("ShowWindow")
)

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

// PrintWindowTitles outputs the title of all Windows to the console
func PrintWindowTitles() {
	printTitle := syscall.NewCallback(func(_hwnd syscall.Handle) uintptr {
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

		fmt.Println(windowTitleString)

		// continue
		return 1
	})

	// enumerate all windows, printing each title
	procEnumWindows.Call(printTitle, 0)
}
