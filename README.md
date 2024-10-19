# Hi! This prohect is made to create a prototype of a server and server conection using autorisation 

## Functionality 

### File `hash\hash_password.go`

- 2 simple functions to hash and compare hash

### File `server\new_server.go`

- A structure and a fuction to create a server
- You can edit it, for example using a database or a .json file to store data about the server

### File `server\new_conection.go`

- A structure and a method of structure *server* to create a conection to a server
- The conection is allowed only by entering the password of the server

### File `server\conection_methods.go`

- Two methods of *server_conection*

1. `CheckConection()`. Checks that the conection wasn't timed out or ha too many operations
2. `Touch()`. Increacing the number of operations with the conection

### File `server/server_methods.go`

- This is a file for your imagination!
- Add methods for your server!
- Here are some examples:
    1. Database oprations
    2. Web server methods
    3. Other server operations

## Tests

- For my program I've made some basic tests in `main_test.go` file!

### Error conection

- Creates a conection with wrong password
- Wants:
    > Eror 401 Unauthorized: Passwords do not match

### Normal conection

- Creates a conection that has no errors

### Too many operations

- Tryes to make operations more than allowed
- Wants:
    > Error 429 Too Many Requests: Too many operations with server conection

### Timeout

- Timeouts the conection and tryes to use it
- Wants: 
    > Error 408 Request Timeout: Server connection timed out

## How to get the project? 

- Install **golang** [here](https://go.dev/doc/install)
- Insatall **git** [here](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- Use a command `git clone https://github.com/VandiKond/server_prototype.git`
- To get packages you need write `go get golang.org/x/crypto`. If you have already installed it you do not need to write it
____
- Or downdload a .zip file of the project [here](https://github.com/VandiKond/server_prototype/archive/refs/heads/master.zip)

## How to run the project? 

- If you want to run the tests, i've made write in the console `go test`
____

- You can also edit and run the main.go file
- `go run main.go`