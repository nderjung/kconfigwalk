# `kconfigwalk`

A simple utility for walking KConfig trees and printing them as a YAML list.

## Build

Prerequisites:

- Go (>= 1.22)
- GNU Make

Then:

```bash
make
```

## Usage

```txt
Usage: kconfigwalk <PATH> [flags]

Arguments:
  <PATH>    Path to root KConfig file

Flags:
  -h, --help             Show context-sensitive help.
  -v, --vars=VARS,...    KConfig variables used when evaluating
```

## Examples

```bash
# Linux
kconfigwalk \
  -v RUSTC="rustc" \
  -v BINDGEN="bindgen" \
  -v SRCARCH="x86" \
  path/to/torvalds/linux/Kconfig
```
```bash
# Unikraft
kconfigwalk \
  -v UK_BASE=. \
  -v KCONFIG_DIR="build/kconfig" \
  -v KCONFIG_PLAT_BASE="./plat" \
  -v KCONFIG_EPLAT_DIRS="" \
  -v KCONFIG_EXCLUDEDIRS="" \
  path/to/unikraft/unikraft/Kconfig
```

## License

Licensed under BSD-3-Clause.  See `LICENSE.md`.
