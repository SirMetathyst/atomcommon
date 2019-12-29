package atomkit_test

import (
	"testing"

	"github.com/SirMetathyst/atom"
	"github.com/SirMetathyst/atomkit"
)

func TestSetLocalPosition2(t *testing.T) {

	// Setup
	atom.Reset()

	// Arrange
	id := atom.CreateEntity()
	inData1 := atomkit.LocalPosition2Data{X: 10, Y: 10}

	// Act
	atomkit.SetLocalPosition2(id, inData1)

	// Assert
	outData := atomkit.LocalPosition2(id)
	if outData != inData1 {
		t.Errorf("assert: want %v, got %v", inData1, outData)
	}

	// Arrange
	inData2 := atomkit.LocalPosition2Data{X: 10, Y: 10}

	// Act
	atomkit.SetLocalPosition2(id, inData2)

	// Assert
	outData = atomkit.LocalPosition2(id)
	if outData != inData2 {
		t.Errorf("assert: want %v, got %v", inData2, outData)
	}
}

func TestDeleteLocalPosition2(t *testing.T) {

	// Setup
	atom.Reset()

	// Arrange
	id := atom.CreateEntity()
	inData1 := atomkit.LocalPosition2Data{X: 10, Y: 10}

	// Act
	atomkit.SetLocalPosition2(id, inData1)

	// Assert
	v := atomkit.HasLocalPosition2(id)
	if v != true {
		t.Errorf("assert: want %v, got %v", true, v)
	}

	// Act
	atomkit.DeleteLocalPosition2(id)

	// Assert
	v = atomkit.HasLocalPosition2(id)
	if v != false {
		t.Errorf("assert: want %v, got %v", false, v)
	}
}

func BenchmarkSetLocalPosition2(b *testing.B) {
	e := atom.NewEntityManager()
	c := atomkit.NewLocalPosition2Component()
	ctx := e.RegisterComponent(atomkit.LocalPosition2Key, c)
	c.SetContext(ctx)
	id := e.CreateEntity()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		atomkit.SetLocalPosition2X(e, id, atomkit.LocalPosition2Data{X: 10, Y: 10})
	}
}

func BenchmarkDeleteLocalPosition2(b *testing.B) {
	e := atom.NewEntityManager()
	c := atomkit.NewLocalPosition2Component()
	ctx := e.RegisterComponent(atomkit.LocalPosition2Key, c)
	c.SetContext(ctx)
	id := e.CreateEntity()
	data := atomkit.LocalPosition2Data{X: 10, Y: 10}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		atomkit.SetLocalPosition2X(e, id, data)
		atomkit.DeleteLocalPosition2X(e, id)
	}
}

func BenchmarkLocalPosition2(b *testing.B) {
	e := atom.NewEntityManager()
	c := atomkit.NewLocalPosition2Component()
	ctx := e.RegisterComponent(atomkit.LocalPosition2Key, c)
	c.SetContext(ctx)
	id := e.CreateEntity()
	atomkit.SetLocalPosition2X(e, id, atomkit.LocalPosition2Data{X: 10, Y: 10})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		atomkit.LocalPosition2X(e, id)
	}
}
