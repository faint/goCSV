package goCSV

// Line ...
type Line struct {
	Keys   *Key
	Values []string
}

func createLine(k *Key, s []string) Line {
	line := Line{}
	line.Keys = k
	line.Values = s
	return line
}

// GetValueBy ...
func (line *Line) GetValueBy(key string) (string, bool) {
	n, result := line.Keys.GetIndex(key)
	if !result {
		return "", false
	}

	return line.Values[n], true
}

// GetValueByN ...
func (line *Line) GetValueByN(n int) (string, bool) {
	if line.Values[n] == "" {
		return "", false
	}
	return line.Values[n], true
}
