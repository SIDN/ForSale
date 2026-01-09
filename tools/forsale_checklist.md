# Implementation Checklist for draft-davids-forsalereg-19

## Record Format Requirements

### Version Tag
- [ ] MUST include version tag `v=FORSALE1;` (case-sensitive, exact match)
- [ ] Version tag MUST be at the beginning of the record
- [ ] Version tag MUST use exact bytes: %x76.3D.46.4F.52.53.41.4C.45.31.3B

### Content Tags
- [ ] MUST NOT include more than one tag-value pair per TXT record
- [ ] Content tags MUST be lowercase and case-sensitive
- [ ] Every tag-value pair in the RRset MUST be unique
- [ ] Multiple instances of the same content tag type MAY occur (with different values)

### RDATA Structure
- [ ] RDATA MUST consist of a single character-string (max 255 octets)
- [ ] MUST NOT use multiple character-strings requiring concatenation
- [ ] Total length per record: version tag + optional content ≤ 255 octets

## Content Tag Types

### fcod= (Code Format)
- [ ] Content value MUST consist of at least 1 octet
- [ ] Maximum length: 239 octets (255 minus version tag)
- [ ] Format: `fcod=` followed by 1-239 octets

### ftxt= (Free Text)
- [ ] Content value MUST consist of at least 1 octet
- [ ] Maximum length: 239 octets
- [ ] SHOULD provide meaningful context (more than single character)
- [ ] SHOULD avoid URIs (use furi= instead)
- [ ] Format: `ftxt=` followed by 1-239 octets

### furi= (URI Format)
- [ ] Content value MUST contain exactly one URI or IRI
- [ ] URI MUST conform to RFC 3986 syntax
- [ ] IRI MUST conform to RFC 3987 syntax
- [ ] URIs MUST use proper percent-encoding (e.g., spaces as %20)
- [ ] RECOMMENDED schemes only: http, https, mailto, tel
- [ ] Format: `furi=` followed by valid URI/IRI

### fval= (Asking Price)
- [ ] Format: `fval=` followed by currency code + amount
- [ ] Currency code: 1 or more uppercase letters (A-Z)
- [ ] RECOMMENDED: Use ISO 4217 three-letter codes (USD, EUR, GBP, JPY)
- [ ] Amount: integer part with optional fractional part
- [ ] Decimal separator MUST be period (.) - %x2E
- [ ] Total length (currency + amount): 2-239 characters
- [ ] Examples: `EUR999`, `BTC0.000010`, `USD750.50`

## Character Encoding

### Text Content
- [ ] RECOMMENDED: Encode text in UTF-8 (RFC 3629)
- [ ] RECOMMENDED: Conform to Network Unicode format (RFC 5198)
- [ ] RECOMMENDED: Use Unicode subset per RFC 9839 Section 4.3
- [ ] SHOULD avoid control characters: %x09, %x0A, %x0D
- [ ] IDNs MAY appear as A-labels or U-labels (U-labels as UTF-8)

### Parsing
- [ ] MUST parse raw TXT RDATA content, not escaped presentation format
- [ ] If multiple character-strings exist, SHOULD concatenate before UTF-8 interpretation
- [ ] Processors SHOULD handle UTF-8 encoded non-ASCII content

## Record Validation

### Valid Records
- [ ] MUST have valid version tag to be considered valid
- [ ] If content is absent but version tag valid, SHOULD assume domain is for sale
- [ ] If content is invalid but version tag valid, SHOULD assume domain is for sale
- [ ] TXT records without version tag MUST NOT be interpreted as _for-sale indicators

### Invalid Content
- [ ] MUST NOT contain text suggesting domain is NOT for sale
- [ ] Empty content tags are considered invalid (e.g., `fcod=`)
- [ ] Unrecognized tags are considered invalid (e.g., `foo=bar`)

## RRset Handling

### Multiple Records
- [ ] No limit on number of TXT records in RRset
- [ ] Processors MAY select one or more records from RRset
- [ ] All records in RRset MUST have same TTL (RFC 2181)

### TTL Recommendations
- [ ] RECOMMENDED: TTL ≤ 3600 seconds (1 hour)

## Placement Rules

### Valid Placements
- [ ] MAY be placed at any DNS level except under .arpa
- [ ] MUST be a leaf node name (not `xyz._for-sale.example`)
- [ ] Valid: `_for-sale.example.`
- [ ] Valid: `_for-sale.aaa.example.`
- [ ] Valid: `_for-sale.exco.bbb.example.`

### Invalid Placements
- [ ] MUST NOT use under .arpa infrastructure TLD
- [ ] Processors MUST ignore records under .arpa
- [ ] MUST NOT use wildcards like `_for-sale.*.example.`
- [ ] Out of scope: Special-Use Domain Names (RFC 6761) like .onion, .alt

