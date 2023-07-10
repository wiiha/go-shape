package shp

import "testing"

func TestBoxExtend(t *testing.T) {
	a := Box{-124.763068, 45.543541, -116.915989, 49.002494}
	b := Box{-92.888114, 42.49192, -86.805415, 47.080621}
	a.Extend(b)
	c := Box{-124.763068, 42.49192, -86.805415, 49.002494}
	if a.MinX != c.MinX {
		t.Errorf("a.MinX = %v, want %v", a.MinX, c.MinX)
	}
	if a.MinY != c.MinY {
		t.Errorf("a.MinY = %v, want %v", a.MinY, c.MinY)
	}
	if a.MaxX != c.MaxX {
		t.Errorf("a.MaxX = %v, want %v", a.MaxX, c.MaxX)
	}
	if a.MaxY != c.MaxY {
		t.Errorf("a.MaxY = %v, want %v", a.MaxY, c.MaxY)
	}
}

func TestPackShape(t *testing.T) {
	p := &Point{1, 2}
	b, err := PackShape(p)
	if err != nil {
		t.Fatalf("PackShape failed: %v", err)
	}
	if b == nil {
		t.Fatalf("PackShape returned nil")
	}
	t.Logf("PackShape returned %v", b)

	// Unpack it
	ns, err := UnpackShape(b)
	if err != nil {
		t.Fatalf("UnpackShape failed: %v", err)
	}
	if ns == nil {
		t.Fatalf("UnpackShape returned nil")
	}
	p2, ok := ns.(*Point)
	if !ok {
		t.Fatalf("UnpackShape returned %T, want *Point", ns)
	}
	if p2.X != p.X {
		t.Errorf("p2.X = %v, want %v", p2.X, p.X)
	}
	if p2.Y != p.Y {
		t.Errorf("p2.Y = %v, want %v", p2.Y, p.Y)
	}

}
