# gobyexample
Learning Go with https://gobyexample.com

## Structure

Each chapter in the project is separated into its own file in the `src` directory and they are run using the `task` command provided by [taskfile.dev](https://taskfile.dev).

```
➜ task --list
task: Available tasks for this project:
* helloworld:   Hello world
```

## Workflow

The project uses [gitversion.net](https://gitversion.net) for [semantic versioning](http://semver.org).

Each chapter in the guide is created in its own branch and then merged back to the main branch. GitVersion is configured to automatically increment the version whenever a branch is merged to main.

```

```