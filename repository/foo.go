package repository

import (
	"math/rand"
	"time"

	"github.com/samber/lo"
)

type FooId uint
type Foo struct {
	ID       FooId
	IsActive bool
}

const QUERY_EXECUTION_TIME_SECOND = "0.1"

func Find(after FooId, limit uint, isActiveThreshould uint) []Foo {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := r.Intn(100)
	rng := lo.Range(int(limit) * 2)
	time.Sleep(100 * time.Millisecond)
	return lo.Map(rng, func(item int, index int) Foo {
		return Foo{
			ID:       FooId(item),
			IsActive: randomNumber < int(isActiveThreshould),
		}
	})
}
