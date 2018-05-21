# Simple CLI App to Query GitHub users

The purpose of this application was understanding how to build Go applications with appropriate structure and linking two files together.

You'll have to add this directory to your Go workspace before you can run it.

I used a Go library for reading user flags

first, run this command to get it:
```$ go get github.com/ogier/pflag ```

second, you may not have some of the libraries installed so get this tool as well

```$ go get golang.org/x/tools/cmd/goimports ```

finally, install the tools using the following command:

```$ goimports -w user.go```

then run 
```$ go install```
and finally, 
```$ go run user.go cli.go -u githubuser1,githubuser2,etc```

output:
```
$ go run user.go cli.go -u torvalds,suraj44
Searching user(s): [torvalds suraj44]
Username:	 torvalds
Name:		 Linus Torvalds
Email:		 
Bio:		 

Username:	 suraj44
Name:		 Suraj Singh
Email:		 
Bio:		 I'm currently in my sophomore year, pursuing Computer Engineering at the National Institute of Technology Karnataka, India. I'm interested in systems and ML


```
