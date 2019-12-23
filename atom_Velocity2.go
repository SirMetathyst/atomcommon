package atomcommon

import "github.com/SirMetathyst/atom"

// Velocity2Key ...
const Velocity2Key uint = 2475506976

// Velocity2Data ...
type Velocity2Data struct {
	X float32
	Y float32	
}

// Velocity2Component ...
type Velocity2Component struct {
	context atom.Context
	data map[atom.EntityID]Velocity2Data
}

// NewVelocity2Component ...
func NewVelocity2Component() *Velocity2Component {
	return &Velocity2Component{
		data: make(map[atom.EntityID]Velocity2Data),
	}
}

func init() {
	x := NewVelocity2Component()
	context := atom.Default().RegisterComponent(Velocity2Key, x)
	x.context = context 
}

// EntityDeleted ...
func (c *Velocity2Component) EntityDeleted(id atom.EntityID) {
	delete(c.data, id)
}

// HasEntity ...
func (c *Velocity2Component) HasEntity(id atom.EntityID) bool {
	_, ok := c.data[id]
	return ok
}

// SetVelocity2 ...
func (c *Velocity2Component) SetVelocity2(id atom.EntityID, velocity2 Velocity2Data) {
	if c.context.HasEntity(id) {
		if c.HasEntity(id) {
			c.data[id] = velocity2
			c.context.ComponentUpdated(Velocity2Key, id)
		} else {
			c.data[id] = velocity2
			c.context.ComponentAdded(Velocity2Key, id)
		}
	}
}

// Velocity2 ...
func (c *Velocity2Component) Velocity2(id atom.EntityID) Velocity2Data {
	return c.data[id]
}

// DeleteVelocity2 ...
func (c *Velocity2Component) DeleteVelocity2(id atom.EntityID) {
	delete(c.data, id)
	c.context.ComponentDeleted(Velocity2Key, id)
}

// SetVelocity2X ...
func SetVelocity2X(e *atom.EntityManager, id atom.EntityID, velocity2 Velocity2Data) {
	v := e.Component(Velocity2Key)
	c := v.(*Velocity2Component)
	c.SetVelocity2(id, velocity2)
}

// SetVelocity2 ...
func SetVelocity2(id atom.EntityID, velocity2 Velocity2Data) {
	SetVelocity2X(atom.Default(), id, velocity2)
}

// Velocity2X ...
func Velocity2X(e *atom.EntityManager, id atom.EntityID) Velocity2Data {
	v := e.Component(Velocity2Key)
	c := v.(*Velocity2Component)
	return c.Velocity2(id)
}

// Velocity2 ...
func Velocity2(id atom.EntityID) Velocity2Data {
	return Velocity2X(atom.Default(), id)
}

// DeleteVelocity2X ...
func DeleteVelocity2X(e *atom.EntityManager, id atom.EntityID) {
	v := e.Component(Velocity2Key)
	c := v.(*Velocity2Component)
	c.DeleteVelocity2(id)
}

// DeleteVelocity2 ...
func DeleteVelocity2(id atom.EntityID) {
	DeleteVelocity2X(atom.Default(), id)
}

// HasVelocity2X ...
func HasVelocity2X(e *atom.EntityManager, id atom.EntityID) bool {
	v := e.Component(Velocity2Key)
	return v.HasEntity(id)
}

// HasVelocity2 ...
func HasVelocity2(id atom.EntityID) bool {
	return HasVelocity2X(atom.Default(), id)
}