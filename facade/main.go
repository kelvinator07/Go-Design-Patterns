package main

import "fmt"

type Buffer struct {
	width, height int
	buffer []rune
}

func NewBuffer(width int, height int) *Buffer {
	return &Buffer{width, height, make([]rune, width * height)}
}

func (b *Buffer) At(index int) rune {
	return b.buffer[index]
}

type ViewPort struct {
	buffer *Buffer
	offset int
}

func NewViewPort(buffer *Buffer) *ViewPort {
	return &ViewPort{buffer, 0}
}

func (v *ViewPort) GetCharacterAt(index int) rune {
	return v.buffer.At(v.offset + index)
}

// facade
type Console struct {
	buffer []*Buffer
	viewPorts []*ViewPort
	offset int
}

func NewConsole() *Console{
	b := NewBuffer(200, 150)
	v := NewViewPort(b)
	return &Console{[]*Buffer{b}, []*ViewPort{v}, 0}
}

func (c *Console) GetCharacterAt(index int) rune {
	return c.viewPorts[0].GetCharacterAt(index)
}

func main() {
	c := NewConsole()
	u := c.GetCharacterAt(1)
	fmt.Println(u)
}
