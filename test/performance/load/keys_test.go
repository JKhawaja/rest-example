// +build load

package load

import (
	"testing"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

// NOTE: must have instance of server running on localhost:8080 before
// performing these load tests.

// NOTE: load testing exceeds API rate limits for third-party services usually

func TestLoadListKeys(t *testing.T) {
	t.Run("LoadListKeys", func(t *testing.T) {
		rate := uint64(100) // per second
		duration := 4 * time.Second
		targeter := vegeta.NewStaticTargeter(vegeta.Target{
			Method: "POST",
			Body:   []byte(`["dave", "tom", "john"]`),
			URL:    "http://localhost:8080/keys",
		})
		attacker := vegeta.NewAttacker()

		var metrics vegeta.Metrics
		for res := range attacker.Attack(targeter, rate, duration) {
			metrics.Add(res)
		}
		metrics.Close()

		t.Logf("99th percentile: %s\n", metrics.Latencies.P99)
	})
}
