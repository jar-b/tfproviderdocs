package contents

import (
	"testing"
)

func TestCheckFunctionArgumentsSection(t *testing.T) {
	testCases := []struct {
		Name         string
		Path         string
		ProviderName string
		ExpectError  bool
	}{
		{
			Name:         "passing",
			Path:         "testdata/function_arguments/passing.md",
			ProviderName: "test",
		},
		{
			Name:         "missing heading",
			Path:         "testdata/function_arguments/missing_heading.md",
			ProviderName: "test",
			ExpectError:  true,
		},
		{
			Name:         "wrong heading level",
			Path:         "testdata/function_arguments/wrong_heading_level.md",
			ProviderName: "test",
			ExpectError:  true,
		},
		{
			Name:         "wrong heading text",
			Path:         "testdata/function_arguments/wrong_heading_text.md",
			ProviderName: "test",
			ExpectError:  true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			doc := NewDocument(testCase.Path, testCase.ProviderName)

			if err := doc.Parse(); err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			got := doc.checkFunctionArgumentsSection()

			if got == nil && testCase.ExpectError {
				t.Errorf("expected error, got no error")
			}

			if got != nil && !testCase.ExpectError {
				t.Errorf("expected no error, got error: %s", got)
			}
		})
	}
}
