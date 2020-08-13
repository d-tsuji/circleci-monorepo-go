# CircleCI Monorepo in Go

This is an example of how to set up CircleCI in a repository for the Go language's Monolepo configuration.

One of the most common patterns we see is that we only want to build and test modules that have diffs in their commits. However, if there are dependencies between Go Modules in go.mod, you'll want to test the caller as well if the caller's module changes, even if there are no differences as commits. We don't intend to test only jobs that have differences.

Take a look at [.cirleci/config.yml](https://github.com/d-tsuji/circleci-monorepo-go/blob/master/.circleci/config.yml).

## Sample Structure

Consider a Go repository that would consist of three Go Modules, as shown below.

```
|--common
|  |--README.md
|  |--go.mod
|--service1
|  |--README.md
|  |--go.mod
|--service2
|  |--README.md
|  |--go.mod
```
