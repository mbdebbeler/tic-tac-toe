# TDD TicTacToe in Go
I'm learning Go. This is a command line TicTacToe game written in Go. I am trying my best to TDD my way through. 

## To Run the Tests
Get into your terminal and check whether Go is installed: 
```which go```

If not, [install Go](https://golang.org/doc/install). 

Then, clone the repository and `cd` in.
```cd tic-tac-toe
go build
go test .```

Running `go test .` instead of `go test` silences test terminal output during tests. 

## To Play
From inside the directory, 
```
go build
./ticTacToe
```

If you have exported your $GOROOT, you can run:
```go install```
and then run the game from any directory with the command ticTacToe.