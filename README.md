# go-docbase

A Go library for accessing the [Docbase](https://docbase.io)

[![Go Report Card](https://goreportcard.com/badge/github.com/kyoh86/go-docbase)](https://goreportcard.com/report/github.com/kyoh86/go-docbase)
[![Coverage Status](https://img.shields.io/codecov/c/github/kyoh86/go-docbase.svg)](https://codecov.io/gh/kyoh86/go-docbase)

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

client := docbase.NewAuthClient("Your DocBase Domain", "Your API Token")
...
```

And see [example](./cmd/go-docbase-sample/main.go).

## API Coverage Status

| Service | Function | Implemented | Tested |
| --- | --- | --- | --- |
| Post | List | ☑ | ☑ |
| Post | Create | ☑ | ☑ |
| Post | Get | ☑ | ☑ |
| Post | Edit | ☑ | ☑ |
| Post | Archive | ☑ | ☑ |
| Post | Unarchive | ☑ | ☑ |
| Post | Delete | ☑ | ☑ |
| User | List | ☑ | ☑ |
| Comment | Create | ☑ | ☑ |
| Comment | Delete | ☑ | ☑ |
| Attachment | Upload | ☑ | ☑ |
| Tag | List | ☑ | ☑ |
| Group | Create | ☑ | ☑ |
| Group | Get | ☑ | ☑ |
| Group | List | ☑ | ☑ |
| Group | AddUsers | ☑ | ☑ |
| Group | RemoveUsers | ☑ | ☑ |

# LICENSE

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)](http://www.opensource.org/licenses/MIT)

This is distributed under the [MIT License](http://www.opensource.org/licenses/MIT).
