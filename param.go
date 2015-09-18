package vodka

import (
	"html/template"
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

/*
func (c *Context) ParamMustString(key string, defaults ...string) string {
	if len(key) == 0 {
		return ""
	}
	if key[0] != ':' && key[0] != '*' {
		key = ":" + key
	}

	for _, v := range *c {
		if v.Name == key {
			return v.Value
		}
	}
	if len(defaults) > 0 {
		return defaults[0]
	}
	return ""
}

func (c *Context) ParamMustStrings(key string, defaults ...[]string) []string {
	if len(key) == 0 {
		return []string{}
	}
	if key[0] != ':' && key[0] != '*' {
		key = ":" + key
	}

	var s = make([]string, 0)
	for _, v := range *c {
		if v.Name == key {
			s = append(s, v.Value)
		}
	}
	if len(s) > 0 {
		return s
	}
	if len(defaults) > 0 {
		return defaults[0]
	}
	return []string{}
}

func (c *Context) ParamMustEscape(key string, defaults ...string) string {
	if len(key) == 0 {
		return ""
	}
	if key[0] != ':' && key[0] != '*' {
		key = ":" + key
	}

	for _, v := range *c {
		if v.Name == key {
			return template.HTMLEscapeString(v.Value)
		}
	}
	if len(defaults) > 0 {
		return defaults[0]
	}
	return ""
}
*/

func (c *Context) ParamMustInt(key string, defaults ...int) int {
	v, err := c.ParamInt(key)
	if len(defaults) > 0 && err != nil {
		return defaults[0]
	}
	return v

}

func (c *Context) ParamMustInt32(key string, defaults ...int32) int32 {
	r, err := c.ParamInt32(key)
	if len(defaults) > 0 && err != nil {
		return defaults[0]
	}

	return int32(r)
}

func (c *Context) ParamMustInt64(key string, defaults ...int64) int64 {
	r, err := c.ParamInt64(key)
	if len(defaults) > 0 && err != nil {
		return defaults[0]
	}
	return r
}

func (c *Context) ParamMustUint(key string, defaults ...uint) uint {
	v, err := c.ParamUint(key)
	if len(defaults) > 0 && err != nil {
		return defaults[0]
	}
	return uint(v)
}

func (c *Context) ParamMustUint32(key string, defaults ...uint32) uint32 {
	r, err := c.ParamUint32(key)
	if len(defaults) > 0 && err != nil {
		return defaults[0]
	}

	return uint32(r)
}

func (c *Context) ParamMustUint64(key string, defaults ...uint64) uint64 {
	r, err := c.ParamUint64(key)
	if len(defaults) > 0 && err != nil {
		return defaults[0]
	}
	return r
}

func (c *Context) ParamMustFloat32(key string, defaults ...float32) float32 {
	r, err := c.ParamFloat32(key)
	if len(defaults) > 0 && err != nil {
		return defaults[0]
	}
	return float32(r)
}

func (c *Context) ParamMustFloat64(key string, defaults ...float64) float64 {
	r, err := c.ParamFloat64(key)
	if len(defaults) > 0 && err != nil {
		return defaults[0]
	}
	return r
}

func (c *Context) ParamMustBool(key string, defaults ...bool) bool {
	r, err := c.ParamBool(key)
	if len(defaults) > 0 && err != nil {
		return defaults[0]
	}
	return r
}

func (c *Context) ParamEscape(key string, defaults ...string) string {
	s := c.Param(key)
	if len(defaults) > 0 && len(s) == 0 {
		return defaults[0]
	}
	return template.HTMLEscapeString(s)
}
