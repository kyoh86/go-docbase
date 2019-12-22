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

### v1

```go
import (
	"github.com/kyoh86/go-docbase/docbase"
)

transport := docbase.TokenTransport{Token: "Your API Token"}
client := docbase.NewClient(transport.Client())
...
```

And see [example](./cmd/go-docbase-sample/main.go).

### v2

```go
import (
	"github.com/kyoh86/go-docbase/docbase/v2"
)

client := docbase.NewAuthClient("Your DocBase Domain", "Your API Token")
```

And see [example](./v2/cmd/go-docbase-sample/main.go).

## API Coverage Status

### v1

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

### v2
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
