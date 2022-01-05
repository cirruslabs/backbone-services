# How to build the project

This repository uses [Bazel](https://bazel.build/). We recommend installing it via [Bazelisk](https://github.com/bazelbuild/bazelisk#about-bazelisk)
which will respect different Bazel versions defined in `.bazelversion` files.

## Dependency Management

[Gazelle](https://github.com/bazelbuild/bazel-gazelle) is used for generating `BUILD.bazel` files. To update dependencies
in BUILD files run the following command:

```bash
bazel run //:gazelle -- update
```

To add a new dependency use the regular `go get ...` to modify `go.mod` file and run `gazelle-update-repos` command
to pick up the changes:

```bash
bazel run //:gazelle-update-repos
```
