package contents

import (
	"fmt"
)

func (d *Document) checkFunctionArgumentsSection() error {
	section := d.Sections.FunctionArguments

	if section == nil {
		return fmt.Errorf("missing arguments section: ## Arguments")
	}

	heading := section.Heading

	if heading.Level != 2 {
		return fmt.Errorf("arguments section heading level (%d) should be: 2", heading.Level)
	}

	headingText := string(heading.Text(d.source))
	expectedHeadingText := "Arguments"

	if headingText != expectedHeadingText {
		return fmt.Errorf("arguments section heading (%s) should be: %s", headingText, expectedHeadingText)
	}

	return nil
}
