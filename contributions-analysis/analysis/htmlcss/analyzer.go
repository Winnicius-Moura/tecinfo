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
//
// Scoring breakdown (total: 100 pts):
//
//	Structure & Semantics — 38 pts
//	Badge rules           — 22 pts
//	Social link rules     — 22 pts
//	CSS fidelity          — 18 pts
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

	// --- Structure & Semantics (38 pts) ---
	run(checkArticleWrapper(doc))         // 5
	run(checkH3NonEmpty(doc))             // 5
	run(checkParagraph(doc))              // 3
	run(checkH4TitlesText(doc))           // 8
	run(checkSectionContainer(doc))       // 6
	run(checkSectionSocialContainer(doc)) // 6
	run(checkStructureOrder(doc))         // 5

	// --- Badge rules (22 pts) ---
	run(checkBadgeCount(doc))           // 9
	run(checkBadgeBackgroundColor(doc)) // 7
	run(checkBadgeTextColor(doc))       // 6

	// --- Social link rules (22 pts) ---
	run(checkSocialLinkCount(doc))  // 9
	run(checkSocialLinkHref(doc))   // 4
	run(checkSocialLinkTarget(doc)) // 5
	run(checkSocialLinkIcon(doc))   // 4

	// --- CSS fidelity (18 pts) ---
	run(checkHasStyleBlock(htmlContent)) // 3
	run(checkArticleCss(htmlContent))    // 3
	run(checkBadgeCss(htmlContent))      // 4
	run(checkContainerCss(htmlContent))  // 3
	run(checkSocialLinkCss(htmlContent)) // 5

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
// Structure & Semantics checks
// ---------------------------------------------------------------------------

func checkArticleWrapper(doc *xhtml.Node) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "article-wrapper",
		MaxPoints: 5,
		Expected:  "<article> as the root semantic element wrapping all content",
	}
	if find(doc, func(n *xhtml.Node) bool { return isElement(n, "article") }) != nil {
		r.Passed = true
		r.Points = 5
		r.Actual = "Found <article> wrapper"
	} else {
		r.Actual = "No <article> element found"
		r.Diff = "Wrap all card content inside an <article> tag — it is the required root element"
	}
	return r
}

// checkH3NonEmpty verifies that an <h3> element exists inside <article> and contains non-empty text.
func checkH3NonEmpty(doc *xhtml.Node) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "h3-username",
		MaxPoints: 5,
		Expected:  "<h3> inside <article> with a non-empty username as text content",
	}
	h3 := find(doc, func(n *xhtml.Node) bool { return isElement(n, "h3") })
	if h3 == nil {
		r.Actual = "No <h3> element found"
		r.Diff = "Add an <h3> element containing your username"
		return r
	}
	name := strings.TrimSpace(textContent(h3))
	if name == "" {
		r.Actual = "<h3> found but it has no text content"
		r.Diff = "Fill the <h3> element with your actual username"
		return r
	}
	r.Passed = true
	r.Points = 5
	r.Actual = fmt.Sprintf("<h3> found with content: %q", name)
	return r
}

func checkParagraph(doc *xhtml.Node) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "paragraph-bio",
		MaxPoints: 3,
		Expected:  "<p> element present for the bio (content is optional)",
	}
	if find(doc, func(n *xhtml.Node) bool { return isElement(n, "p") }) != nil {
		r.Passed = true
		r.Points = 3
		r.Actual = "Found <p> element"
	} else {
		r.Actual = "No <p> element found"
		r.Diff = "Add a <p> element after <h3> — even if you leave it empty, the tag must be present"
	}
	return r
}

// checkH4TitlesText verifies that both <h4> elements exist with their exact required text.
func checkH4TitlesText(doc *xhtml.Node) objects.CheckResult {
	const title1 = "Programming languages I use"
	const title2 = "Social Links"
	r := objects.CheckResult{
		Rule:      "h4-titles-text",
		MaxPoints: 8,
		Expected:  fmt.Sprintf(`Exactly 2 <h4> elements with the exact text %q and %q`, title1, title2),
	}

	h4s := findAll(doc, func(n *xhtml.Node) bool { return isElement(n, "h4") })
	if len(h4s) != 2 {
		r.Actual = fmt.Sprintf("Found %d <h4> element(s), expected exactly 2", len(h4s))
		r.Diff = fmt.Sprintf("Use exactly 2 <h4> headings with the specified text")
		return r
	}

	t1 := strings.TrimSpace(textContent(h4s[0]))
	t2 := strings.TrimSpace(textContent(h4s[1]))

	if t1 == title1 && t2 == title2 {
		r.Passed = true
		r.Points = 8
		r.Actual = fmt.Sprintf("Both <h4> titles match: %q and %q", t1, t2)
	} else {
		r.Actual = fmt.Sprintf("Got %q and %q", t1, t2)
		r.Diff = fmt.Sprintf("The text must be exactly %q (first) and %q (second) — check for typos, extra spaces, or wrong language", title1, title2)
	}
	return r
}

