package vodka

// Query returns query parameter by name.
func (c *Context) Query(name string) string {
	if c.query == nil {
		c.query = c.request.URL.Query()
	}
	return c.query.Get(name)
}
