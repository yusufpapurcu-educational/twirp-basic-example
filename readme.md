# Basic example with twirp
Intro : https://twitchtv.github.io/twirp/docs/example.html

## How to run?
```
$ protoc --twirp_out=. --go_out=. rpc/haberdasher/service.proto
$ go run ./cmd/server/main.go
```

Run this in another window and sure server is up.
```
$ go run ./cmd/client/main.go
I have a nice new hat: inches:12 color:"black" name:"bowler"
```
