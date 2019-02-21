# Taktgeber für die kaffeesimulation

## How to connect

```
wscat -c ws://localhost:8080/clock
```

## How to Run

in Mac/Linux

```
$ ./src/timebase/timebase
```

für kompilieren Windows => good luck

```
$ GOOS=windows GOARCH=386 go build -o hello.exe ./src/timebase/timebase.go
```
