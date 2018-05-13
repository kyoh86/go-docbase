# go-docbase

[![Go Report Card](https://goreportcard.com/badge/github.com/kyoh86/go-docbase)](https://goreportcard.com/report/github.com/kyoh86/go-docbase)
[![CircleCI](https://img.shields.io/circleci/project/github/kyoh86/go-docbase.svg)](https://circleci.com/gh/kyoh86/go-docbase)
[![Coverage Status](https://img.shields.io/codecov/c/github/kyoh86/go-docbase.svg)](https://codecov.io/gh/kyoh86/go-docbase)


A Go library for accessing the [Docbase](https://docbase.io/) API v1.

API Docs: https://help.docbase.io/posts/45703

## Install

```sh
go get github.com/kyoh86/go-docbase
```

## Usage

```go
import (
	"github.com/kyoh86/go-docbase/docbase"
)

transport := docbase.TokenTransport{Token: "Your API Token"}
client := docbase.NewClient(transport.Client())
...
```

## API Coverage Status

* ○: Implemented and tested.
* △: Implementing.
* ×: Not implemented.

| Service | Function | Status |
| --- | --- | --- |
| Post | List | ○ |
| Post | Create | ○ |
| Post | Get | △ |
| Post | Update | △ |
| Post | Delete | △ |
| Comment | Create | △ |
| Comment | Delete | △ |
| Team | List | ○ |
| Group | List | △ |
| Tag | List | △ |
| Attachment | Post | × |

# LICENSE

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)](http://www.opensource.org/licenses/MIT)

This is distributed under the [MIT License](http://www.opensource.org/licenses/MIT).
