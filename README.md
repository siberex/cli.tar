# cli.tar

CLI tool that creates a TAR file from the list of inputs.

Also provides custom Bazel rules `pkg_tar` and `file_size`


# Prerequisites

Install [bazelisk](https://docs.bazel.build/versions/main/install-bazelisk.html):

```bash
go install github.com/bazelbuild/bazelisk@latest
ln -s "$GOPATH/bin/bazelisk" /usr/local/bin/bazel
```


# Bazel targets

File size custom rule:

```bash
bazel run //rules:size
```

pkg_tar custom rule:

```bash
bazel build //rules:archive
```

CLI TAR go binary:

```bash
bazel run //cmd/tar -- output.tar input1 input2 input3
```


# Linters

```bash
bazel run //:buildifier
```

# Tests

```bash
bazel test //rules:pkg_tar_test
```
