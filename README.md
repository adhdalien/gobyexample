# gobyexample
Learning Go with https://gobyexample.com

## Workflow

Each chapter in the project is separated into its own file in the `src` directory and they are run using the `task` command provided by [taskfile.dev](https://taskfile.dev).

Each chapter in the guide is created in its own branch and then merged back to the main branch.

Each commit is signed and verified with a GPG key.

**Starting a new chapter:**

```
➜ task newchap -- chapterName
task: [newchap] git checkout -b chapterName
Switched to a new branch 'chapterName'
task: [newchap] mkdir cmd/chapterName
task: [newchap] touch cmd/chapterName/main.go
task: [newchap] echo "package main" >> cmd/chapterName/main.go
task: [newchap] echo "func main() {}" >> cmd/chapterName/main.go
```

**Finishing a chapter:**

```
➜ NAME=chapterName task endchap -- "Commit message"
```

**Running the code for a specific chapter:**

```
➜ task runchap -- structs
task: [runchap] go run cmd/structs/main.go
{Bob 20}
{Alice 30}
{Fred 0}
&{Ann 40}
&{Jon 42}
Sean
50
51
```