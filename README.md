# mer

CLI to convert currency using Yahoo! finance API.

[![CI](https://github.com/winebarrel/mer/actions/workflows/ci.yml/badge.svg)](https://github.com/winebarrel/mer/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/winebarrel/mer.svg)](https://pkg.go.dev/github.com/winebarrel/mer)


## Usage

```
Usage: mer <from> <to> [<src>]

Arguments:
  <from>     Exchange source currency code.
  <to>       Exchange destination currency code.
  [<src>]    Exchange source.

Flags:
  -h, --help       Show help.
  -c, --comma      Add comma to a number.
      --version
```

```sh
$ mer usd jpy 3
427.203

mer jpy eur 10
0.064

$ echo 1.53 | mer usd eur
1.36185

$ mer usd jpy 3000
21

$ mer usd jpy 3,000 # commas are ignored
21

$ mer usd jpy 1000
142401

$ mer -c usd jpy 1000
142,401
```

## Installation

```sh
brew install winebarrel/mer/mer
```
