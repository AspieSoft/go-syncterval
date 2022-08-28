# Go Syncterval

[![donation link](https://img.shields.io/badge/buy%20me%20a%20coffee-square-blue)](https://buymeacoffee.aspiesoft.com)

A synchronized interval for go.

Rather than creating a new go routine for every repetitive task, if that task is low priority, you can use this module to share one routine with other modules.

## Installation

```shell script

  go get github.com/AspieSoft/go-syncterval

```

## Usage

```go

import (
  "fmt"
  "time"

  "github.com/AspieSoft/go-syncterval"
)

func loopFn(){
  fmt.Println("This Loop Will Run Every 3 Seconds!")
}

func main(){
  syncterval.New(3 * time.Second, loopFn)
}

```
