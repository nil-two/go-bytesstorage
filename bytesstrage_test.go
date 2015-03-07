package bs_test

import (
	"fmt"
	"testing"
	"text/template"

	"github.com/kusabashira/go-bytesstorage"
)

func TestLoadSimple(t *testing.T) {
	w := bs.NewBytesStorage()
	fmt.Fprint(w, "abc")

	actual := string(w.Load())
	expect := "abc"
	if actual != expect {
		t.Fatalf("Expected %v, but %v:", expect, actual)
	}
}

func TestLoadWithTextTemplate(t *testing.T) {
	w := bs.NewBytesStorage()
	tmpl, err := template.New("test").Parse("{{.Id}}: {{.Score}}")
	if err != nil {
		t.Fatalf("Mistook the use of the text/template %v:", err)
	}
	err = tmpl.Execute(w, map[string]string{"Id": "abc", "Score": "100"})
	if err != nil {
		t.Fatalf("Mistook the use of the text/template %v:", err)
	}

	actual := string(w.Load())
	expect := "abc: 100"
	if actual != expect {
		t.Fatalf("Expected %v, but %v:", expect, actual)
	}
}
