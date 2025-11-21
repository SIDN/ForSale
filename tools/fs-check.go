package main

import (
    "fmt"
    "net"
    "net/url"
    "os"
    "regexp"
    "sort"
    "strings"
)

type ValidationResult struct {
    Record           string
    HasVersion       bool
    Tag              string
    Value            string
    ValidSyntax      bool
    Warnings         []string
    Errors           []string
}

type Summary struct {
    Domain          string
    ForSale         bool
    ValidRecords    int
    InvalidRecords  int
    DuplicatesFound bool
    Results         []ValidationResult
}

var (
    versionPrefix = "v=FORSALE1;"
    // Allowed content tags
    tags = []string{"fcod=", "ftxt=", "furi=", "fval="}
    // fval currency and amount: uppercase letters + integer part + optional . fractional digits
    reFval = regexp.MustCompile(`^[A-Z]+[0-9]+(\.[0-9]+)?$`)
    // Simple mailto/tel validators
    reMailto = regexp.MustCompile(`^mailto:[^ \t]+$`)
    reTel    = regexp.MustCompile(`^tel:\+?[0-9][0-9\-\.\(\) ]*$`)
)

func trimSpacesAfterVersion(s string) string {
    // Allow any number of spaces immediately after version tag
    if strings.HasPrefix(s, versionPrefix) {
        rest := s[len(versionPrefix):]
        return versionPrefix + strings.TrimLeft(rest, " ")
    }
    return s
}

func parseRecord(raw string) ValidationResult {
    r := ValidationResult{Record: raw}

    // Normalize spaces after version
    raw = trimSpacesAfterVersion(strings.TrimSpace(raw))

    // Must start with exact version tag
    if !strings.HasPrefix(raw, versionPrefix) {
        r.Errors = append(r.Errors, "missing or invalid version tag (must start with 'v=FORSALE1;')")
        return r
    }
    r.HasVersion = true

    content := raw[len(versionPrefix):]
    if content == "" {
        // Content absent is allowed; domain still considered for sale
        r.ValidSyntax = true
        return r
    }

    // Ensure only one tag-value pair: find which tag appears
    var foundTag string
    for _, t := range tags {
        if strings.HasPrefix(content, t) {
            foundTag = t[:len(t)-1] // strip '=' for reporting
            break
        }
    }
    if foundTag == "" {
        r.Errors = append(r.Errors, "invalid content tag (must be one of fcod=, ftxt=, furi=, fval=)")
        return r
    }

    // Extract value
    value := content[len(foundTag)+1:] // +1 to account for '='
    r.Tag = foundTag
    r.Value = value

    // Check that there's only one tag-value pair: no semicolons introducing another tag
    for _, t := range tags {
        if strings.Contains(value, ";"+t) {
            r.Warnings = append(r.Warnings, "ambiguous content: value contains ';"+t+"' which may be confusing")
        }
    }
    if strings.Contains(value, ";") {
        // Semicolons inside value are allowed, but could be confusing
        r.Warnings = append(r.Warnings, "value contains ';' which may appear as a second tag — ensure proper escaping")
    }

    // Validate per tag
    switch foundTag {
    case "fcod":
        if len(value) < 1 {
            r.Errors = append(r.Errors, "fcod value must be at least 1 octet")
        }
    case "ftxt":
        if len(value) < 1 {
            r.Errors = append(r.Errors, "ftxt value must be at least 1 octet")
        }
        // Warn on raw script tags (not forbidden, but risky)
        if strings.Contains(strings.ToLower(value), "<script") {
            r.Warnings = append(r.Warnings, "ftxt contains '<script' — potential XSS risk if rendered")
        }
    case "furi":
        if !validateURI(value, &r) {
            // validateURI will add specific errors/warnings
        }
    case "fval":
        if !reFval.MatchString(value) {
            r.Errors = append(r.Errors, "fval must match '^[A-Z]+[0-9]+(\\.[0-9]+)?$' (currency + amount)")
        } else {
            // Check recommended ISO 4217 3-letter fiat currency (non-fatal)
            cur := leadingLetters(value)
            if len(cur) != 3 {
                r.Warnings = append(r.Warnings, "currency code is not 3 letters; non-standard codes are allowed but not recommended")
            }
        }
    default:
        // Should not happen
        r.Errors = append(r.Errors, "unknown tag")
    }

    r.ValidSyntax = len(r.Errors) == 0
    return r
}

func leadingLetters(s string) string {
    i := 0
    for i < len(s) {
        c := s[i]
        if c < 'A' || c > 'Z' {
            break
        }
        i++
    }
    return s[:i]
}

