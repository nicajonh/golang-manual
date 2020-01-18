
package main
import (
"fmt"
)
type Op interface {
	getName() string
}
type A struct {
}
type B struct {
}
type Factory struct {
}

func (a *A) getName() string {
	return "A"
}
func (b *B) getName() string {
	return "B"
}
func (f *Factory) create(name string) Op {
	switch name {
	case `a`:
		return new(A)
	case `b`:
		return new(B)
	default:
		panic(`name not exists`)
	}
	return nil
}
func main() {
	var f = new(Factory)
	p := f.create(`a`)
	fmt.Println(p.getName())
	p = f.create(`b`)
	fmt.Println(p.getName())
}
