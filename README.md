# GO-Lang-Practise

## Make basic CLI

Let try to use this CLI with following commands:

```bash
# Default CLI:
> go run ./cli

A basic cli with Cobra!

Usage:
  cli [command]

Available Commands:
  help        Help about any command
  whereami    Display absolute path of current workfing directory.

Flags:
  -h, --help   help for cli

Use "cli [command] --help" for more information about a command.
```

```bash
> go run ./cli whereami --name kimyvgy

Hello kimyvgy!
You are at:  /home/kimnguyen/kimyvgy-forks/GO-Lang-Practise
```
