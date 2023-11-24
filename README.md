# Testcontainers with Bazel build tool

This example shows how to use [Testcontainers](https://www.testcontainers.org/) with Bazel build tool.

## Install Bazel
The easiest way to use Bazel is using [bazelisk](https://github.com/bazelbuild/bazelisk)

```bash
$ brew install bazelisk
```

## How to run?

```bash
$ bazelisk build //...
$ bazelisk test //... --test_output=all
$ bazelisk build //customers/... --test_output=all
$ bazelisk build //products/...
$ bazelisk test //customers/... --test_output=all
$ bazelisk test //products/... --test_output=all
```


## Update go dependencies

```bash
$ gazelle update-repos -from_file=go.work -to_macro=deps.bzl%go_deps -prune=True
```
