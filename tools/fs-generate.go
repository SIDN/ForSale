package main

import (
    "bufio"
    "fmt"
    "net/url"
    "os"
    "regexp"
    "strings"
)

var (
    reFval   = regexp.MustCompile(`^[A-Z]+[0-9]+(\.[0-9]+)?$`)
    reMailto = regexp.MustCompile(`^mailto:[^ \t]+$`)
    reTel    = regexp.MustCompile(`^tel:\+?[0-9][0-9\\-\\.\\(\\) ]*$`)
)

func ask(prompt string) string {
    fmt.Print(prompt)
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    return strings.TrimSpace(text)
}

func escapeQuotes(s string) string {
    return strings.ReplaceAll(s, `"`, `\"`)
}

func validateFval(v string) error {
    if !reFval.MatchString(v) {
        return fmt.Errorf("must be uppercase currency letters followed by amount, e.g. USD750 or EUR99.99")
    }
    return nil
}

func validateFuri(u string) error {
    if strings.ContainsAny(u, " \t") {
        return fmt.Errorf("URI must not contain spaces; use percent-encoding")
    }
    if strings.HasPrefix(u, "mailto:") {
        if !reMailto.MatchString(u) {
            return fmt.Errorf("invalid mailto URI syntax")
        }
        return nil
    }
    if strings.HasPrefix(u, "tel:") {
        if !reTel.MatchString(u) {
            return fmt.Errorf("invalid tel URI syntax")
        }
        return nil
    }
    parsed, err := url.Parse(u)
    if err != nil {
        return fmt.Errorf("invalid URI: %v", err)
    }
    if parsed.Scheme == "" || parsed.Host == "" {
        return fmt.Errorf("URI must include scheme and host")
    }
    return nil
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

    fmt.Println("Generating _for-sale TXT records for domain:", domain)

    records := []string{}
    seenPairs := make(map[string]struct{})

    for {
        fmt.Println("\nChoose content type:")
        fmt.Println(" 1) fval (asking price)")
        fmt.Println(" 2) furi (contact URI)")
        fmt.Println(" 3) ftxt (free text)")
        fmt.Println(" 4) fcod (code)")
        fmt.Println(" 5) view (show current records)")
        fmt.Println(" 6) done (finish)")
        choice := ask("Enter choice (1-6): ")

        if choice == "6" {
            break
        }
        if choice == "5" {
            fmt.Println("\nCurrent records:")
            if len(records) == 0 {
                fmt.Println("  (none yet)")
            } else {
                for _, r := range records {
                    fmt.Println("  " + r)
                }
            }
            continue
        }

        var tag, value string
        switch choice {
        case "1":
            tag = "fval"
            for {
                v := ask("Enter asking price (or type 'cancel' to return): ")
                if strings.ToLower(v) == "cancel" {
                    fmt.Println("Cancelled, returning to main menu.")
                    tag, value = "", ""
                    break
                }
                if err := validateFval(v); err != nil {
                    fmt.Println("Invalid:", err)
                    continue
                }
                value = v
                break
            }
        case "2":
            tag = "furi"
            for {
                v := ask("Enter contact URI (or type 'cancel' to return): ")
                if strings.ToLower(v) == "cancel" {
                    fmt.Println("Cancelled, returning to main menu.")
                    tag, value = "", ""
                    break
                }
                if err := validateFuri(v); err != nil {
                    fmt.Println("Invalid:", err)
                    continue
                }
                value = v
                break
            }
        case "3":
            tag = "ftxt"
            for {
                v := ask("Enter free text (or type 'cancel' to return): ")
                if strings.ToLower(v) == "cancel" {
                    fmt.Println("Cancelled, returning to main menu.")
                    tag, value = "", ""
                    break
                }
                if len(v) < 1 {
                    fmt.Println("Invalid: must not be empty")
                    continue
                }
                value = v
                break
            }
        case "4":
            tag = "fcod"
            for {
                v := ask("Enter code value (or type 'cancel' to return): ")
                if strings.ToLower(v) == "cancel" {
                    fmt.Println("Cancelled, returning to main menu.")
                    tag, value = "", ""
                    break
                }
                if len(v) < 1 {
                    fmt.Println("Invalid: must not be empty")
                    continue
                }
                value = v
                break
            }
        default:
            fmt.Println("Invalid choice")
            continue
        }

        if tag == "" {
            // user cancelled, skip adding
            continue
        }

        valueEsc := escapeQuotes(value)
        key := tag + "=" + valueEsc
        if _, exists := seenPairs[key]; exists {
            fmt.Println("Duplicate record detected — not allowed by the draft. Skipping.")
            continue
        }
        seenPairs[key] = struct{}{}

        record := fmt.Sprintf(`_for-sale.%s. IN TXT "v=FORSALE1;%s=%s"`, domain, tag, valueEsc)
        records = append(records, record)
        fmt.Println("Record added.")
    }

    fmt.Println("\nFinal DNS zone file snippet:")
    for _, r := range records {
        fmt.Println(r)
    }
}