// checkSectionContainer enforces the use of <section> (not <div>) with class="container".
func checkSectionContainer(doc *xhtml.Node) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "section-container",
		MaxPoints: 6,
		Expected:  `<section class="container"> — must use the <section> semantic tag, not <div>`,
	}
	// Fail if they used a div with the right class — semantic tag is required
	if find(doc, func(n *xhtml.Node) bool {
		return isElement(n, "section") && hasClass(n, "container")
	}) != nil {
		r.Passed = true
		r.Points = 6
		r.Actual = `Found <section class="container">`
	} else if find(doc, func(n *xhtml.Node) bool {
		return isElement(n, "div") && hasClass(n, "container")
	}) != nil {
		r.Actual = `Found <div class="container"> but the spec requires <section>`
		r.Diff = `Replace <div class="container"> with <section class="container"> — use the correct semantic element`
	} else {
		r.Actual = "No element with class=\"container\" found"
		r.Diff = `Add a <section class="container"> to wrap the language badges`
	}
	return r
}

// checkSectionSocialContainer enforces the use of <section> (not <div>) with class="social-container".
func checkSectionSocialContainer(doc *xhtml.Node) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "section-social-container",
		MaxPoints: 6,
		Expected:  `<section class="social-container"> — must use the <section> semantic tag, not <div>`,
	}
	if find(doc, func(n *xhtml.Node) bool {
		return isElement(n, "section") && hasClass(n, "social-container")
	}) != nil {
		r.Passed = true
		r.Points = 6
		r.Actual = `Found <section class="social-container">`
	} else if find(doc, func(n *xhtml.Node) bool {
		return isElement(n, "div") && hasClass(n, "social-container")
	}) != nil {
		r.Actual = `Found <div class="social-container"> but the spec requires <section>`
		r.Diff = `Replace <div class="social-container"> with <section class="social-container">`
	} else {
		r.Actual = "No element with class=\"social-container\" found"
		r.Diff = `Add a <section class="social-container"> to wrap the social links`
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

// checkBadgeCount accepts 1 or 2 badges — the project limits contributors to at most 2 languages.
func checkBadgeCount(doc *xhtml.Node) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "badge-count",
		MaxPoints: 9,
		Expected:  "Between 1 and 2 <div class=\"badge\"> elements inside .container",
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

	if count >= 1 && count <= 2 {
		r.Passed = true
		r.Points = 9
	} else if count == 0 {
		r.Diff = "Add at least one <div class=\"badge\"> inside .container"
	} else {
		r.Diff = fmt.Sprintf("Found %d badges — the gallery allows a maximum of 2 programming languages per contributor", count)
	}
	return r
}

func checkBadgeBackgroundColor(doc *xhtml.Node) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "badge-background-color",
		MaxPoints: 7,
		Expected:  "Each <div class=\"badge\"> has a background-color defined via inline style attribute",
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
		r.Points = 7
		r.Actual = "All badges have inline background-color"
	} else {
		r.Actual = fmt.Sprintf("Missing background-color in: %s", strings.Join(missing, ", "))
		r.Diff = "Set the language official color as background-color directly on each .badge via the style attribute"
	}
	return r
}

func checkBadgeTextColor(doc *xhtml.Node) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "badge-text-color",
		MaxPoints: 6,
		Expected:  "Each <div class=\"badge\"> has an inline color attribute for text contrast",
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
		styleWithoutBg := strings.ReplaceAll(attr(badge, "style"), "background-color", "")
		if !strings.Contains(styleWithoutBg, "color") {
			missing = append(missing, fmt.Sprintf("badge %d", i+1))
		}
	}

	if len(missing) == 0 {
		r.Passed = true
		r.Points = 6
		r.Actual = "All badges have inline text color"
	} else {
		r.Actual = fmt.Sprintf("Missing inline color in: %s", strings.Join(missing, ", "))
		r.Diff = "Add color: white or color: black as inline style on each .badge to ensure readable contrast over the background"
	}
	return r
}

// ---------------------------------------------------------------------------
// Social link rules
// ---------------------------------------------------------------------------

