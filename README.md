# gobyexample
Learning Go with https://gobyexample.com

## Workflow

Each chapter in the project is separated into its own file in the `src` directory and they are run using the `task` command provided by [taskfile.dev](https://taskfile.dev).

Each chapter in the guide is created in its own branch and then merged back to the main branch.

Each commit is signed and verified with a GPG key.

**Starting a new chapter:**

```
➜ NAME=chapterName task newchap
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

## Taskfile commands

Each chapter is run with the `task` command and needs to be defined in the `Taskfile.yaml` file using the following format:

```yaml
  values:
    desc: Chapter description
    cmds:
      - go run cmd/chapterName/main.go
    silent: true
```