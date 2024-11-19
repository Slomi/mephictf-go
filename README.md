# mephictf-go

## Usage guide

### Setting up

1. Fork this repository. You might want to make it private, if you want to hide your solutions from others.
2. Run `git clone <your_repo_url>` on your local machine.
3. Run `git remote add upstream https://github.com/LeKSuS-04/mephictf-go.git` to add repository with tasks as remote source.
4. Run `git config --local pull.rebase true` to ensure that new tasks will be pulled correctly.

### Solving tasks

1. Write code in files that don't end with `_test.go`.
2. Run `go test ./<path>` to test specific task or `go test ./...` to run all tests in the repository.
3. Commit your solution with `git add <path>` and `git commit -m "<commit_msg>"` and push it to github with `git push`.
4. View results of automatic testing on `actions` tab in github.

### Pulling new tasks

1. Make sure that all of your changes are commited.
2. Run `git pull upstream master` to pull new tasks.
3. Although unlikely, you might stumble into merge conflicts. [Here's how to resolve them](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/addressing-merge-conflicts/resolving-a-merge-conflict-using-the-command-line).

## Code snippets and examples from lectures

1. [1-basic](./examples/1-basic/)

## Tasks

1. Basics of the language
   - [helloworld](/helloworld/helloworld.go)
   - [quickmafs](/quickmafs/quickmafs.go)
   - [lrucache](/lrucache/lrucache.go)
   - [packagemanager](/packagemanager/packagemanager.go)
2. Packages, structs, interfaces, stdlib and iterators
   - [shapes](/shapes/shapes.go)
   - [functional](/functional/functional.go)
   - [logparser](/logparser/logreader.go)

## Contribution

Feel free to open PRs to this repository if you find any mistakes or inaccuracies. If you wish to discuss them beforehand, open an issue.
