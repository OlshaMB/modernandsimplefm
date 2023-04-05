# modernandsimplefm
Do you know nginx file server list?<br>
This project is exactly that but more extensible and with more modern ui
## How to run
`./masfm`
## Config file
location: ./config.toml
```toml
dir = "/path/to/folder/to/host" # Should be absolute
location = "localhost:8080"
```
## License
The project is licensed under **Apache 2.0** except everything in `pkg/view` folder is licensed under **Zlib**
## Build
Deps:
- golang (1.19)
- GNU make
- unix coreutils

`make bundle`