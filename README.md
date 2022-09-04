# EuroPiGo

Alternate firmware for the [Allen Synthesis EuroPi](https://github.com/Allen-Synthesis/EuroPi) written in Go.

# Getting started

Install Go

https://go.dev/doc/install

Install TinyGo

https://tinygo.org/getting-started/install/

Install the TinyGo VSCode plugin

https://marketplace.visualstudio.com/items?itemName=tinygo.vscode-tinygo


# Build the example

Use the `tinygo flash` command while your pico is in BOOTSEL mode to compile the script and copy it to your attached EuroPi pico.

```
tinygo flash --target pico examples/diagnostics.go
```

# Developing using picoprobe

Follow the notes on setting up a Picoprobe.

https://tinygo.org/docs/reference/microcontrollers/pico/#notes

Once you have confirmed that setup is working, you can add a default build task to build and flash your project via picoprobe.

```
Ctrl + Shift + P > Tasks: Configure Default Build Task`
```

Use the following example task configurations to add the automated build and flash tasks. Set "tinygo flash task" as your default build command:

```json
{
    "version": "2.0.0",
    "tasks": [
        {
            "label": "tinygo build task",
            "type": "shell",
            "command": "tinygo build -o out.elf -target pico -size short -opt 1 ${workspaceRoot}",
            "problemMatcher": [],
            "group": {
                "kind": "build",
                "isDefault": true
            }
        },
        {
            "label": "tinygo flash task",
            "type": "shell",
            "command": "openocd -f interface/picoprobe.cfg -f target/rp2040.cfg -c \"program out.elf verify reset exit\"",
            "problemMatcher": [],
            "group": {
                "kind": "build",
                "isDefault": true
            },
            "dependsOn": ["tinygo build task"]
        }
    ]
}
```

Now you can build your project using `Ctrl + Shift + B` or search for the command:

```
Ctrl + Shift + P > Tasks: Run Build Task`
```

# Why should I use this?

You probably shouldn't. This is my passion project because I love Go. You should probably be using the official [EuroPi firmware](https://github.com/Allen-Synthesis/EuroPi). But if you are interested in writing a dedicated script for the EuroPi that requires concurrency and faster performance than MicroPython can offer, then maybe this is the right firmware for you!