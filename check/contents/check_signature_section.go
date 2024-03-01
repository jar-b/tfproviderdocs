package contents

import (
	"fmt"
	"strings"

	"github.com/bflad/tfproviderdocs/markdown"
)

type CheckSignatureSectionOptions struct {
	ExpectedCodeBlockLanguage string
}

func (d *Document) checkSignatureSection() error {
	checkOpts := &CheckSignatureSectionOptions{
		ExpectedCodeBlockLanguage: markdown.FencedCodeBlockLanguageTerraform,
	}

	if d.CheckOptions != nil && d.CheckOptions.SignatureSection != nil {
		checkOpts = d.CheckOptions.SignatureSection
	}

	section := d.Sections.Signature

	if section == nil {
		return fmt.Errorf("missing signature section: ## Signature")
	}

	heading := section.Heading

	if heading.Level != 2 {
		return fmt.Errorf("signature section heading level (%d) should be: 2", heading.Level)
	}

	headingText := string(heading.Text(d.source))
	expectedHeadingText := "Signature"

	if headingText != expectedHeadingText {
		return fmt.Errorf("signature section heading (%s) should be: %s", headingText, expectedHeadingText)
	}

	// CDKTF conversion will leave the original terraform code blocks if unsuccessful
	if checkOpts.ExpectedCodeBlockLanguage != markdown.FencedCodeBlockLanguageTerraform {
		return nil
	}

	for _, fencedCodeBlock := range section.FencedCodeBlocks {
		language := markdown.FencedCodeBlockLanguage(fencedCodeBlock, d.source)

		if language != checkOpts.ExpectedCodeBlockLanguage {
			return fmt.Errorf("signature section code block language (%s) should be: ```%s", language, checkOpts.ExpectedCodeBlockLanguage)
		}

		text := markdown.FencedCodeBlockText(fencedCodeBlock, d.source)
		functionName := strings.TrimPrefix(d.ResourceName, d.ProviderName+"_")

		if !strings.Contains(text, functionName) {
			return fmt.Errorf("signature section code block text should contain function name: %s", functionName)
		}
	}

	return nil
}
