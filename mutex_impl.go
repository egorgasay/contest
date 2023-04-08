package contest

func New() Mutex {
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
	return &MyMutex{
		ch: ch,
	}
}

func (m *MyMutex) LockChannel() <-chan struct{} {
	return m.ch
}

func (m *MyMutex) Lock() {
	<-m.ch

}

func (m *MyMutex) Unlock() {
	m.ch <- struct{}{}
	//}
}
