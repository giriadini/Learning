**Installation**

https://golang.org/doc/install

**Test the installation**

*go version*

**Run Go in Docker**

We can also run go in a small docker container: 

*cd golang\introduction*

- *docker build --target dev . -t go*
- *docker run -it -v ${PWD}:/work go sh*
- *go version*

**Hello world**
**Create mod file**

*go mod init  github.com/Learning/golang/introduction/app*

**Run the code**

*go run app.go* 

**Build the binary**

*go build*
 
**run the binary**

*./app*

**install**

*go install  github.com/Learning/golang/introduction/app*


**Install code runner to run the code in vscode**

**check box allow coderuner to run from terminal to enter user inputs**