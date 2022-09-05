# EuroPiGo

Alternate firmware for the [Allen Synthesis EuroPi](https://github.com/Allen-Synthesis/EuroPi) written in Go.

## Getting started

Install Go

[https://go.dev/doc/install](https://go.dev/doc/install)

Install TinyGo

[https://tinygo.org/getting-started/install](https://tinygo.org/getting-started/install)

Install the TinyGo VSCode plugin

[https://marketplace.visualstudio.com/items?itemName=tinygo.vscode-tinygo](https://marketplace.visualstudio.com/items?itemName=tinygo.vscode-tinygo)

## Build the example

Use the `tinygo flash` command while your pico is in BOOTSEL mode to compile the script and copy it to your attached EuroPi pico.

```shell
tinygo flash --target pico examples/diagnostics.go
```

> **_NOTE:_** After the first time you flash your TinyGo program you will no longer need to reboot in BOOTSEL mode to flash your script. Sweet!

## Serial printing

When your EuroPi pico is connected via USB you can view printed serial output using a tool like `minicom`.

For example, a line in your code like:

```go
println(fmt.Sprintf("K1: %2.2f\n", e.K1.ReadVoltage()))
```

> **_NOTE:_** Using the `println` function handles newlines better than `fmt.Println`.

You can launch minicom to view the printed output:

```shell
minicom -b 115200 -o -D /dev/ttyACM0
```

## Developing using picoprobe

Follow the notes on setting up a Picoprobe.

[](https://tinygo.org/docs/reference/microcontrollers/pico/#notes)

Once you have confirmed that setup is working, you can add a default build task to build and flash your project via picoprobe.

```plain
Ctrl + Shift + P > Tasks: Configure Default Build Task
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

```shell
Ctrl + Shift + P > Tasks: Run Build Task`
```

## Debugging using picoprobe

TODO

## Why should I use this?

You probably shouldn't. This is my passion project because I love Go. You should probably be using the official [EuroPi firmware](https://github.com/Allen-Synthesis/EuroPi). But if you are interested in writing a dedicated script for the EuroPi that requires concurrency and faster performance than MicroPython can offer, then maybe this is the right firmware for you!
