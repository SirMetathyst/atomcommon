package logic

import (
	"github.com/SirMetathyst/zinc"
	"github.com/SirMetathyst/zinckit"
)

// Position2System ...
type Position2System struct {
	group         zinc.G
	entityManager *zinc.EntityManager
}

// NewPosition2System ...
func NewPosition2System() *Position2System {
	return &Position2System{
		entityManager: zinc.Default(),
		group:         zinc.Default().Group(zinc.AllOf(zinckit.LocalPosition2Key, zinckit.Velocity2Key)),
	}
}

// NewPosition2SystemWith ...
func NewPosition2SystemWith(e *zinc.EntityManager) *Position2System {
	return &Position2System{
		entityManager: e,
		group:         e.Group(zinc.AllOf(zinckit.LocalPosition2Key, zinckit.Velocity2Key)),
	}
}

// Update ...
func (s Position2System) Update(dt float64) {
	for _, id := range s.group.Entities() {
		velocity := zinckit.Velocity2X(s.entityManager, id)
		position := zinckit.LocalPosition2X(s.entityManager, id)
		position.X += velocity.X * float32(dt)
		position.Y += velocity.Y * float32(dt)
		zinckit.SetLocalPosition2X(s.entityManager, id, position)
	}
}
