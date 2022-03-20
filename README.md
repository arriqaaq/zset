# zset

Getting Started
===============

## Installing

To start using hash, install Go and run `go get`:

```sh
$ go get -u github.com/arriqaaq/zset
```

This will retrieve the library.

## Usage

```go
package main

import (
	"fmt"

	"github.com/arriqaaq/zset"
)

type kv struct{ k, v int }

func main() {
	key := "zset1"

	// ZAdd (accepts any value)
	n := zset.New()
	n.ZAdd(key, 1, "ced", nil)
	n.ZAdd(key, 1, "efg", &kv{1, 2})

	// ZScore
	_, score := n.ZScore(key, "ced")
	fmt.Println("score: ", int(score))
	// score: 1

	// ZRank
	rank := n.ZRank(key, "ced")
	fmt.Println("zrank: ", int(rank))
	// zrank: 0

	// ZRevRank
	rank = n.ZRevRank(key, "ced")
	fmt.Println("zrevrank: ", int(rank))
	// zrevrank: 1

	// ZIncrBy
	n.ZIncrBy(key, 300, "ced")
	_, score = n.ZScore(key, "ced")
	fmt.Println("score: ", int(score))
	// score: 301
}
```

## Supported Commands

```go
Supported  commands

ZADD
ZCARD
ZINCRBY
ZPOPMAX
ZPOPMIN
ZRANGE
ZRANGEBYSCORE
ZRANGEWITHSSCORES
ZRANK
ZREM
ZREVRANGE
ZREVRANGEWITHSSCORES
ZREVRANK
ZSCORE
```
