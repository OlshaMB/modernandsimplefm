# modernandsimplefm
TODO: description
## How to run
`./masfm`
## Config file
location: ./config.toml
```toml
dir = "/path/to/folder/to/host" # Should be absolute
location = "localhost:8080"
```
## Build
Deps:
- golang (1.19)
- GNU make
- unix coreutils

`make bundle`