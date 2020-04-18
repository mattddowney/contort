Contort
=======

Change the behavior and appearance of running Windows applications from the command line. 
Prevent closing, disable menus, hide an app's gui and more.

What you can do with contort:
* Prevent a user from closing an application
* Disable minimizing and maximizing
* Hide and show windows
* Move windows around on the screen

Installation
------------

Clone the repo: `git clone https://github.com/mattddowney/contort`

Navigate into the cloned folder and run: `go install`

Usage
-----

```
usage: contort [<flags>] <command> [<args> ...]

Modify the behavior and appearance of a window.

Flags:
  --help     Show context-sensitive help (also try --help-long and --help-man).
  --version  Show application version.

Commands:
  help [<command>...]
    Show help.


  disable [<flags>] <window>
    Disable part of a window's GUI.

    --close     Disable a window's close button.
    --maximize  Disable a window's maximize button.
    --menu      Disable a window's menu bar. Once disabled, a window's menu bar
                cannot be re-enabled.
    --minimize  Disable a window's minimize button.

  enable [<flags>] <window>
    Enable part of a window's GUI.

    --close     Enable a window's close button.
    --maximize  Enable a window's maximize button.
    --minimize  Enable a window's minimize button.

  hide <window>
    Hide a window.


  list
    List all window titles.


  maximize <window>
    Maximize a window.


  minimize <window>
    Minimize a window.


  move --x=X --y=Y <window>
    Move a window.

    --x=X  X position
    --y=Y  Y position

  restore <window>
    Restore a window to it's state before minimizing or maximizing.


  show <window>
    Show a window.
```