func validateURI(u string, r *ValidationResult) bool {
    if strings.ContainsAny(u, " \t") {
        r.Errors = append(r.Errors, "URI must not contain spaces; use percent-encoding")
        return false
    }
    // mailto and tel are special cases
    if strings.HasPrefix(u, "mailto:") {
        if !reMailto.MatchString(u) {
            r.Errors = append(r.Errors, "invalid mailto URI syntax")
            return false
        }
        return true
    }
    if strings.HasPrefix(u, "tel:") {
        if !reTel.MatchString(u) {
            r.Errors = append(r.Errors, "invalid tel URI syntax")
            return false
        }
        return true
    }

    // Parse as generic URL (expects http/https ideally)
    parsed, err := url.Parse(u)
    if err != nil {
        r.Errors = append(r.Errors, "invalid URI: "+err.Error())
        return false
    }
    if parsed.Scheme == "" || parsed.Host == "" {
        r.Errors = append(r.Errors, "URI must include scheme and host")
        return false
    }
    // Recommend schemes
    switch strings.ToLower(parsed.Scheme) {
    case "http", "https":
        // ok
    default:
        r.Warnings = append(r.Warnings, "non-recommended scheme (recommended: http, https, mailto, tel)")
    }
    return true
}

func lookupTXT(name string) ([]string, error) {
    return net.LookupTXT(name)
}

func validateDomain(domain string) Summary {
    target := "_for-sale." + domain
    txts, err := lookupTXT(target)
    summary := Summary{Domain: domain}

    if err != nil {
        // No records => not for sale (no version found)
        return summary
    }

    results := make([]ValidationResult, 0, len(txts))
    seenPairs := map[string]struct{}{}
    duplicates := false

    for _, rec := range txts {
        vr := parseRecord(rec)
        results = append(results, vr)

        if vr.HasVersion {
            summary.ForSale = true
            if vr.Tag != "" {
                key := vr.Tag + "=" + vr.Value
                if _, ok := seenPairs[key]; ok {
                    duplicates = true
                } else {
                    seenPairs[key] = struct{}{}
                }
            }
        }
    }

    for _, r := range results {
        if r.ValidSyntax {
            summary.ValidRecords++
        } else {
            summary.InvalidRecords++
        }
    }
    summary.DuplicatesFound = duplicates
    summary.Results = results
    return summary
}

func printSummary(s Summary) {
    fmt.Printf("Domain: %s\n", s.Domain)
    fmt.Printf("For sale: %v\n", s.ForSale)
    fmt.Printf("Valid records: %d, Invalid records: %d\n", s.ValidRecords, s.InvalidRecords)
    fmt.Printf("Duplicate tag-value pairs in RRset: %v\n", s.DuplicatesFound)
    fmt.Println()

    // Sort for stable output: errors first, then warnings
    sort.Slice(s.Results, func(i, j int) bool {
        if s.Results[i].ValidSyntax == s.Results[j].ValidSyntax {
            return s.Results[i].Record < s.Results[j].Record
        }
        // invalid first
        return !s.Results[i].ValidSyntax && s.Results[j].ValidSyntax
    })

    for idx, r := range s.Results {
        fmt.Printf("Record %d:\n", idx+1)
        fmt.Printf("  Raw: %q\n", r.Record)
        fmt.Printf("  Has version: %v\n", r.HasVersion)
        if r.Tag != "" {
            fmt.Printf("  Tag: %s\n", r.Tag)
            fmt.Printf("  Value: %s\n", r.Value)
        } else {
            fmt.Printf("  Tag: (none)\n")
        }
        fmt.Printf("  Valid syntax: %v\n", r.ValidSyntax)

        if len(r.Errors) > 0 {
            fmt.Println("  Errors:")
            for _, e := range r.Errors {
                fmt.Printf("    - %s\n", e)
            }
        }
        if len(r.Warnings) > 0 {
            fmt.Println("  Warnings:")
            for _, w := range r.Warnings {
                fmt.Printf("    - %s\n", w)
            }
        }
        fmt.Println()
    }
}

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s <domain>\n", os.Args[0])
        os.Exit(2)
    }
    domain := strings.TrimSpace(os.Args[1])
    if domain == "" {
        fmt.Fprintln(os.Stderr, "Domain must not be empty")
        os.Exit(2)
    }
    s := validateDomain(domain)
    printSummary(s)
    if s.ForSale {
        os.Exit(0)
    } else {
        // Non-zero exit could be used to signal "not for sale" in scripts; adjust as desired.
        os.Exit(1)
    }
}
