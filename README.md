# golang-basics
practice golang basics code

# to generate the go.mod file where you can maintain your dependencies
go mod init hello

# how to build for different platforms
in terminal if you'll put go env - it will give you goos - which is considered as your environment
and go will create executable for that operating system

so if you wanna change that

GOOS="linux" go build
or
GOOS="windows" go build
or
"GOOS"="darwin" go build 
- for mac

//case sensitive so use it in capitals 


# Memory management in golang
Allocation and deallocation in golang happens automatically

you can use 2 methods

1. New()  
when New() is called only memory is allocated but not initialized
you will get memry address but zeroed storage

2. Make()
when Make() is called memory is allocated also initialised
you will get memory address and non-zeroed storage

Note-  GOGC from go env you can use to set the garbage collector percentage

# check more with runtiem package to know more

example= to check the number of CPU
call runtime.NumCPU()

