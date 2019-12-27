package atomcommon_test

import (
	"testing"

	"github.com/SirMetathyst/atom"
	"github.com/SirMetathyst/atomcommon"
)

func TestSetLocalPosition2(t *testing.T) {

	// Setup
	atom.Reset()

	// Arrange
	id := atom.CreateEntity()
	inData1 := atomcommon.LocalPosition2Data{X: 10, Y: 10}

	// Act
	atomcommon.SetLocalPosition2(id, inData1)

	// Assert
	outData := atomcommon.LocalPosition2(id)
	if outData != inData1 {
		t.Errorf("assert: want %v, got %v", inData1, outData)
	}

	// Arrange
	inData2 := atomcommon.LocalPosition2Data{X: 10, Y: 10}

	// Act
	atomcommon.SetLocalPosition2(id, inData2)

	// Assert
	outData = atomcommon.LocalPosition2(id)
	if outData != inData2 {
		t.Errorf("assert: want %v, got %v", inData2, outData)
	}
}

func TestDeleteLocalPosition2(t *testing.T) {

	// Setup
	atom.Reset()

	// Arrange
	id := atom.CreateEntity()
	inData1 := atomcommon.LocalPosition2Data{X: 10, Y: 10}

	// Act
	atomcommon.SetLocalPosition2(id, inData1)

	// Assert
	v := atomcommon.HasLocalPosition2(id)
	if v != true {
		t.Errorf("assert: want %v, got %v", true, v)
	}

	// Act
	atomcommon.DeleteLocalPosition2(id)

	// Assert
	v = atomcommon.HasLocalPosition2(id)
	if v != false {
		t.Errorf("assert: want %v, got %v", false, v)
	}
}

func BenchmarkSetLocalPosition2(b *testing.B) {
	e := atom.NewEntityManager()
	c := atomcommon.NewLocalPosition2Component()
	ctx := e.RegisterComponent(atomcommon.LocalPosition2Key, c)
	c.SetContext(ctx)
	id := e.CreateEntity()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		atomcommon.SetLocalPosition2X(e, id, atomcommon.LocalPosition2Data{X: 10, Y: 10})
	}
}

func BenchmarkDeleteLocalPosition2(b *testing.B) {
	e := atom.NewEntityManager()
	c := atomcommon.NewLocalPosition2Component()
	ctx := e.RegisterComponent(atomcommon.LocalPosition2Key, c)
	c.SetContext(ctx)
	id := e.CreateEntity()
	data := atomcommon.LocalPosition2Data{X: 10, Y: 10}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		atomcommon.SetLocalPosition2X(e, id, data)
		atomcommon.DeleteLocalPosition2X(e, id)
	}
}

func BenchmarkLocalPosition2(b *testing.B) {
	e := atom.NewEntityManager()
	c := atomcommon.NewLocalPosition2Component()
	ctx := e.RegisterComponent(atomcommon.LocalPosition2Key, c)
	c.SetContext(ctx)
	id := e.CreateEntity()
	atomcommon.SetLocalPosition2X(e, id, atomcommon.LocalPosition2Data{X: 10, Y: 10})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		atomcommon.LocalPosition2X(e, id)
	}
}
