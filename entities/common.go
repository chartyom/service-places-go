package entities

type Common struct{}

// StringError generate string
func (c *Common) StringError(v map[string]string) string {
	var r string
	l := len(v)
	i := 1
	for k, v := range v {
		r += k + ":" + v
		if i != l {
			r += ";"
		}
		i++
	}
	return r
}
