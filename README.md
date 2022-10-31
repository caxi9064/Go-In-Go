# Go-in-Go
The oldest board game Go, written in Go. 

## How to run
First clone the repo. 

While you are in the root of the repo, run 
```
    go build ./cmd/main
```

You'll find an executable in the root directory. You can run this on mac by running
```
    ./main
```

Alternatively, you can call it something else by using the -o flag in build. If you are on windows youll have to run the executable differently.

Some notes about imports/ packages in go:

If you want to add some functionality to the program, use a package. go.sum and go.mod being in the root directory means that all the files within the subdirectories get access to the module and its requirements. 

Say you want to add a directory named foo and call its functionality in main:
```
.
├── cmd
│   └── main <- calling file
└── pkg
    ├── GUI
    ├── GameLogic
    └── foo <- new file!
```    
In foo, import all the necessary third party libraries you need. In main, you can import foo in main with the following line:
import "github.com/johnowagon/Go-in-Go/pkg/foo"

NOTE: exported functions in foo start with capital letters:
```go
  func Hello(){
    fmt.Println("Hello");
  }
```
