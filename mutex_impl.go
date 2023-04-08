package contest

// New constructs a new Mutex
func New() Mutex {
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
	return &MyMutex{
		ch: ch,
	}
}

// LockChannel provides the main mutex channel.
func (m *MyMutex) LockChannel() <-chan struct{} {
	return m.ch
}

// Lock acquires lock.
func (m *MyMutex) Lock() {
	<-m.ch

}

// Unlock releases lock.
func (m *MyMutex) Unlock() {
	m.ch <- struct{}{}
}
