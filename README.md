# cli.tar

CLI tool that creates a TAR file from the list of inputs



# Prerequisites

Install [bazelisk](https://docs.bazel.build/versions/main/install-bazelisk.html):

```bash
go install github.com/bazelbuild/bazelisk@latest
ln -s "$GOPATH/bin/bazelisk" /usr/local/bin/bazel
```


# CLI TAR binary

```bash
bazel run //cmd/tar -- output.tar input1 input2 input3
```
