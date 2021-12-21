# LearningGO
Learning GO programming.
  
> Download GO : [https://golang.org/dl/](https://golang.org/dl/)  
> Effective GO Documentation : https://golang.org/doc/effective_go#introduction  
> How Do You Structure Your Go Apps : https://oreil.ly/0zHY4  
> FOSS - Open Source Licenses : https://oreil.ly/KVlrd  
> Profiling GO : https://oreil.ly/HHe9c  
> Type Parameters-Draft Design : https://oreil.ly/POhSg  
  
VSCode Extension : 
- [GO](golang.go)
- Open Command Palette and search GO Tools. Install all tools listed . 
  
## GO Commands  
1. `go version` - To check GO version.  
2. `go env` - To check GO environment config.  
3. `go run` - Compile code into binary in temporary directory. Once execution completed, will delete the binary.  
4. `go build` - Build binary for later use. Creates executable .exe file.  
5. `go install` - Takes an argument(source code repository) followed by an @ plus the version of the tool. It then downloads, compiles, and install into *$GOPATH/bin* direrctory.  
6. `go fmt` - Automatically reformat codes to match the standard format.  
7. `go mod init <repo>` - Modules initialization  
8. `go doc PACKAGE_NAME.IDENTIFIER_NAME` - To display documentation for specific identifier in the package.  
9. `go list -m -versions MODULE_PATH` - To check versions of the module available.  
10. `go get MODULE_PATH@VERSION` - To upgrade module version.  
11. `go get -u=patch MODULE_PATH` - To upgrade patch version for the same module version.  
12. `go mod tidy` - To remove unused versions.  
13. `go mod vendor` - To create vendor directory contains all of modules dependencies.  
14. `go test` - To run test and generate report.  
15. `go test -v -cover -coverprofile=c.out` - Calculates coverage information and includes summary in the test output and save into a file.  
16. `go tool cover -html=c.out` - Generates HTML representation of the source code.  
17. `go test -bench=. -benchmem` - Run all benchmarks and includes memory allocation information.  
18. `go test -race` - Concurrency race checker.  
  
## GO Tools  
1. **hey** - Load tests HTTP servers.  
> go install github.com/rakyll/hey@latest  
2. **goimports** - Enhanced version of `go fmt` which also cleans up import statements.  
> go install golang.org/x/tools/cmd/goimport@latest.  
> go imports -l -w .  
>> "-l" = prints the file with incorrect formatting.  
>> "-w"= modify the files in place.  
>> "."  = specifies the file to be scanned which in this case is everything in the current directory and all of its subdirectory.  
3. **golint** & **go vet** - To ensure the code follows sytle guidelines.  
> go install golang.org/x/lint/golint@latest.  
> golint ./...  
> go vet ./...  
4. **golangci-lint** - Combined tools for `golint`, `go vet` and set of other code quality tools.  
> go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest.  
> golangci-lint run.  
>
> To configure, go to root project and find .golangci.yml.  
>> Documentation: https://oreil.ly/vufj1  
4. **shadow** - To detect shadowed variables. Include shadow in vet task in Makefile.  
> go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
> shadow ./...  
5. **wire** - A dependency injection helper.
> go get github.com/google/wire/cmd/wire
>> Documentation: https://github.com/google/wire

  