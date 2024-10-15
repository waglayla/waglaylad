
waglaylad - Only for Stratum Bridge & Pool Operators (Deprecated)
====

[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://choosealicense.com/licenses/isc/)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/waglayla/waglaylad)

#### Build from Source

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Ensure Go was installed properly and is a supported version:

```bash
$ go version
```

- Run the following commands to obtain and install waglaylad including all dependencies:

```bash
$ git clone https://github.com/waglayla/waglaylad
$ cd waglaylad
$ go install . ./cmd/...
```

- waglaylad (and utilities) should now be installed in `$(go env GOPATH)/bin`. If you did
  not already add the bin directory to your system path during Go installation,
  you are encouraged to do so now.


## Getting Started

waglaylad has several configuration options available to tweak how it runs, but all
of the basic operations work with zero configuration.

```bash
$ waglaylad
```

## License

waglaylad is licensed under the copyfree [ISC License](https://choosealicense.com/licenses/isc/).
