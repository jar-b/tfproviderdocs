package contents

import (
	"testing"
)

func TestCheckSignatureSection(t *testing.T) {
	testCases := []struct {
		Name         string
		Path         string
		ProviderName string
		ExpectError  bool
	}{
		{
			Name:         "passing",
			Path:         "testdata/signature/passing.md",
			ProviderName: "test",
		},
		{
			Name:         "missing code block language",
			Path:         "testdata/signature/missing_code_block_language.md",
			ProviderName: "test",
			ExpectError:  true,
		},
		{
			Name:         "missing function path",
			Path:         "testdata/signature/missing_function_path.md",
			ProviderName: "test",
			ExpectError:  true,
		},
		{
			Name:         "missing heading",
			Path:         "testdata/signature/missing_heading.md",
			ProviderName: "test",
			ExpectError:  true,
		},
		{
			Name:         "wrong heading level",
			Path:         "testdata/signature/wrong_heading_level.md",
			ProviderName: "test",
			ExpectError:  true,
		},
		{
			Name:         "wrong heading text",
			Path:         "testdata/signature/wrong_heading_text.md",
			ProviderName: "test",
			ExpectError:  true,
		},
		{
			Name:         "wrong code block language",
			Path:         "testdata/signature/wrong_code_block_language.md",
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

			got := doc.checkSignatureSection()

			if got == nil && testCase.ExpectError {
				t.Errorf("expected error, got no error")
			}

			if got != nil && !testCase.ExpectError {
				t.Errorf("expected no error, got error: %s", got)
			}
		})
	}
}
