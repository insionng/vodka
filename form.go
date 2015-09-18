package vodka

// Form returns form parameter by name.
func (c *Context) Form(name string) string {
	return c.request.FormValue(name)
}
