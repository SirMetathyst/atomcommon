package logic

import (
	"github.com/SirMetathyst/atom"
	"github.com/SirMetathyst/atomcommon"
)

// Position2DSystem ...
type Position2DSystem struct {
	group         *atom.G
	entityManager *atom.EntityManager
}

// NewPosition2DSystem ...
func NewPosition2DSystem() *Position2DSystem {
	return &Position2DSystem{
		entityManager: atom.Default(),
		group:         atom.Default().Group(atom.AllOf(atomcommon.Position2Key, atomcommon.Velocity2Key)),
	}
}

// NewPosition2DSystemWith ...
func NewPosition2DSystemWith(e *atom.EntityManager) *Position2DSystem {
	return &Position2DSystem{
		entityManager: e,
		group:         e.Group(atom.AllOf(atomcommon.Position2Key, atomcommon.Velocity2Key)),
	}
}

// Update ...
func (s Position2DSystem) Update(dt float32) {
	for _, id := range s.group.Entities() {
		velocity := atomcommon.Velocity2X(s.entityManager, id)
		position := atomcommon.Position2X(s.entityManager, id)
		position.X += velocity.X * dt
		position.Y += velocity.Y * dt
		atomcommon.SetPosition2X(s.entityManager, id, position)
	}
}
