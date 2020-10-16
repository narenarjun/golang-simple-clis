# Calc CLI

```bash
$>> go run main.go --help
```

```bash
A simple calc cli with cobra and golang which does additon and multiplication.

Usage:
  mycalc [flags]
  mycalc [command]

Available Commands:
  add         Adds two or more number
  even        adds even numbers only from the given
  help        Help about any command
  multiply    multiply two or more numbers
  odd         adds odd numbers only from the given

Flags:
      --config string   config file (default is $HOME/.mycalc.yaml)
  -h, --help            help for mycalc
  -t, --toggle          Help message for toggle

Use "mycalc [command] --help" for more information about a command.
```
