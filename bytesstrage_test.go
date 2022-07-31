package bs_test

import (
	"fmt"
	"testing"
	"text/template"

	"github.com/nil2nekoni/go-bytesstorage"
)

func Example() {
	w := bs.NewBytesStorage()

	tmpl, _ := template.New("test").Parse("{{.Id}}: {{.Score}}")
	tmpl.Execute(w, map[string]string{"Id": "abc", "Score": "100"})

	fmt.Println(string(w.Load()))
	// Output:
	// abc: 100
}


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
