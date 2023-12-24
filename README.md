# mer

CLI to convert currency.

[![CI](https://github.com/winebarrel/mer/actions/workflows/ci.yml/badge.svg)](https://github.com/winebarrel/mer/actions/workflows/ci.yml)

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

$ echo 1.53 | mer usd eur
1.36185

$ mer -c usd jpy 1000
142,401
```

## Installation

```sh
brew install winebarrel/mer/mer
```