// checkSocialLinkCount accepts 1 or 2 social links — the project limits contributors to at most 2.
func checkSocialLinkCount(doc *xhtml.Node) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "social-link-count",
		MaxPoints: 9,
		Expected:  "Between 1 and 2 <a class=\"social-link\"> elements inside .social-container",
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

	if count >= 1 && count <= 2 {
		r.Passed = true
		r.Points = 9
	} else if count == 0 {
		r.Diff = "Add at least one <a class=\"social-link\"> inside .social-container"
	} else {
		r.Diff = fmt.Sprintf("Found %d links — the gallery allows a maximum of 2 social links per contributor", count)
	}
	return r
}

// checkSocialLinkHref verifies that each social link has a real URL (not empty, not a placeholder).
func checkSocialLinkHref(doc *xhtml.Node) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "social-link-href",
		MaxPoints: 4,
		Expected:  "Each .social-link has a real URL as href (https://...)",
	}

	links := findAll(doc, func(n *xhtml.Node) bool {
		return isElement(n, "a") && hasClass(n, "social-link")
	})

	var invalid []string
	for i, link := range links {
		h := attr(link, "href")
		isReal := (strings.HasPrefix(h, "https://") || strings.HasPrefix(h, "http://")) &&
			!strings.Contains(h, "your-username")
		if !isReal {
			invalid = append(invalid, fmt.Sprintf("link %d (%q)", i+1, h))
		}
	}

	if len(invalid) == 0 {
		r.Passed = true
		r.Points = 4
		r.Actual = "All social links have real URLs"
	} else {
		r.Actual = fmt.Sprintf("Invalid or placeholder href in: %s", strings.Join(invalid, ", "))
		r.Diff = "Replace placeholder URLs with your actual profile links (must start with https://)"
	}
	return r
}

func checkSocialLinkTarget(doc *xhtml.Node) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "social-link-target",
		MaxPoints: 5,
		Expected:  `Each .social-link has target="_blank" so links open in a new tab`,
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
		r.Diff = `Add target="_blank" to every <a class="social-link"> element`
	}
	return r
}

// checkSocialLinkIcon verifies that each social link contains an <img class="social-icon"> child.
func checkSocialLinkIcon(doc *xhtml.Node) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "social-link-icon",
		MaxPoints: 4,
		Expected:  "Each .social-link contains an <img class=\"social-icon\"> with a Devicons CDN src",
	}

	links := findAll(doc, func(n *xhtml.Node) bool {
		return isElement(n, "a") && hasClass(n, "social-link")
	})
	if len(links) == 0 {
		r.Actual = "No .social-link elements found"
		r.Diff = "Add .social-link elements first"
		return r
	}

	var missing []string
	for i, link := range links {
		icon := find(link, func(n *xhtml.Node) bool {
			return isElement(n, "img") && hasClass(n, "social-icon")
		})
		if icon == nil {
			missing = append(missing, fmt.Sprintf("link %d", i+1))
		}
	}

	if len(missing) == 0 {
		r.Passed = true
		r.Points = 4
		r.Actual = "All social links have an <img class=\"social-icon\">"
	} else {
		r.Actual = fmt.Sprintf("Missing <img class=\"social-icon\"> inside: %s", strings.Join(missing, ", "))
		r.Diff = "Add an <img class=\"social-icon\" src=\"...\"> inside each .social-link, pointing to the Devicons CDN"
	}
	return r
}

// ---------------------------------------------------------------------------
// CSS fidelity checks
// ---------------------------------------------------------------------------

// checkHasStyleBlock verifies that a <style> block exists AND is positioned after </article>.
func checkHasStyleBlock(htmlContent string) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "has-style-block",
		MaxPoints: 3,
		Expected:  "A <style> block present after the closing </article> tag",
	}
	lower := strings.ToLower(htmlContent)
	styleIdx := strings.Index(lower, "<style")
	articleCloseIdx := strings.Index(lower, "</article>")

	hasStyle := styleIdx >= 0 && strings.Contains(lower, "</style>")
	isAfterArticle := articleCloseIdx >= 0 && styleIdx > articleCloseIdx

	if hasStyle && isAfterArticle {
		r.Passed = true
		r.Points = 3
		r.Actual = "<style> block found after </article>"
	} else if hasStyle {
		r.Actual = "<style> block found but it is not positioned after </article>"
		r.Diff = "Move the <style> block so it appears after the closing </article> tag"
	} else {
		r.Actual = "No <style> block found"
		r.Diff = "Add a <style> block with your CSS rules, placed after </article>"
	}
	return r
}

