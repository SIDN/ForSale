# ♻️ ForSale - Advanced usage

Here are some more (advanced) examples. See [draft-davids-forsalereg](https://datatracker.ietf.org/doc/draft-davids-forsalereg/) for all details.

## 🌐 Unicode

~~~
_for-sale.example.nl.	1800	IN	TXT	"v=FORSALE1;ftxt=人間が読める形式の情報がここにあります。"
_for-sale.example.nl.	1800	IN	TXT	"v=FORSALE1;ftxt=ᛞᚨᚷᚱ ᛖᚱ ᛒᛃᚨᚱᛏ"
_for-sale.example.nl.	1800	IN	TXT	"v=FORSALE1;furi=https://⠎⠑⠁⠎⠕⠝⠎⠛⠗⠑⠑⠞⠊⠝⠛⠎.sidnlabs.nl/"
_for-sale.example.nl.	1800	IN	TXT	"v=FORSALE1;fval=JPY235566"
~~~

> [!IMPORTANT]
> Unicode characters should be [encoded as UTF-8](https://www.fileformat.info/info/unicode/char/1f600/index.htm), like this: 😀 (U+1F600) → `F0 9F 98 80`

> [!CAUTION]
> It is not always possible to enter Unicode directly in DNS records and zonefiles. Sometimes a [presentation format](https://datatracker.ietf.org/doc/html/rfc1035#section-5.1) has to be used.
> Like this: 😀 (U+1F600) → `F0 9F 98 80` → `_for-sale.example.nl.	1800	IN	TXT	"v=FORSALE1;ftxt=\240\159\152\128"`

> [!NOTE]
> Please note that representing a '😀' in a ForSale TXT record consumes not one, but four characters as the example above shows.

> [!TIP]
> This little one liner might help to test things: `dig TXT _for-sale.example.nl +short | perl -pe 's/\\([0-9]{3})/chr($1)/ge'`
 
## What is the `fcod=` tag used for?

The `fcod=` tag is a flexible code that only has meaning for parties that have agreed on how to use it. It acts as a private instruction between cooperating systems, such as a domain registry and its registrars, and can be used to control what happens when a domain is for sale - for example, which page is shown or how information is displayed. It has no universal meaning, and outsiders should not attempt to interpret it.

### A few `fcod=` examples

**Encoded proprietary landing page URL**

A domain name registry may allow registrars to enter a "for sale" URL into its back-end system. From that URL, a unique code is generated. This code is then inserted as the value of the `fcod=`content tag in the "_for-sale" TXT record of a domain name, as shown in the example below.

When a user checks the availability of a domain name using a registry-provided tool (for example, a web interface), the registry may use this code to redirect the user to the appropriate "for sale" URL. This URL may include a query component containing the domain name, for example:

`https://forsale-url.example.com/exco?d=example.org`

The rationale for this approach is that the controlling parties retain authority over redirection URLs and any other information derived from the content tag. This helps prevent users from being redirected to unintended or malicious destinations or from being presented with unintended content. In addition, this approach allows the interpretation of `fcod=` content values to be adjusted centrally in back-end systems (for example, determining which "for sale" URL to redirect to) without requiring modifications to the "_for-sale" TXT records.

The following example shows a string encoded using Base64, preceded by the prefix EXCO- ('Example Company'), used as the value of the content tag:

`_for-sale.example.nl. 1800 IN TXT "v=FORSALE1;fcod=EXCO-S2lscm95IHdhcyBoZXJl"`

**Arbitrary formatting or conditional display instructions**

A party may include proprietary code containing arbitrary formatting or conditional display instructions, such as adding an extra banner (for example, "eligibility criteria apply") or specifying a particular style, including color, font, emojis, or logos. For example:

`_for-sale.example.nl. 1800 IN TXT "v=FORSALE1;fcod=CSS-AGgAMQAgAHsAZgBvAG4AdAAtAHcAZQBpAGcAaAB0ADoAIAA3ADAAMAB9"`

`_for-sale.example.nl. 1800 IN TXT "v=FORSALE1;fcod=LOGO-AHMAYQBsAGUALgBqAHAAZw=="`

**Digital signatures**

A party could include a digital signature for various reasons, for example to prove that it has validated a particular furi= value. For example:

`_for-sale.example.nl. 1800 IN TXT "v=FORSALE1;fcod=SIG-AGQAYQBzAGQAYQBzAGQACg=="`

The interpretation of the `fcod=` content value is entirely up to the issuing party.

> [!IMPORTANT]
> Remember: `fcod=` content values only have meaning between parties that have agreed on their use. They have no universal meaning, unlike the `furi=` and `fval=` content tag-value pairs.


## Is there a quick way to generate 'presentation format' ?

Yes! Try this:

```
python3 -c "
s = 'v=FORSALE1;ftxt=😀'; 
print(''.join(c if ord(c) < 128 else ''.join(f'\\{b:03d}' for b in c.encode('utf-8')) for c in s))
"
```
The output will be:

`v=FORSALE1;ftxt=\240\159\152\128`

Obvisously this is just a simple quick example. There are probably better ways.

## ⚒ WORK IN PROGRESS - PLEASE CHECK BACK LATER

