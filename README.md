# Go-in-Go
The oldest board game Go, written in Go. 

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
