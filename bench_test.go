package main

import (
	"poc-app-pager/repository"
	"testing"

	"github.com/samber/lo"
)

func Benchmark_QueryOnly(b *testing.B) {
	b.ResetTimer()
	for range b.N {
		repository.Find(1, 100, 100)
	}
}

func Benchmark_QueryWithFilterPass100Percent(b *testing.B) {
	b.ResetTimer()
	retryCounts := []uint{}
	for range b.N {
		retryCounts = append(retryCounts, findWithRefetch(100))
	}

	averageRetryCount := lo.Mean(retryCounts)
	b.Log("average", averageRetryCount)
}

func Benchmark_QueryWithFilterPass75Percent(b *testing.B) {
	b.ResetTimer()
	retryCounts := []uint{}
	for range b.N {
		retryCounts = append(retryCounts, findWithRefetch(75))
	}

	averageRetryCount := lo.Mean(retryCounts)
	b.Log("average", averageRetryCount)
}

func Benchmark_QueryWithFilterPass50Percent(b *testing.B) {
	b.ResetTimer()
	retryCounts := []uint{}
	for range b.N {
		retryCounts = append(retryCounts, findWithRefetch(50))
	}

	averageRetryCount := lo.Mean(retryCounts)
	b.Log("average", averageRetryCount)
}

func Benchmark_QueryWithFilterPass25Percent(b *testing.B) {
	b.ResetTimer()
	retryCounts := []uint{}
	for range b.N {
		retryCounts = append(retryCounts, findWithRefetch(25))
	}

	averageRetryCount := lo.Mean(retryCounts)
	b.Log("average", averageRetryCount)
}

func Benchmark_QueryWithFilterPass10Percent(b *testing.B) {
	b.ResetTimer()
	retryCounts := []uint{}
	for range b.N {
		retryCounts = append(retryCounts, findWithRefetch(10))
	}

	averageRetryCount := lo.Mean(retryCounts)
	b.Log("average", averageRetryCount)
}

func Benchmark_QueryWithFilterPass5Percent(b *testing.B) {
	b.ResetTimer()
	retryCounts := []uint{}
	for range b.N {
		retryCounts = append(retryCounts, findWithRefetch(5))
	}

	averageRetryCount := lo.Mean(retryCounts)
	b.Log("average", averageRetryCount)
}

func Benchmark_QueryWithFilterPass4Percent(b *testing.B) {
	b.ResetTimer()
	retryCounts := []uint{}
	for range b.N {
		retryCounts = append(retryCounts, findWithRefetch(4))
	}

	averageRetryCount := lo.Mean(retryCounts)
	b.Log("average", averageRetryCount)
}


func Benchmark_QueryWithFilterPass3Percent(b *testing.B) {
	b.ResetTimer()
	retryCounts := []uint{}
	for range b.N {
		retryCounts = append(retryCounts, findWithRefetch(3))
	}

	averageRetryCount := lo.Mean(retryCounts)
	b.Log("average", averageRetryCount)
}


func Benchmark_QueryWithFilterPass2Percent(b *testing.B) {
	b.ResetTimer()
	retryCounts := []uint{}
	for range b.N {
		retryCounts = append(retryCounts, findWithRefetch(2))
	}

	averageRetryCount := lo.Mean(retryCounts)
	b.Log("average", averageRetryCount)
}

func Benchmark_QueryWithFilterPass1Percent(b *testing.B) {
	b.ResetTimer()
	retryCounts := []uint{}
	for range b.N {
		retryCounts = append(retryCounts, findWithRefetch(1))
	}

	averageRetryCount := lo.Mean(retryCounts)
	b.Log("average", averageRetryCount)
}

func findWithRefetch(threshould uint) uint {
	result := []repository.Foo{}

	var retryCount uint
	for {
		foos := repository.Find(1, 100, threshould)
		for _, foo := range foos {
			if foo.IsActive {
				result = append(result, foo)
			}
		}

		if len(result) >= 100 {
			break
		}
		retryCount++
	}
	return retryCount
}
