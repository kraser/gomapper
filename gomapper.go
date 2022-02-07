// gomapper
package gomapper

import (
	"sort"
	"strings"
)

type Mapper map[string][]string

func Init() *Mapper {
	m := new(Mapper)
	*m = make(map[string][]string)
	return m
}

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns
// the empty string. To access multiple values, use the map
// directly.
func (m Mapper) Get(key string) string {
	if m == nil {
		return ""
	}
	vs := m[key]
	if len(vs) == 0 {
		return ""
	}
	return vs[0]
}

// Set sets the key to value. It replaces any existing
// values.
func (m Mapper) Set(key, value string) {
	m[key] = []string{value}
}

// Add adds the value to key. It appends to any existing
// values associated with key.
func (m Mapper) Add(key, value string) {
	m[key] = append(m[key], value)
}

// Del deletes the values associated with key.
func (m Mapper) Del(key string) {
	delete(m, key)
}

// Has checks whether a given key is set.
func (m Mapper) Has(key string) bool {
	_, ok := m[key]
	return ok
}

// Encode encodes the values into ``URL encoded'' form
// ("bar=baz&foo=quux") sorted by key.
func (m Mapper) Implode() string {
	if m == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := m[k]
		//keyEscaped := QueryEscape(k)
		for _, m := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(m)
		}
	}
	return buf.String()
}
