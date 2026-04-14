package htmlcss

import (
	"fmt"
	"strings"

	xhtml "golang.org/x/net/html"

	"github.com/wnn-dev/contributions-analysis/objects"
)

// ApprovalThreshold is the minimum percentage to approve a submission
const ApprovalThreshold = 70.0

// Analyze parses and evaluates an HTML/CSS submission against the article.html template.
// Total score: 100 points.
func Analyze(htmlContent string) (*objects.HtmlCssAnalysisReport, error) {
	doc, err := xhtml.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	var passed []objects.CheckResult
	var failed []objects.CheckResult
	maxScore := 0

	run := func(r objects.CheckResult) {
		maxScore += r.MaxPoints
		if r.Passed {
			passed = append(passed, r)
		} else {
			r.Points = 0
			failed = append(failed, r)
		}
	}

	// --- Structure (40 pts) ---
	run(checkArticleWrapper(doc))
	run(checkH3Tag(doc))
	run(checkParagraph(doc))
	run(checkH4Count(doc))
	run(checkContainerDiv(doc))
	run(checkSocialContainer(doc))
	run(checkStructureOrder(doc))

	// --- Badge rules (25 pts) ---
	run(checkBadgeCount(doc))
	run(checkBadgeBackgroundColor(doc))
	run(checkBadgeTextColor(doc))

	// --- Social link rules (25 pts) ---
	run(checkSocialLinkCount(doc))
	run(checkSocialLinkHref(doc))
	run(checkSocialLinkTarget(doc))

	// --- CSS checks (10 pts) ---
	run(checkHasStyleBlock(htmlContent))
	run(checkBadgeCss(htmlContent))
	run(checkSocialLinkCss(htmlContent))

	totalScore := 0
	for _, c := range passed {
		totalScore += c.Points
	}

	percentage := float64(totalScore) / float64(maxScore) * 100

	return &objects.HtmlCssAnalysisReport{
		Score:        float64(totalScore),
		MaxScore:     float64(maxScore),
		Percentage:   percentage,
		Approved:     percentage >= ApprovalThreshold,
		PassedChecks: passed,
		FailedChecks: failed,
	}, nil
}

// ---------------------------------------------------------------------------
// Structure checks
// ---------------------------------------------------------------------------

func checkArticleWrapper(doc *xhtml.Node) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "article-wrapper",
		MaxPoints: 5,
		Expected:  "<article> element as the root wrapper",
	}
	if find(doc, func(n *xhtml.Node) bool { return isElement(n, "article") }) != nil {
		r.Passed = true
		r.Points = 5
		r.Actual = "Found <article> wrapper"
	} else {
		r.Actual = "No <article> element found"
		r.Diff = "Wrap all content inside an <article> tag"
	}
	return r
}

func checkH3Tag(doc *xhtml.Node) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "h3-username",
		MaxPoints: 5,
		Expected:  "<h3> element for the username",
	}
	if find(doc, func(n *xhtml.Node) bool { return isElement(n, "h3") }) != nil {
		r.Passed = true
		r.Points = 5
		r.Actual = "Found <h3> element"
	} else {
		r.Actual = "No <h3> element found"
		r.Diff = "Add an <h3> element with the username"
	}
	return r
}

func checkParagraph(doc *xhtml.Node) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "paragraph-bio",
		MaxPoints: 5,
		Expected:  "<p> element for the bio",
	}
	if find(doc, func(n *xhtml.Node) bool { return isElement(n, "p") }) != nil {
		r.Passed = true
		r.Points = 5
		r.Actual = "Found <p> element"
	} else {
		r.Actual = "No <p> element found"
		r.Diff = "Add a <p> element for the bio section"
	}
	return r
}

func checkH4Count(doc *xhtml.Node) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "h4-count",
		MaxPoints: 10,
		Expected:  "Exactly 2 <h4> section titles",
	}
	h4s := findAll(doc, func(n *xhtml.Node) bool { return isElement(n, "h4") })
	count := len(h4s)
	r.Actual = fmt.Sprintf("Found %d <h4> element(s)", count)
	if count == 2 {
		r.Passed = true
		r.Points = 10
	} else {
		r.Diff = fmt.Sprintf("Expected exactly 2 <h4> elements, found %d", count)
	}
	return r
}