### Coexistence with Wildcards
- [ ] MAY coexist with wildcard records at other labels
- [ ] Example valid: `*.example` with `_for-sale.example`

## Processing Requirements

### Discovery
- [ ] Processors MUST check for valid version tag before processing
- [ ] MUST ignore records without valid version tag
- [ ] MAY apply robustness principle (liberal in acceptance)
- [ ] MAY accept spaces (%x20) after version tag

### Error Handling
- [ ] MUST handle missing keys gracefully
- [ ] SHOULD implement try-catch for all operations
- [ ] When content ambiguous/invalid with valid version: assume for sale
- [ ] MAY sanitize unexpected control characters by:
  - Replacing with space (%x20), OR
  - Replacing with Unicode REPLACEMENT CHARACTER (U+FFFD)

### Security
- [ ] MUST validate and sanitize content before display
- [ ] MUST NOT automatically follow URIs without user consent
- [ ] SHOULD present URI to user for inspection before navigation
- [ ] MUST protect against XSS attacks
- [ ] MUST protect against SQL injection
- [ ] MUST protect against Unicode manipulation attacks
- [ ] SHOULD perform URI reputation checks before display
- [ ] SHOULD display disclaimers for fval= prices ("indicative only")
- [ ] MUST NOT make automated purchases based solely on fval= data

### Content Display
- [ ] SHOULD handle non-ASCII data correctly
- [ ] SHOULD display UTF-8 encoded content properly
- [ ] MAY sanitize control characters for display

## Operational Considerations

### Wildcards
- [ ] MUST be cautious with wildcard interactions
- [ ] Version tag distinguishes valid records from wildcard expansion
- [ ] Watch for edge cases with CNAMEs and DNAMEs

### Currency (fval=)
- [ ] RECOMMENDED: Use standard ISO 4217 codes for fiat currencies
- [ ] MAY support cryptocurrency codes (BTC, ETH, etc.)
- [ ] Currency code MUST be uppercase letters only

### Robustness
- [ ] MAY be liberal in accepting variations
- [ ] MAY accept spaces after version tag
- [ ] MAY implement proprietary stricter formats by agreement

### Scope
- [ ] Mechanism requires domain to be resolvable in DNS
- [ ] May not work during redemption period
- [ ] May not work in pendingDelete status
- [ ] May not work when DNSSEC validation fails (bogus state)

## Implementation Recommendations

### Content Creation
- [ ] SHOULD NOT publish _for-sale if domain not actually for sale
- [ ] SHOULD provide at least one content tag for engagement
- [ ] SHOULD avoid ambiguous constructs (e.g., semicolons in fcod values)
- [ ] SHOULD remove _for-sale indicator when domain no longer for sale

### Privacy
- [ ] SHOULD warn users about public visibility
- [ ] SHOULD warn about data scraping risks (email, phone)
- [ ] SHOULD warn about spam/unwanted contact potential

### User Interface
- [ ] SHOULD require user confirmation before following URIs
- [ ] SHOULD display disclaimers for prices
- [ ] SHOULD indicate data source (DNS TXT record)
- [ ] SHOULD provide fallback to WHOIS/RDAP if content missing/invalid

## IANA Considerations
- [ ] Entry required in "Underscored and Globally Scoped DNS Node Names" registry
- [ ] RR Type: TXT
- [ ] Node Name: _for-sale

## Testing Checklist

### Valid Record Tests
- [ ] Test: `v=FORSALE1;` (minimal valid)
- [ ] Test: `v=FORSALE1;fcod=ABC123`
- [ ] Test: `v=FORSALE1;ftxt=For sale`
- [ ] Test: `v=FORSALE1;furi=https://example.com`
- [ ] Test: `v=FORSALE1;fval=USD1000`
- [ ] Test: Multiple records in RRset

### Invalid Record Tests
- [ ] Test: Missing version tag
- [ ] Test: Wrong version tag case (`V=FORSALE1`)
- [ ] Test: Multiple character-strings
- [ ] Test: Content exceeding 255 octets
- [ ] Test: Multiple tag-value pairs in single record
- [ ] Test: Under .arpa TLD (must ignore)
- [ ] Test: Non-leaf placement

### Edge Cases
- [ ] Test: UTF-8 content
- [ ] Test: Non-ASCII IDN in furi=
- [ ] Test: Control characters in ftxt=
- [ ] Test: Wildcard coexistence
- [ ] Test: Empty content after tag (e.g., `fcod=`)
- [ ] Test: Unrecognized tags
- [ ] Test: Spaces after version tag
