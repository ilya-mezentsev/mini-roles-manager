# Minimal startup of mini-roles-manager

## Requirements:
* Golang 1.16
* GNU Make 4.3

## Steps:
* Clone repo:
```bash
$ git clone https://github.com/ilya-mezentsev/mini-roles-manager.git && cd mini-roles-manager
```

* Build backend:
```bash
$ make backend-build
```

### If you want to run by Makefile target
* Copy application data file to backend/config/app-data.json
```bash
$ cp /path/to/exported/file.json backend/config/app-data.json
```

* Run:
```bash
$ make backend-run-minimal
```

### If you want to run it somewhere else:
* Copy application binary to desired place:
```bash
$ cp backend/main /path/to/desired/folder
```

* Run:
```bash
$ cd /path/to/desired/folder && ./main -mode minimal -app-data /path/to/exported/file.json -port 8080
```
