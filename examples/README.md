# ♻️ ForSale - Advanced usage

Here are some more (advanced) examples. See [draft-davids-forsalereg](https://datatracker.ietf.org/doc/draft-davids-forsalereg/) for all details.

## 🌐 Unicode

~~~
_for-sale.example.nl.	1800	IN	TXT	"v=FORSALE1;ftxt=人間が読める形式の情報がここにあります。"
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

## ⚒ WORK IN PROGRESS - PLEASE CHECK BACK LATER

