package vodka

type (
	Group struct {
		vodka Vodka
	}
)

func (g *Group) Use(m ...Middleware) {
	for _, h := range m {
		g.vodka.middleware = append(g.vodka.middleware, wrapMiddleware(h))
	}
}

func (g *Group) Connect(path string, h Handler) {
	g.vodka.Connect(path, h)
}

func (g *Group) Delete(path string, h Handler) {
	g.vodka.Delete(path, h)
}

func (g *Group) Get(path string, h Handler) {
	g.vodka.Get(path, h)
}

func (g *Group) Head(path string, h Handler) {
	g.vodka.Head(path, h)
}

func (g *Group) Options(path string, h Handler) {
	g.vodka.Options(path, h)
}

func (g *Group) Patch(path string, h Handler) {
	g.vodka.Patch(path, h)
}

func (g *Group) Post(path string, h Handler) {
	g.vodka.Post(path, h)
}

func (g *Group) Put(path string, h Handler) {
	g.vodka.Put(path, h)
}

func (g *Group) Trace(path string, h Handler) {
	g.vodka.Trace(path, h)
}

func (g *Group) WebSocket(path string, h HandlerFunc) {
	g.vodka.WebSocket(path, h)
}

func (g *Group) Static(path, root string) {
	g.vodka.Static(path, root)
}

func (g *Group) ServeDir(path, root string) {
	g.vodka.ServeDir(path, root)
}

func (g *Group) ServeFile(path, file string) {
	g.vodka.ServeFile(path, file)
}

func (g *Group) Group(prefix string, m ...Middleware) *Group {
	return g.vodka.Group(prefix, m...)
}
