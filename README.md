This utility has been deprecated in favor of **DynamicBuildRequires** which have been successfuly integrated in rpm itself.

# mock-install
A small utility to request package installation within a [mock](https://github.com/rpm-software-management/mock) build root when using the [pm_request plugin](https://github.com/rpm-software-management/mock/wiki/Plugin-PMRequest).

## Usage:

Add inside your [rpm](http://rpm.org/) package [spec file](http://rpm.org/documentation.html):

```specfile
BuildRequires: mock-install
[…]

%prep
[…] # Do something to compute yourbr1 … yourbrX
mock-install yourbr1 … yourbrX
```

The command is a no-op if it can not detect *pm_request*, making the resulting spec file safe to use outside mock (you just have to install the needed packages some other way).

## Building

*mock-install* is written in [Go](https://golang.org/) using only basic standard Go packages, install a Go compiler and build it with:

```sh
go build -o mock-install cmd/mock-install/mock-install.go
```
