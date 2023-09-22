package cmd

import (
	"fmt"
	"testing"
)

func Test_newPathComputer_singlePath(t *testing.T) {
	pc := newPathComputer("")

	subj := "private.dp.edna.compatibility-value"
	pc.addSubject(subj)

	paths := pc.sortedPaths()

	expected := []string{"./dp/private", "./dp/private/edna", "./dp/private/edna/compatibility"}

	if fmt.Sprintf("%v", paths) != fmt.Sprintf("%v", expected) {
		t.Errorf("Expected %v, got %v", expected, paths)
	}
}

func Test_newPathComputer_multiPath(t *testing.T) {
	pc := newPathComputer("")

	pc.addSubject("private.dp.edna.compatibility-value")
	pc.addSubject("private.dp.edna.examples-value")
	pc.addSubject("private.dp.edna.compatibility-Publisher-Model-Person")
	pc.addSubject("public.kunde.mmm.verifiedmeteringvolumes-value")
	pc.addSubject("public.kunde.sesam.kunde-value")
	pc.addSubject("comp-test-oot")

	paths := pc.sortedPaths()

	expectedPatsh := `
./comp_test_oot
./dp/private
./dp/private/edna
./dp/private/edna/compatibility
./dp/private/edna/compatibility_publisher_model_person
./dp/private/edna/examples
./kunde/public
./kunde/public/mmm
./kunde/public/mmm/verifiedmeteringvolumes
./kunde/public/sesam
./kunde/public/sesam/kunde`

	ps := ""
	for _, path := range paths {
		ps = fmt.Sprintf("%s\n%s", ps, path)
	}

	if ps != expectedPatsh {
		t.Errorf("Expected %v, got %v", expectedPatsh, ps)
	}
}
