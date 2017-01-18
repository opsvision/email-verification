# Email Verification Utility [![Build Status](https://travis-ci.org/opsvision/email-verification.svg?branch=master)](https://travis-ci.org/opsvision/email-verification) [![Go Report Card](https://goreportcard.com/badge/github.com/opsvision/email-verification)](https://goreportcard.com/report/github.com/opsvision/email-verification)
A small utility tool for verifying email addresses

## Download and Build
```
$ git clone http://github.com/opsvision/email-verification
$ go build -o verify main.go
```

## Usage
```
Usage of ./verify:
  -email string
        the email address to verify (default "jdoe@acme.com")
  -sender string
        the sender email address (default "jdoe@acme.com")
```

### Example
```
$ ./verify -email jdoe@acme.com
2017/01/18 15:13:22 Checking email jdoe@acme.com
jdoe@acme.com|INVALID
```

## Issues
Note that this utility can return false positives depending on how the remote server is configured. Use at your own risk.

## Acknowledgement
* Author: [@dishmael](https://github.com/dishmael)
* Company: [@opsvision](https://github.com/opsvision)
