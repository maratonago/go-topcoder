# go-topcoder #

go-topcoder is a Go client library for accessing the [Topcoder API][].

[![license](https://img.shields.io/github/license/maratonago/go-topcoder.svg)](https://github.com/maratonago/go-topcoder/blob/master/LICENSE)
[![Build Status](https://travis-ci.org/maratonago/go-topcoder.svg?branch=master)](https://travis-ci.org/maratonago/go-topcoder)
[![Test Coverage](https://coveralls.io/repos/github/maratonago/go-topcoder/badge.svg?branch=master)](https://coveralls.io/github/maratonago/go-topcoder?branch=master)
[![GoDoc](https://godoc.org/github.com/maratonago/go-topcoder/topcoder?status.svg)](https://godoc.org/github.com/maratonago/go-topcoder/topcoder)

go-topcoder requires Go version 1.6.2 or greater.

## Getting Started ##

### Usage ###
#### As a library: ####
```go
import "github.com/maratonago/go-topcoder/topcoder"
```

#### Sample App ####
```go
import (
  "encoding/json"
  "fmt"

  "github.com/maratonago/go-topcoder/topcoder"
)

func main() {
  profile, _, _ := topcoder.UserProfile.PublicProfile("paulocezar")

  sProfile, err := json.MarshalIndent(profile, "", "  ")
  if err != nil {
    panic(err)
  }

  fmt.Println(string(sProfile))
}
```

Above code would produce an output similar to:
```json
{
  "handle": "paulocezar",
  "country": "Brazil",
  "memberSince": "2009-10-08T18:26:00.000-04:00",
  "quote": "",
  "photoLink": "https://topcoder-prod-media.s3.amazonaws.com/member/profile/paulocezar-1447891532103.png",
  "copilot": false,
  "overallEarning": 250,
  "ratingSummary": [
    {
      "name": "Algorithm",
      "rating": 1448,
      "colorStyle": "color: #6666FF"
    }
  ],
  "Achievements": [
    {
      "date": "2016-03-15T00:00:00.000-04:00",
      "description": "First Marathon Competition"
    },
    {
      "date": "2013-09-17T00:00:00.000-04:00",
      "description": "One Hundred Rated Algorithm Competitions"
    },
    {
      "date": "2013-03-03T00:00:00.000-05:00",
      "description": "Solved Hard Div1 Problem in SRM"
    }
  ]
}
```

## LICENSE ##

Released under the [MIT License](https://github.com/maratonago/go-topcoder/blob/master/LICENSE).


[Topcoder API]: http://docs.tcapi.apiary.io/
[Topcoder Data Feeds]: http://apps.topcoder.com/wiki/display/tc/Algorithm+Data+Feeds
