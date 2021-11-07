# LearningGO
Learning GO programming.
  
> Download GO : [https://golang.org/dl/](https://golang.org/dl/)  
  
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
    
## GO Tools  
1. **hey** - Load tests HTTP servers.  
> go install github.com/rakyll/hey@latest  
2. **goimports** - Enhanced version of `go fmt` which also cleans up import statements.  
> go install golang.org/x/tools/cmd/goimport@latest.  
> go imports -l -w .  
>> "-l" = prints the file with incorrect formatting.  
>> "-w"= modify the files in place.  
>> "."  = specifies the file to be scanned which in this case is everything in the current directory and all of its subdirectory.  
     

  