Contort
=======

Change the behavior and appearance of running Windows applications from the command line. 
Prevent closing, disable menus, hide an app's gui and more.

Examples
--------

Some of what you can do with contort...

### Disable an application's close button

Sometimes you want an application to run, but you don't want your users to be able to close it.

Simply run `contort disable <app_name> --close`.

![contortPreventClose](https://user-images.githubusercontent.com/524968/79677714-804dd380-81b1-11ea-95df-3572ff458331.gif)

### Disable minimizing and maximizing

With contort you can toggle multiple window styles at once.

To prevent minimizing and maximizing, run `contort disable <app_name> --minimize --maximize`.

![contortDisableMinMax](https://user-images.githubusercontent.com/524968/79677723-85ab1e00-81b1-11ea-809e-405bb5fa3ed8.gif)

### Disable an application's menubar

Quite often key settings are located in an application's menu.

Lock the menu down by disabling it completely with `contort disable <app_name> --menu`.

![contortDisableMenu](https://user-images.githubusercontent.com/524968/79677778-ecc8d280-81b1-11ea-9202-be9266ca25a7.gif)

### Hide and show windows

Maybe you want your GUI application to run in the background without your user's knowledge.

This is possible by running `contort hide <app_name>`.

![contortHideShow](https://user-images.githubusercontent.com/524968/79677779-f4887700-81b1-11ea-9677-29d86ac31252.gif)

Installation
------------

Clone the repo: `git clone https://github.com/mattddowney/contort`

Navigate into the cloned folder and run: `go install`

Usage
-----

```
usage: contort [<flags>] <command> [<args> ...]

Modify the behavior and appearance of a window.
https://github.com/mattddowney/contort

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
    --titlebar  Disable a window's titlebar.

  enable [<flags>] <window>
    Enable part of a window's GUI.

    --close     Enable a window's close button.
    --maximize  Enable a window's maximize button.
    --minimize  Enable a window's minimize button.
    --titlebar  Enable a window's titlebar.

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