func checkContainerDiv(doc *xhtml.Node) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "container-div",
		MaxPoints: 5,
		Expected:  `<section class="container"> wrapping the language badges`,
	}
	if find(doc, func(n *xhtml.Node) bool {
		return (isElement(n, "section") || isElement(n, "div")) && hasClass(n, "container")
	}) != nil {
		r.Passed = true
		r.Points = 5
		r.Actual = "Found .container element"
	} else {
		r.Actual = "No .container element found"
		r.Diff = `Add a <section class="container"> wrapping the badges`
	}
	return r
}

func checkSocialContainer(doc *xhtml.Node) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "social-container",
		MaxPoints: 5,
		Expected:  `<section class="social-container"> wrapping social links`,
	}
	if find(doc, func(n *xhtml.Node) bool {
		return (isElement(n, "section") || isElement(n, "div")) && hasClass(n, "social-container")
	}) != nil {
		r.Passed = true
		r.Points = 5
		r.Actual = "Found .social-container element"
	} else {
		r.Actual = "No .social-container element found"
		r.Diff = `Add a <section class="social-container"> wrapping the social links`
	}
	return r
}

func checkStructureOrder(doc *xhtml.Node) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "structure-order",
		MaxPoints: 5,
		Expected:  "h3 → p → h4 → .container → h4 → .social-container",
	}

	article := find(doc, func(n *xhtml.Node) bool { return isElement(n, "article") })
	if article == nil {
		r.Actual = "No <article> element to inspect"
		r.Diff = "Wrap all content in <article> first"
		return r
	}

	var order []string
	for c := article.FirstChild; c != nil; c = c.NextSibling {
		if c.Type != xhtml.ElementNode {
			continue
		}
		switch {
		case c.Data == "h3":
			order = append(order, "h3")
		case c.Data == "p":
			order = append(order, "p")
		case c.Data == "h4":
			order = append(order, "h4")
		case (c.Data == "section" || c.Data == "div") && hasClass(c, "container"):
			order = append(order, ".container")
		case (c.Data == "section" || c.Data == "div") && hasClass(c, "social-container"):
			order = append(order, ".social-container")
		}
	}

	expected := []string{"h3", "p", "h4", ".container", "h4", ".social-container"}
	r.Actual = strings.Join(order, " → ")

	match := len(order) == len(expected)
	if match {
		for i, v := range order {
			if v != expected[i] {
				match = false
				break
			}
		}
	}

	if match {
		r.Passed = true
		r.Points = 5
	} else {
		r.Diff = fmt.Sprintf("Expected: %s — Got: %s", strings.Join(expected, " → "), r.Actual)
	}
	return r
}

// ---------------------------------------------------------------------------
// Badge rules
// ---------------------------------------------------------------------------

func checkBadgeCount(doc *xhtml.Node) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "badge-count",
		MaxPoints: 15,
		Expected:  "Exactly 2 .badge elements inside .container (2 languages only)",
	}

	container := find(doc, func(n *xhtml.Node) bool {
		return (isElement(n, "section") || isElement(n, "div")) && hasClass(n, "container")
	})
	if container == nil {
		r.Actual = "No .container found"
		r.Diff = "Add a .container element first"
		return r
	}

	badges := findAll(container, func(n *xhtml.Node) bool {
		return isElement(n, "div") && hasClass(n, "badge")
	})
	count := len(badges)
	r.Actual = fmt.Sprintf("Found %d badge(s)", count)

	if count == 2 {
		r.Passed = true
		r.Points = 15
	} else {
		r.Diff = fmt.Sprintf("Expected exactly 2 badges, found %d — only 2 programming languages are allowed", count)
	}
	return r
}

