package cli

type GlobalInterface interface {
	GlobalConfigure(g *Global)
	Lock()
	Unlock()
}
