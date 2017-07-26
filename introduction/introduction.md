BASICS
-
1. Project infrastructure
   - GOPATH structure
   - Project structure
   - File structure
   - main function
1. Language
   - tour.golang.org / play.golang.org
   - [Dos & Don'ts](https://docs.google.com/document/d/152oQ24u1BMc0t2NJppEE5Cs6hiXtkGsa2p9zzUf4GgE/)
   - multiple return parameters
   - named return parameters
   - type conversion/assertion
   - for, range, switch
   
IDIOMATIC
-
1. Formatting
   - go fmt
   - happy path
   - minimal abstractions
   - single responsibility functions
1. Error Handling
   - if err != nil
   - early exit
   - panic/recover (+in goroutines)
   - fmt.Errorf with additional information
1. Testing
   - examples
   - unit testing
   - table testing
   - benchmarks
1. Interfaces
   - return structs, expect interfaces
   - minimal interface
   - interface composition
1. Objects & Functions
   - struct
   - function receiver
   - pass by value/reference

ADVANCED
-
1. Context
   - Cancel+Timeout
   - value storage + typed keys
1. Concurrency
   - go func
   - WaitGroup
   - Channels (read/write)
   - select
1. HTTP Server
   - endpoints
   - handlers ([stdlib routing](http://blog.merovius.de/2017/06/18/how-not-to-use-an-http-router.html))
   - middleware
   - HTTP server testing
