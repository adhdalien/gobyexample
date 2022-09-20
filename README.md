# gobyexample
Learning Go with https://gobyexample.com

## Structure

Each chapter in the project is separated into its own file in the `src` directory and they are run using the `task` command provided by [taskfile.dev](https://taskfile.dev).

**Starting a new chapter:**

```
➜ NAME=chapterName task newchap
```

**Finishing a chapter:**

```
➜ NAME=chapterName task endchap -- "Commit message"
```

## Workflow

The project uses [gitversion.net](https://gitversion.net) for [semantic versioning](http://semver.org).

Each chapter in the guide is created in its own branch and then merged back to the main branch. GitVersion is configured to automatically increment the version whenever a branch is merged to main.

Each commit is signed and verified with a GPG key.

```
$ git checkout -b chapterName
$ mkdir cmd/chapterName && touch cmd/chapterName/main.go
$ code cmd/chapterName/main.go
$ git add . && git commit -S -s -m "chapterName implementation"
$ git checkout main && git merge chapterName
$ git push
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