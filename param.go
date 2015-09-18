package vodka

import (
	"strconv"
)

// Param returns path parameter by name.
func (c *Context) Param(name string) (value string) {
	l := len(c.pnames)
	for i, n := range c.pnames {
		if n == name && i < l {
			value = c.pvalues[i]
			break
		}
	}
	return
}

func (c *Context) ParamInt(key string) (int, error) {
	return strconv.Atoi(c.Param(key))
}

func (c *Context) ParamInt32(key string) (int32, error) {
	v, err := strconv.ParseInt(c.Param(key), 10, 32)
	return int32(v), err
}

func (c *Context) ParamInt64(key string) (int64, error) {
	return strconv.ParseInt(c.Param(key), 10, 64)
}

func (c *Context) ParamUint(key string) (uint, error) {
	v, err := strconv.ParseUint(c.Param(key), 10, 64)
	return uint(v), err
}

func (c *Context) ParamUint32(key string) (uint32, error) {
	v, err := strconv.ParseUint(c.Param(key), 10, 32)
	return uint32(v), err
}

func (c *Context) ParamUint64(key string) (uint64, error) {
	return strconv.ParseUint(c.Param(key), 10, 64)
}

func (c *Context) ParamBool(key string) (bool, error) {
	return strconv.ParseBool(c.Param(key))
}

func (c *Context) ParamFloat32(key string) (float32, error) {
	v, err := strconv.ParseFloat(c.Param(key), 32)
	return float32(v), err
}

func (c *Context) ParamFloat64(key string) (float64, error) {
	return strconv.ParseFloat(c.Param(key), 64)
}
