## Junk FTP (Go)

### Default Port
8080

### Implemented command.

- ls
- cd
- send
- get
- exit


### build
```bash
$ go build ftp.go 
```

### Start
```bash
$ ./ftp&
```

### Connect
```bash
$ ftp
ftp> open
(to) 
usage: open host-name [port]
ftp> open localhost 8080
```