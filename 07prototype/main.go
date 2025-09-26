package main

type Cloneable interface {
	Clone() Cloneable
}

type PrototypeManager struct {
	prototypes map[string]Cloneable
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{
		prototypes: make(map[string]Cloneable),
	}
}

func (p *PrototypeManager) SetPrototype(key string, prototype Cloneable) {
	if prototype == nil {
		return
	}
	p.prototypes[key] = prototype
}

func (p *PrototypeManager) GetPrototype(key string) Cloneable {
	if v, ok := p.prototypes[key]; ok {
		return v.Clone()
	}
	return nil
}

type Type1 struct {
	Name string
}

func (t *Type1) Clone() Cloneable {
	return t
}

type Type2 struct {
	Name string
}

func (t *Type2) Clone() Cloneable {
	return t
}

func main() {
	p := NewPrototypeManager()
	p.SetPrototype("type1", &Type1{Name: "dasdasdsad"})
	p1 := p.GetPrototype("type1")
	if _, ok := p1.(*Type1); ok {
		
	}
}
