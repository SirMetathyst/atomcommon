package logic

import (
	"github.com/SirMetathyst/atom"
	"github.com/SirMetathyst/atomkit"
)

// Position2System ...
type Position2System struct {
	group         atom.G
	entityManager *atom.EntityManager
}

// NewPosition2System ...
func NewPosition2System() *Position2System {
	return &Position2System{
		entityManager: atom.Default(),
		group:         atom.Default().Group(atom.AllOf(atomkit.LocalPosition2Key, atomkit.Velocity2Key)),
	}
}

// NewPosition2SystemWith ...
func NewPosition2SystemWith(e *atom.EntityManager) *Position2System {
	return &Position2System{
		entityManager: e,
		group:         e.Group(atom.AllOf(atomkit.LocalPosition2Key, atomkit.Velocity2Key)),
	}
}

// Update ...
func (s Position2System) Update(dt float64) {
	for _, id := range s.group.Entities() {
		velocity := atomkit.Velocity2X(s.entityManager, id)
		position := atomkit.LocalPosition2X(s.entityManager, id)
		position.X += velocity.X * float32(dt)
		position.Y += velocity.Y * float32(dt)
		atomkit.SetLocalPosition2X(s.entityManager, id, position)
	}
}
