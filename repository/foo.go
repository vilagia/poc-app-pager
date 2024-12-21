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

const QUERY_EXECUTION_TIME_SECOND = 100 * time.Millisecond

func Find(after FooId, limit uint, isActiveThreshould uint) []Foo {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := r.Intn(100)
	rng := lo.Range(int(limit) * 2)
	time.Sleep(QUERY_EXECUTION_TIME_SECOND)
	return lo.Map(rng, func(item int, index int) Foo {
		return Foo{
			ID:       FooId(item),
			IsActive: randomNumber < int(isActiveThreshould),
		}
	})
}