func checkBadgeBackgroundColor(doc *xhtml.Node) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "badge-background-color",
		MaxPoints: 5,
		Expected:  "Each badge has inline background-color style attribute",
	}

	container := find(doc, func(n *xhtml.Node) bool {
		return (isElement(n, "section") || isElement(n, "div")) && hasClass(n, "container")
	})
	if container == nil {
		r.Actual = "No .container found"
		r.Diff = "Add a .container element first"
		return r
	}

	badges := findAll(container, func(n *xhtml.Node) bool {
		return isElement(n, "div") && hasClass(n, "badge")
	})

	var missing []string
	for i, badge := range badges {
		if !strings.Contains(attr(badge, "style"), "background-color") {
			missing = append(missing, fmt.Sprintf("badge %d", i+1))
		}
	}

	if len(missing) == 0 {
		r.Passed = true
		r.Points = 5
		r.Actual = "All badges have background-color"
	} else {
		r.Actual = fmt.Sprintf("Missing background-color in: %s", strings.Join(missing, ", "))
		r.Diff = "Add background-color as inline style to each .badge element"
	}
	return r
}

func checkBadgeTextColor(doc *xhtml.Node) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "badge-text-color",
		MaxPoints: 5,
		Expected:  "Each badge has inline color style attribute",
	}

	container := find(doc, func(n *xhtml.Node) bool {
		return (isElement(n, "section") || isElement(n, "div")) && hasClass(n, "container")
	})
	if container == nil {
		r.Actual = "No .container found"
		r.Diff = "Add a .container element first"
		return r
	}

	badges := findAll(container, func(n *xhtml.Node) bool {
		return isElement(n, "div") && hasClass(n, "badge")
	})

	var missing []string
	for i, badge := range badges {
		// strip background-color so we don't confuse it with plain "color"
		styleWithoutBg := strings.ReplaceAll(attr(badge, "style"), "background-color", "")
		if !strings.Contains(styleWithoutBg, "color") {
			missing = append(missing, fmt.Sprintf("badge %d", i+1))
		}
	}

	if len(missing) == 0 {
		r.Passed = true
		r.Points = 5
		r.Actual = "All badges have text color"
	} else {
		r.Actual = fmt.Sprintf("Missing color in: %s", strings.Join(missing, ", "))
		r.Diff = "Add color as inline style to each .badge element"
	}
	return r
}

// ---------------------------------------------------------------------------
// Social link rules
// ---------------------------------------------------------------------------

func checkSocialLinkCount(doc *xhtml.Node) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "social-link-count",
		MaxPoints: 15,
		Expected:  "Exactly 2 .social-link anchor elements (2 social links only)",
	}

	social := find(doc, func(n *xhtml.Node) bool {
		return (isElement(n, "section") || isElement(n, "div")) && hasClass(n, "social-container")
	})
	if social == nil {
		r.Actual = "No .social-container found"
		r.Diff = "Add a .social-container element first"
		return r
	}

	links := findAll(social, func(n *xhtml.Node) bool {
		return isElement(n, "a") && hasClass(n, "social-link")
	})
	count := len(links)
	r.Actual = fmt.Sprintf("Found %d social link(s)", count)

	if count == 2 {
		r.Passed = true
		r.Points = 15
	} else {
		r.Diff = fmt.Sprintf("Expected exactly 2 social links, found %d — only 2 are allowed", count)
	}
	return r
}

func checkSocialLinkHref(doc *xhtml.Node) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "social-link-href",
		MaxPoints: 5,
		Expected:  "Each .social-link has a valid non-empty href",
	}

	links := findAll(doc, func(n *xhtml.Node) bool {
		return isElement(n, "a") && hasClass(n, "social-link")
	})

	var missing []string
	for i, link := range links {
		h := attr(link, "href")
		if h == "" || h == "#" {
			missing = append(missing, fmt.Sprintf("link %d", i+1))
		}
	}

	if len(missing) == 0 {
		r.Passed = true
		r.Points = 5
		r.Actual = "All social links have valid href"
	} else {
		r.Actual = fmt.Sprintf("Missing or empty href in: %s", strings.Join(missing, ", "))
		r.Diff = "Add a real URL as href to each .social-link element"
	}
	return r
}

