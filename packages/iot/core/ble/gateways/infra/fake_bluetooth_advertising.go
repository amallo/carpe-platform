//go:build !tinygo

package infra

type FakeBluetoothAdvertising struct {
	started      bool
	startCallCount int
}

func NewFakeBluetoothAdvertising() *FakeBluetoothAdvertising {
	return &FakeBluetoothAdvertising{
		started:      false,
		startCallCount: 0,
	}
}

func (f *FakeBluetoothAdvertising) Start() error {
	f.started = true
	f.startCallCount++
	return nil
}

func (f *FakeBluetoothAdvertising) IsStarted() bool {
	return f.started
}

func (f *FakeBluetoothAdvertising) StartCallCount() int {
	return f.startCallCount
}
