//go:build !darwin && !linux && !freebsd && !openbsd && !windows && !solaris && !aix
// +build !darwin,!linux,!freebsd,!openbsd,!windows,!solaris,!aix

package disk

import (
	"context"

	"github.com/vfpinheiro/gopsutil/internal/common"
)

func IOCountersWithContext(ctx context.Context, names ...string) (map[string]IOCountersStat, error) {
	return nil, common.ErrNotImplementedError
}

func PartitionsWithContext(ctx context.Context, all bool) ([]PartitionStat, error) {
	return []PartitionStat{}, common.ErrNotImplementedError
}

func UsageWithContext(ctx context.Context, path string) (*UsageStat, error) {
	return nil, common.ErrNotImplementedError
}

func SerialNumberWithContext(ctx context.Context, name string) (string, error) {
	return "", common.ErrNotImplementedError
}

func LabelWithContext(ctx context.Context, name string) (string, error) {
	return "", common.ErrNotImplementedError
}
