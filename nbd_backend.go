package main

type NBDBackend struct {
	volume *Volume
}

func NewNBDBackend(volume *Volume) *NBDBackend {
	return &NBDBackend{
		volume: volume,
	}
}

func (b *NBDBackend) Size() (int64, error) {
	return b.volume.size, nil
}

func (b *NBDBackend) Sync() error {
	return b.volume.Flush()
}

func (b *NBDBackend) ReadAt(p []byte, off int64) (int, error) {
	return b.volume.ReadAt(p, off)
}

func (b *NBDBackend) WriteAt(p []byte, off int64) (int, error) {
	return b.volume.WriteAt(p, off)
}
