package zk_smt

import (
	"github.com/gatechain/gate-zk-smt/metrics"
	"github.com/panjf2000/ants/v2"
)

// Option is a function that configures SMT.
type Option func(*GateSparseMerkleTree)

func InitializeVersion(version Version) Option {
	return func(smt *GateSparseMerkleTree) {
		smt.version = version
	}
}

func BatchSizeLimit(limit int) Option {
	return func(smt *GateSparseMerkleTree) {
		smt.batchSizeLimit = limit
	}
}

func DBCacheSize(size int) Option {
	return func(smt *GateSparseMerkleTree) {
		smt.dbCacheSize = size
	}
}

func GoRoutinePool(pool *ants.Pool) Option {
	return func(smt *GateSparseMerkleTree) {
		smt.goroutinePool = pool
	}
}

func GCThreshold(threshold uint64) Option {
	return func(smt *GateSparseMerkleTree) {
		if smt.gcStatus != nil {
			smt.gcStatus.threshold = threshold
			smt.gcStatus.segment = threshold / 10
		}

	}
}

func EnableMetrics(metrics metrics.Metrics) Option {
	return func(smt *GateSparseMerkleTree) {
		smt.metrics = metrics
	}
}