// checkArticleCss verifies that the article selector is styled with the required card properties.
func checkArticleCss(htmlContent string) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "article-css",
		MaxPoints: 3,
		Expected:  "article CSS rule with max-width, border-radius and box-shadow",
	}
	style := styleBlockContent(htmlContent)
	hasArticleRule := strings.Contains(style, "article")
	hasBorderRadius := strings.Contains(style, "border-radius")
	hasBoxShadow := strings.Contains(style, "box-shadow")

	if hasArticleRule && hasBorderRadius && hasBoxShadow {
		r.Passed = true
		r.Points = 3
		r.Actual = "article selector styled with border-radius and box-shadow"
	} else {
		var missing []string
		if !hasArticleRule {
			missing = append(missing, "article selector")
		}
		if !hasBorderRadius {
			missing = append(missing, "border-radius")
		}
		if !hasBoxShadow {
			missing = append(missing, "box-shadow")
		}
		r.Actual = fmt.Sprintf("Missing in CSS: %s", strings.Join(missing, ", "))
		r.Diff = "Add an article { } CSS rule that includes border-radius and box-shadow to style the card"
	}
	return r
}

// checkBadgeCss verifies that the .badge selector has padding and border-radius.
func checkBadgeCss(htmlContent string) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "badge-css",
		MaxPoints: 4,
		Expected:  ".badge CSS rule with padding and border-radius",
	}
	style := styleBlockContent(htmlContent)
	hasBadgeRule := strings.Contains(style, ".badge")
	hasPadding := strings.Contains(style, "padding")
	hasBorderRadius := strings.Contains(style, "border-radius")

	if hasBadgeRule && hasPadding && hasBorderRadius {
		r.Passed = true
		r.Points = 4
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
		r.Diff = "Add a .badge { } CSS rule with padding and border-radius"
	}
	return r
}

// checkContainerCss verifies that .container uses flexbox layout.
func checkContainerCss(htmlContent string) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "container-css",
		MaxPoints: 3,
		Expected:  ".container CSS rule with display: flex",
	}
	style := styleBlockContent(htmlContent)
	if strings.Contains(style, ".container") && strings.Contains(style, "display") && strings.Contains(style, "flex") {
		r.Passed = true
		r.Points = 3
		r.Actual = ".container styled with display: flex"
	} else {
		r.Actual = "Missing .container CSS with display: flex"
		r.Diff = "Add a .container { display: flex; ... } CSS rule to lay out the badges side by side"
	}
	return r
}

// checkSocialLinkCss verifies that .social-link has the required visual properties for the pill shape.
func checkSocialLinkCss(htmlContent string) objects.CheckResult {
	r := objects.CheckResult{
		Rule:      "social-link-css",
		MaxPoints: 5,
		Expected:  ".social-link CSS rule with display: flex, border-radius and text-decoration",
	}
	style := styleBlockContent(htmlContent)
	hasRule := strings.Contains(style, ".social-link")
	hasFlex := strings.Contains(style, "display") && strings.Contains(style, "flex")
	hasBorderRadius := strings.Contains(style, "border-radius")
	hasTextDecoration := strings.Contains(style, "text-decoration")

	if hasRule && hasFlex && hasBorderRadius && hasTextDecoration {
		r.Passed = true
		r.Points = 5
		r.Actual = ".social-link styled with display: flex, border-radius and text-decoration"
	} else {
		var missing []string
		if !hasRule {
			missing = append(missing, ".social-link rule")
		}
		if !hasFlex {
			missing = append(missing, "display: flex")
		}
		if !hasBorderRadius {
			missing = append(missing, "border-radius")
		}
		if !hasTextDecoration {
			missing = append(missing, "text-decoration")
		}
		r.Actual = fmt.Sprintf("Missing in .social-link: %s", strings.Join(missing, ", "))
		r.Diff = "Add a .social-link { } rule with display: flex, border-radius (for the pill shape) and text-decoration: none"
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

// textContent returns the concatenated text content of a node and all its descendants.
func textContent(n *xhtml.Node) string {
	if n.Type == xhtml.TextNode {
		return n.Data
	}
	var sb strings.Builder
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		sb.WriteString(textContent(c))
	}
	return sb.String()
}

// styleBlockContent extracts the raw content between the first <style> and </style> tags.
// Returns an empty string if no style block is found.
func styleBlockContent(htmlContent string) string {
	lower := strings.ToLower(htmlContent)
	start := strings.Index(lower, "<style")
	if start < 0 {
		return ""
	}
	open := strings.Index(lower[start:], ">")
	if open < 0 {
		return ""
	}
	contentStart := start + open + 1
	end := strings.Index(lower[contentStart:], "</style>")
	if end < 0 {
		return ""
	}
	return strings.ToLower(htmlContent[contentStart : contentStart+end])
}
