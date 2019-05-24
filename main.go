package main

import (
	"os"

	"github.com/mattddowney/warpwin/window"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("warpwin", "Modify the behavior and appearance of a window.")

	disableCommand         = app.Command("disable", "Disable part of a window's GUI.")
	disableCommandClose    = disableCommand.Flag("close", "Disable a window's close button.").Bool()
	disableCommandMaximize = disableCommand.Flag("maximize", "Disable a window's maximize button.").Bool()
	disableCommandMenu     = disableCommand.Flag("menu", "Disable a window's menu bar. Once disabled, a window's menu bar cannot be re-enabled.").Bool()
	disableCommandMinimize = disableCommand.Flag("minimize", "Disable a window's minimize button.").Bool()
	disableCommandWindow   = disableCommand.Arg("window", "Window Title").Required().String()

	enableCommand         = app.Command("enable", "Enable part of a window's GUI.")
	enableCommandClose    = enableCommand.Flag("close", "Enable a window's close button.").Bool()
	enableCommandMaximize = enableCommand.Flag("maximize", "Enable a window's maximize button.").Bool()
	enableCommandMinimize = enableCommand.Flag("minimize", "Enable a window's minimize button.").Bool()
	enableCommandWindow   = enableCommand.Arg("window", "Window Title").Required().String()

	hideCommand       = app.Command("hide", "Hide a window.")
	hideCommandWindow = hideCommand.Arg("window", "Window Title").Required().String()

	maximizeCommand       = app.Command("maximize", "Maximize a window.")
	maximizeCommandWindow = maximizeCommand.Arg("window", "Window Title").Required().String()

	minimizeCommand       = app.Command("minimize", "Minimize a window.")
	minimizeCommandWindow = minimizeCommand.Arg("window", "Window Title").Required().String()

	restoreCommand       = app.Command("restore", "Restore a window to it's state before minimizing or maximizing.")
	restoreCommandWindow = restoreCommand.Arg("window", "Window Title").Required().String()

	showCommand       = app.Command("show", "Show a window.")
	showCommandWindow = showCommand.Arg("window", "Window Title").Required().String()
)

func main() {
	app.Version("0.1.0")
	app.UsageTemplate(kingpin.LongHelpTemplate)

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	// Disable UI Items
	case disableCommand.FullCommand():
		hwnd, err := window.GetWindowHandle(*disableCommandWindow)
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}

		if *disableCommandClose {
			window.DisableCloseButton(hwnd)
		}
		if *disableCommandMaximize {
			window.DisableMaximizeButton(hwnd)
		}
		if *disableCommandMenu {
			window.RemoveMenu(hwnd)
		}
		if *disableCommandMinimize {
			window.DisableMinimizeButton(hwnd)
		}

	// Enable UI Items
	case enableCommand.FullCommand():
		hwnd, err := window.GetWindowHandle(*enableCommandWindow)
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}

		if *enableCommandClose {
			window.EnableCloseButton(hwnd)
		}
		if *enableCommandMaximize {
			window.EnableMaximizeButton(hwnd)
		}
		if *enableCommandMinimize {
			window.EnableMinimizeButton(hwnd)
		}

	// Hide Window
	case hideCommand.FullCommand():
		hwnd, err := window.GetWindowHandle(*hideCommandWindow)
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}

		window.Hide(hwnd)

	// Maximize Window
	case maximizeCommand.FullCommand():
		hwnd, err := window.GetWindowHandle(*maximizeCommandWindow)
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}

		window.Maximize(hwnd)

	// Minimize Window
	case minimizeCommand.FullCommand():
		hwnd, err := window.GetWindowHandle(*minimizeCommandWindow)
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}

		window.Minimize(hwnd)

	// Restore Window
	case restoreCommand.FullCommand():
		hwnd, err := window.GetWindowHandle(*showCommandWindow)
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}

		window.Restore(hwnd)

	// Show Window
	case showCommand.FullCommand():
		hwnd, err := window.GetWindowHandle(*showCommandWindow)
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}

		window.Show(hwnd)
	}
}
