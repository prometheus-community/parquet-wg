package parquetwg

import (
	"testing"

	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/util/teststorage"
	"github.com/thanos-io/objstore"
)

func BenchmarkPrometheus(b *testing.B) {
	promQueryableCreate := func(tb testing.TB, bkt objstore.Bucket, st *teststorage.TestStorage) storage.Queryable {
		return st
	}

	RunBenchmarks(b, promQueryableCreate)
}
