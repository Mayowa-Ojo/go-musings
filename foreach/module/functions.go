package module

func (c CustomArray) Foreach(callback Callback) {
	for i, el := range c {
		callback(el, i)
	}
}
