package universe

import (
	"context"
	"sync"
)

type Datasource struct {
	day uint
	mu  sync.Mutex
}

var _ IDataSource = &Datasource{}

func ProvideMemoryDatasource() (*Datasource, error) {
	return &Datasource{
		day: 0,
	}, nil
}

func (d *Datasource) SetCurrentDay(ctx context.Context, day uint) error {
	d.mu.Lock()
	d.day = day
	d.mu.Unlock()

	return nil
}

func (d *Datasource) IncrementCurrentDay(ctx context.Context) (day uint, err error) {
	d.mu.Lock()
	d.day = d.day + 1
	d.mu.Unlock()

	return d.day, nil
}
