package atomcommon

import "github.com/SirMetathyst/atom"

// Position2Key ...
const Position2Key uint = 2789147080

// Position2Data ...
type Position2Data struct {
	X float32
	Y float32	
}

// Position2Component ...
type Position2Component struct {
	context atom.Context
	data map[atom.EntityID]Position2Data
}

// NewPosition2Component ...
func NewPosition2Component() *Position2Component {
	return &Position2Component{
		data: make(map[atom.EntityID]Position2Data),
	}
}

func init() {
	x := NewPosition2Component()
	context := atom.Default().RegisterComponent(Position2Key, x)
	x.context = context 
}

// EntityDeleted ...
func (c *Position2Component) EntityDeleted(id atom.EntityID) {
	delete(c.data, id)
}

// HasEntity ...
func (c *Position2Component) HasEntity(id atom.EntityID) bool {
	_, ok := c.data[id]
	return ok
}

// SetPosition2 ...
func (c *Position2Component) SetPosition2(id atom.EntityID, position2 Position2Data) {
	if c.context.HasEntity(id) {
		if c.HasEntity(id) {
			c.data[id] = position2
			c.context.ComponentUpdated(Position2Key, id)
		} else {
			c.data[id] = position2
			c.context.ComponentAdded(Position2Key, id)
		}
	}
}

// Position2 ...
func (c *Position2Component) Position2(id atom.EntityID) Position2Data {
	return c.data[id]
}

// DeletePosition2 ...
func (c *Position2Component) DeletePosition2(id atom.EntityID) {
	delete(c.data, id)
	c.context.ComponentDeleted(Position2Key, id)
}

// SetPosition2X ...
func SetPosition2X(e *atom.EntityManager, id atom.EntityID, position2 Position2Data) {
	v := e.Component(Position2Key)
	c := v.(*Position2Component)
	c.SetPosition2(id, position2)
}

// SetPosition2 ...
func SetPosition2(id atom.EntityID, position2 Position2Data) {
	SetPosition2X(atom.Default(), id, position2)
}

// Position2X ...
func Position2X(e *atom.EntityManager, id atom.EntityID) Position2Data {
	v := e.Component(Position2Key)
	c := v.(*Position2Component)
	return c.Position2(id)
}

// Position2 ...
func Position2(id atom.EntityID) Position2Data {
	return Position2X(atom.Default(), id)
}

// DeletePosition2X ...
func DeletePosition2X(e *atom.EntityManager, id atom.EntityID) {
	v := e.Component(Position2Key)
	c := v.(*Position2Component)
	c.DeletePosition2(id)
}

// DeletePosition2 ...
func DeletePosition2(id atom.EntityID) {
	DeletePosition2X(atom.Default(), id)
}

// HasPosition2X ...
func HasPosition2X(e *atom.EntityManager, id atom.EntityID) bool {
	v := e.Component(Position2Key)
	return v.HasEntity(id)
}

// HasPosition2 ...
func HasPosition2(id atom.EntityID) bool {
	return HasPosition2X(atom.Default(), id)
}