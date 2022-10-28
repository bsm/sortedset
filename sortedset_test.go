package sortedset_test

import (
	"testing"

	"github.com/bsm/sortedset"
)

func seed(vv ...string) sortedset.Set[string] {
	if vv == nil {
		vv = []string{"b", "d", "f"}
	}

	set := sortedset.NewCap[string](len(vv))
	set = set.Add(vv...)
	return set
}

func TestSet_Len(t *testing.T) {
	set := seed()
	if exp, got := 3, set.Len(); exp != got {
		t.Errorf("expected %v, got %v", exp, got)
	}
}

func TestSet_Clear(t *testing.T) {
	set := seed()
	set = set.Clear()
	if exp, got := 0, set.Len(); exp != got {
		t.Errorf("expected %v, got %v", exp, got)
	}
}

func TestSet_Clone(t *testing.T) {
	set := seed()
	cln := set.Clone()
	set = set.Add("a")

	if exp, got := 4, set.Len(); exp != got {
		t.Errorf("expected %v, got %v", exp, got)
	}
	if exp, got := 3, cln.Len(); exp != got {
		t.Errorf("expected %v, got %v", exp, got)
	}
}

func TestSet_Add(t *testing.T) {
	set := seed()

	set = set.Add("c", "a")
	if exp, got := 5, set.Len(); exp != got {
		t.Errorf("expected %v, got %v", exp, got)
	}

	set = set.Add("b", "c").Add("d")
	if exp, got := 5, set.Len(); exp != got {
		t.Errorf("expected %v, got %v", exp, got)
	}
}

func TestSet_Delete(t *testing.T) {
	set := seed()

	set = set.Delete("c")
	if exp, got := 3, set.Len(); exp != got {
		t.Errorf("expected %v, got %v", exp, got)
	}

	set = set.Delete("b")
	if exp, got := 2, set.Len(); exp != got {
		t.Errorf("expected %v, got %v", exp, got)
	}

	set = set.Delete("b")
	if exp, got := 2, set.Len(); exp != got {
		t.Errorf("expected %v, got %v", exp, got)
	}
}

func TestSet_Has(t *testing.T) {
	set := seed()

	if exp, got := false, set.Has("a"); exp != got {
		t.Errorf("expected %v, got %v", exp, got)
	}
	if exp, got := true, set.Has("b"); exp != got {
		t.Errorf("expected %v, got %v", exp, got)
	}
	if exp, got := false, set.Has("c"); exp != got {
		t.Errorf("expected %v, got %v", exp, got)
	}
	if exp, got := true, set.Has("d"); exp != got {
		t.Errorf("expected %v, got %v", exp, got)
	}
}

func TestSet_Equal(t *testing.T) {
	set1 := seed()
	set2 := seed("d")
	if exp, got := false, set1.Equal(set2); exp != got {
		t.Errorf("expected %v, got %v", exp, got)
	}

	set2 = set2.Add("f", "b")
	if exp, got := true, set1.Equal(set2); exp != got {
		t.Errorf("expected %v, got %v", exp, got)
	}

	set2 = set2.Add("x")
	if exp, got := false, set1.Equal(set2); exp != got {
		t.Errorf("expected %v, got %v", exp, got)
	}
}

func TestSet_Intersects(t *testing.T) {
	set := seed()
	oth := sortedset.New[string]()

	if exp, got := false, set.Intersects(oth); exp != got {
		t.Errorf("expected %v, got %v", exp, got)
	}

	oth = oth.Add("c", "e")
	if exp, got := false, set.Intersects(oth); exp != got {
		t.Errorf("expected %v, got %v", exp, got)
	}

	oth = oth.Add("g", "d")
	if exp, got := true, set.Intersects(oth); exp != got {
		t.Errorf("expected %v, got %v", exp, got)
	}
}

func TestSet_Intersection(t *testing.T) {
	set1 := seed()
	set2 := seed("b", "c", "d", "x")

	res := sortedset.New[string]()
	res = res.Intersection(set1, set2)

	if exp, got := (sortedset.Set[string]{"b", "d"}), res; !exp.Equal(got) {
		t.Errorf("expected %v, got %v", exp, got)
	}
}

func TestSet_IntersectionWith(t *testing.T) {
	set1 := seed()
	set2 := seed("b", "c", "d", "x")
	set1 = set1.IntersectionWith(set2)

	if exp, got := (sortedset.Set[string]{"b", "d"}), set1; !exp.Equal(got) {
		t.Errorf("expected %v, got %v", exp, got)
	}
}

func TestSet_Union(t *testing.T) {
	set1 := seed()
	set2 := seed("b", "c", "d", "x")

	res := sortedset.New[string]()
	res = res.Union(set1, set2)

	if exp, got := (sortedset.Set[string]{"b", "c", "d", "f", "x"}), res; !exp.Equal(got) {
		t.Errorf("expected %v, got %v", exp, got)
	}
}

func TestSet_UnionWith(t *testing.T) {
	set1 := seed()
	set2 := seed("b", "c", "d", "x")
	set1 = set1.UnionWith(set2)

	if exp, got := (sortedset.Set[string]{"b", "c", "d", "f", "x"}), set1; !exp.Equal(got) {
		t.Errorf("expected %v, got %v", exp, got)
	}
}
