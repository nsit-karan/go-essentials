# go-essentials
go quick rampup

# module1 notes
- function overloading is not allowed in go (method overloading is allowed though)
  - implies that i cant have 2 functions with same name and diff return types
- go follows pass by value, i.e, a copy is passed to the function


# module1 build
- Till the code was organized into folders but was not using go modules (go.mod), simpl 'go build' works - go build -o bin/module1 module1/*.go
- now, with the code organized as go modules, the way to run the code is:
  - navigate to module1
  - go run main.go
