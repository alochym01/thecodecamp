# thecodecamp
## Using Makefile

1. Content of Makefile
    ```bash

    run:
        go run cmd/web/*

    mod:
        # real tab space or get error Makefile:2: *** missing separator.  Stop.
        go mod tidy # remove unused go packages
        go mod vendor # make local copy of third party packages

    run cli:

    make run - to run application

    make mod - to run go module
    ```

## Live reload in Development

1. Install air `go get -u github.com/cosmtrek/air`
2. Set shell enviroment by add `GOPATH="/home/hadn/go/bin"` to `~/.bash_profile`
    ```bash
    export PATH="$PATH:/home/hadn/go/bin/"
    ```
3. run `air init` to create air config file should be in same folder with `main.go`
4. air config file - `cmd/web/.air.toml`
   ```bash
   root = "."
   tmp_dir = "tmp"

   [build]
     bin = "./tmp/main"
     cmd = "go build -o ./tmp/main ./cmd/web/main.go"
     delay = 1000
     exclude_dir = ["assets", "tmp", "vendor"]
     exclude_file = []
     exclude_regex = []
     exclude_unchanged = false
     follow_symlink = false
     full_bin = ""
     include_dir = ["cmd", "config", "handler", "models", "repository", "router"]
     include_ext = ["go", "tpl", "tmpl", "html"]
     kill_delay = "0s"
     log = "build-errors.log"
     send_interrupt = false
     stop_on_error = true

   [color]
     app = ""
     build = "yellow"
     main = "magenta"
     runner = "green"
     watcher = "cyan"

   [log]
     time = false

   [misc]
     clean_on_exit = false
   ```
5. run `air` for auto reload

## Using curl for GET/POST/PUT Method

1. Syntax **curl -X POST -H "Content-Type: application/json" -d @FILENAME DESTINATION**
2. POST Method with a record data in json format
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"id": "4", "title": "Blue Train 04", "artist": "Do Nguyen Ha", "price": 6.99}' http://localhost:8080/albums

    curl -X POST -H "Content-Type: application/json" -d @album.json localhost:8080/albums

    // album.json file content
    {"id": "4", "title": "Blue Train 04", "artist": "Do Nguyen Ha", "price": 6.99}
    ```
3. GET Method
    ```bash
    curl -H "Content-Type: application/json" localhost:8080/albums
    ```
4. PUT Method
    ```bash
    curl -X PUT -H "Content-Type: application/json" -d @album.json localhost:8080/albums

    // album.json file content
    {"id": "4", "title": "Blue Train 04", "artist": "Do Nguyen Ha", "price": 66.99}
    ```

## Using git

1. Switch to branch. if branch doesnt exist it will be created
    1. git branch abc
    1. git checkout abc
1. Delete local branch
    1. git branch -D abc
1. Apply to git
    1. git push -u origin master
1. Create git tag
   1. git tag v0.0.1
   1. git push --tags
1. Depending on your exact situation, there are three useful ways to use git diff:
   1. Show differences between index and working tree; that is, changes you haven't staged to commit:
      1. git diff [filename]
   1. Show differences between current commit and index; that is, what you're about to commit (--staged does exactly the same thing, use what you like):
      1. git diff --cached [filename]
   1. Show differences between current commit and working tree:
      1. git diff HEAD [filename]
      1. git diff works recursively on directories, and if no paths are given, it shows all changes.