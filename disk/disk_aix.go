//go:build aix
// +build aix

package disk

import (
	"context"

	"github.com/vfpinheiro/gopsutil/internal/common"
)

func IOCountersWithContext(ctx context.Context, names ...string) (map[string]IOCountersStat, error) {
	return nil, common.ErrNotImplementedError
}

func SerialNumberWithContext(ctx context.Context, name string) (string, error) {
	return "", common.ErrNotImplementedError
}

func LabelWithContext(ctx context.Context, name string) (string, error) {
	return "", common.ErrNotImplementedError
}