func checkSocialLinkTarget(doc *xhtml.Node) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "social-link-target",
		MaxPoints: 5,
		Expected:  `Each .social-link has target="_blank"`,
	}

	links := findAll(doc, func(n *xhtml.Node) bool {
		return isElement(n, "a") && hasClass(n, "social-link")
	})

	var missing []string
	for i, link := range links {
		if attr(link, "target") != "_blank" {
			missing = append(missing, fmt.Sprintf("link %d", i+1))
		}
	}

	if len(missing) == 0 {
		r.Passed = true
		r.Points = 5
		r.Actual = `All social links have target="_blank"`
	} else {
		r.Actual = fmt.Sprintf(`Missing target="_blank" in: %s`, strings.Join(missing, ", "))
		r.Diff = `Add target="_blank" to each .social-link so links open in a new tab`
	}
	return r
}

// ---------------------------------------------------------------------------
// CSS checks
// ---------------------------------------------------------------------------

func checkHasStyleBlock(htmlContent string) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "has-style-block",
		MaxPoints: 5,
		Expected:  "A <style> block containing CSS rules",
	}
	lower := strings.ToLower(htmlContent)
	if strings.Contains(lower, "<style") && strings.Contains(lower, "</style>") {
		r.Passed = true
		r.Points = 5
		r.Actual = "Found <style> block"
	} else {
		r.Actual = "No <style> block found"
		r.Diff = "Add a <style> block with your CSS rules inside the HTML"
	}
	return r
}

func checkBadgeCss(htmlContent string) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "badge-css",
		MaxPoints: 3,
		Expected:  ".badge CSS rule with padding and border-radius",
	}
	lower := strings.ToLower(htmlContent)
	hasBadgeRule := strings.Contains(lower, ".badge")
	hasPadding := strings.Contains(lower, "padding")
	hasBorderRadius := strings.Contains(lower, "border-radius")

	if hasBadgeRule && hasPadding && hasBorderRadius {
		r.Passed = true
		r.Points = 3
		r.Actual = ".badge styled with padding and border-radius"
	} else {
		var missing []string
		if !hasBadgeRule {
			missing = append(missing, ".badge rule")
		}
		if !hasPadding {
			missing = append(missing, "padding")
		}
		if !hasBorderRadius {
			missing = append(missing, "border-radius")
		}
		r.Actual = fmt.Sprintf("Missing: %s", strings.Join(missing, ", "))
		r.Diff = "Add .badge CSS rule with padding and border-radius properties"
	}
	return r
}

func checkSocialLinkCss(htmlContent string) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "social-link-css",
		MaxPoints: 2,
		Expected:  ".social-link CSS rule with display: flex",
	}
	lower := strings.ToLower(htmlContent)
	if strings.Contains(lower, ".social-link") && strings.Contains(lower, "display") && strings.Contains(lower, "flex") {
		r.Passed = true
		r.Points = 2
		r.Actual = ".social-link styled with display: flex"
	} else {
		r.Actual = "Missing .social-link CSS with display: flex"
		r.Diff = "Add .social-link CSS rule with display: flex"
	}
	return r
}

// ---------------------------------------------------------------------------
// HTML tree helpers
// ---------------------------------------------------------------------------

func find(n *xhtml.Node, match func(*xhtml.Node) bool) *xhtml.Node {
	if match(n) {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if result := find(c, match); result != nil {
			return result
		}
	}
	return nil
}

func findAll(n *xhtml.Node, match func(*xhtml.Node) bool) []*xhtml.Node {
	var results []*xhtml.Node
	if match(n) {
		results = append(results, n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		results = append(results, findAll(c, match)...)
	}
	return results
}

func isElement(n *xhtml.Node, tag string) bool {
	return n.Type == xhtml.ElementNode && n.Data == tag
}

func hasClass(n *xhtml.Node, class string) bool {
	for _, c := range strings.Fields(attr(n, "class")) {
		if c == class {
			return true
		}
	}
	return false
}

func attr(n *xhtml.Node, key string) string {
	for _, a := range n.Attr {
		if a.Key == key {
			return a.Val
		}
	}
	return ""
}
