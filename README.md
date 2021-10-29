# scanner

scanner is a tiny TCP scanner library written in Go. Similar to nmap

## Installation

```bash
VERSION=0.0.1
OS=darwin
ARCH=amd64

wget -qO- https://github.com/oschvr/scanner/releases/download/$VERSION/scanner-$VERSION-$OS-$ARCH.tar.gz | tar -xzvf - -C /usr/local/bin

mv /usr/local/bin/scanner-$VERSION-$OS-$ARCH /usr/local/bin/scanner
```

## Usage

```bash
➜ scanner "scanme.nmap.org"

  ___    ___    __ _   _ __    _ __     ___   _ __
 / __|  / __|  / _` | | '_ \  | '_ \   / _ \ | '__|
 \__ \ | (__  | (_| | | | | | | | | | |  __/ | |
 |___/  \___|  \__,_| |_| |_| |_| |_|  \___| |_|

----------------
Warning 🌕: Scanning... !
Connection to scanme.nmap.org succesful on port 22
Connection to scanme.nmap.org succesful on port 80
Connection to scanme.nmap.org succesful on port 9929
Connection to scanme.nmap.org succesful on port 31337
Success ✅: Scan complete
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[AGPLv3](https://choosealicense.com/licenses/agpl-3.0/)
