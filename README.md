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

import "github.com/arriqaaq/zset"

type kv struct{k,v string}

func main() {
    key:="zset1"

    // ZAdd (accepts any value)
    n := zset.New()
	n.ZAdd(key, 1, "ced", nil)
	n.ZAdd(key, 1, "efg", &kv{1,2})

    // ZScore
	_, score := n.ZScore(testKey, "ced")
	assert.Equal(t, 1, int(score))

    // ZRank
	rank := n.ZRank(testKey, "ced")
	assert.Equal(t, 0, int(rank))

    // ZRevRank
	rank := n.ZRevRank(testKey, "ced")
	assert.Equal(t, 6, int(rank))

    // ZIncrBy
	n.ZIncrBy(testKey, 300, "ced")
	_, score := n.ZScore(testKey, "ced")
	assert.Equal(t, 301, int(score))

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
