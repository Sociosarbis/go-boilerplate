package editor

type TextEditor struct {
	s []byte
	i int
}

func Constructor() TextEditor {
	return TextEditor{}
}

func (this *TextEditor) AddText(text string) {
	n := len(text)
	this.s = append(this.s, make([]byte, n)...)
	copy(this.s[this.i+n:], this.s[this.i:])
	copy(this.s[this.i:], text)
	this.i += n
}

func (this *TextEditor) DeleteText(k int) int {
	if k > this.i {
		k = this.i
	}
	copy(this.s[this.i-k:], this.s[this.i:])
	this.s = this.s[:len(this.s)-k]
	this.i -= k
	return k
}

func (this *TextEditor) CursorLeft(k int) string {
	if k > this.i {
		k = this.i
	}
	this.i -= k
	if this.i > 10 {
		return string(this.s[this.i-10 : this.i])
	}
	return string(this.s[:this.i])
}

func (this *TextEditor) CursorRight(k int) string {
	if this.i+k > len(this.s) {
		k = len(this.s) - this.i
	}
	this.i += k
	if this.i > 10 {
		return string(this.s[this.i-10 : this.i])
	}
	return string(this.s[:this.i])
}
