package global

func GetGlobal() Global {
	globalSync.RLock()
	defer globalSync.RUnlock()

	return *global
}
