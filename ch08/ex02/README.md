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

### Reference(mainly japanese)
 - [FTPのプロトコルについて前編](http://www.atmarkit.co.jp/ait/articles/0107/17/news002.html)
 - [FTPのプロトコルについて後編](http://www.atmarkit.co.jp/ait/articles/0108/03/news001.html)