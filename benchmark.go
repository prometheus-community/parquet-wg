package parquetwg

import (
	"context"
	"math"
	"testing"

	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/util/teststorage"
	"github.com/thanos-io/objstore"
)

// QueryCreateFunc is a callback to create a Queryable from a pre-filled TestStorage. It may or may not use the provided Bucket.
// If the provided bucket is used, it will be used to measure object storage operations or to inject latency.
type QueryCreateFunc func(tb testing.TB, bkt objstore.Bucket, st *teststorage.TestStorage) storage.Queryable

func RunBenchmarks(b *testing.B, f QueryCreateFunc) {
	ctx := context.Background()

	// Example trivial function, add more below for interesting scenarios
	b.Run("simple", func(bb *testing.B) {
		st := teststorage.New(bb)
		bb.Cleanup(func() { _ = st.Close() })
		bkt := objstore.NewInMemBucket()
		bb.Cleanup(func() { _ = bkt.Close() })

		q, err := f(bb, bkt, st).Querier(math.MinInt64, math.MaxInt64)
		if err != nil {
			bb.Fatal("error building querier: ", err)
		}

		bb.ReportAllocs()
		bb.ResetTimer()

		for i := 0; i < b.N; i++ {
			q.Select(ctx, false, nil)
		}
	})
}
