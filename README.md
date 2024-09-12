
# Walker

**Walker** is a Go-based tool designed to help identify secrets within binary files, specifically targeting desktop application penetration testing.

## Installation

### Linux
```bash
git clone https://github.com/DEMON1A/Walker
go build cmd/walker/main.go
./main -h
```

### Windows
```bat
git clone https://github.com/DEMON1A/Walker
go build cmd\walker\main.go
main.exe -h
```

## Usage

### Help Menu
```bash
Usage of C:\Users\admin\temp\main.exe:
  -dir string
        Directory path to list files from (default ".")
  -exclude string
        Rule IDs to exclude from the scan
  -max int
        Maximum length of strings to print (default 4)
  -regex string
        Regex pattern to search for within files
  -scan
        Scan all identified strings using a regex dataset
  -search string
        Search for a specific string within files
  -sensitive
        Toggle case-sensitive or case-insensitive search mode
```

### Examples

#### Scan a Folder
```bash
main.exe -dir . -scan -exclude generic-api-key,http-https-url
```

#### Search for a Specific String
```bash
main.exe -dir . -search "api-key"
```

#### Search Using Regex
```bash
main.exe -dir . -regex ".*\_key"
```

## Credits
- Special thanks to [gitleaks](https://github.com/gitleaks/gitleaks) for their TOML configuration used in Walker to detect secrets.
