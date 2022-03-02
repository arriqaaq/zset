package zset

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func randZSet() *ZSet {
	n := New()

	rand.Seed(time.Now().Unix())
	s := "abcdefghijklmnopqrstuvwxyz"
	randomVal := func() (val string) {
		for i := 0; i < 12; i++ {
			val += string(s[rand.Intn(26)])
		}
		return
	}

	for i := 0; i < 100; i++ {
		n.ZAdd(testKey, float64(rand.Intn(100000)), randomVal(), nil)
	}

	return n
}

var testKey = "zset1"

func makeZSet() *ZSet {
	n := New()

	n.ZAdd(testKey, 1, "ced", nil)
	n.ZAdd(testKey, 2, "acd", nil)
	n.ZAdd(testKey, 3, "bcd", nil)
	n.ZAdd(testKey, 4, "acc", nil)
	n.ZAdd(testKey, 5, "mcd", nil)
	n.ZAdd(testKey, 6, "ccd", nil)
	n.ZAdd(testKey, 7, "ecd", nil)

	return n
}

func TestZSet_ZAdd(t *testing.T) {
	n := makeZSet()
	assert.Equal(t, 7, n.ZCard(testKey))
}

func TestZSet_ZScore(t *testing.T) {
	n := makeZSet()
	_, score := n.ZScore(testKey, "ced")
	assert.Equal(t, 1, int(score))
	_, score = n.ZScore(testKey, "ecd")
	assert.Equal(t, 7, int(score))
	ok, score := n.ZScore(testKey, "defrwefrw")
	assert.Equal(t, false, ok)
	assert.Equal(t, 0, int(score))
}

func TestZSet_ZRank(t *testing.T) {

	n := makeZSet()
	rank := n.ZRank(testKey, "ced")
	assert.Equal(t, 0, int(rank))
	rank = n.ZRank(testKey, "ecd")
	assert.Equal(t, 6, int(rank))
}

func TestZSet_ZRevRank(t *testing.T) {
	n := makeZSet()

	rank := n.ZRevRank(testKey, "ced")
	assert.Equal(t, 6, int(rank))
	rank = n.ZRevRank(testKey, "ecd")
	assert.Equal(t, 0, int(rank))
}

func TestZSet_ZIncrBy(t *testing.T) {
	n := makeZSet()

	n.ZIncrBy(testKey, 300, "ced")

	_, score := n.ZScore(testKey, "ced")
	assert.Equal(t, 301, int(score))
}
