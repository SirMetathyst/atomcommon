package logic

import (
	"github.com/SirMetathyst/atom"
	"github.com/SirMetathyst/atomcommon"
)

// Position2System ...
type Position2System struct {
	group         atom.Grouper
	entityManager *atom.EntityManager
}

// NewPosition2System ...
func NewPosition2System() *Position2System {
	return &Position2System{
		entityManager: atom.Default(),
		group:         atom.Default().Group(atom.AllOf(atomcommon.LocalPosition2Key, atomcommon.Velocity2Key)),
	}
}

// NewPosition2SystemWith ...
func NewPosition2SystemWith(e *atom.EntityManager) *Position2System {
	return &Position2System{
		entityManager: e,
		group:         e.Group(atom.AllOf(atomcommon.LocalPosition2Key, atomcommon.Velocity2Key)),
	}
}

// Update ...
func (s Position2System) Update(dt float32) {
	for _, id := range s.group.Entities() {
		velocity := atomcommon.Velocity2X(s.entityManager, id)
		position := atomcommon.LocalPosition2X(s.entityManager, id)
		position.X += velocity.X * dt
		position.Y += velocity.Y * dt
		atomcommon.SetLocalPosition2X(s.entityManager, id, position)
	}
}
