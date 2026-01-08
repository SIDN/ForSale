# 🪧 ForSale - Advanced usage
> [!NOTE]
>  Here are some more examples.

## Unicode

~~~
_for-sale.example.nl.	1800	IN	TXT	"v=FORSALE1;ftxt=人間が読める形式の情報がここにあります。"
_for-sale.example.nl.	1800	IN	TXT	"v=FORSALE1;furi=https://⠎⠑⠁⠎⠕⠝⠎⠛⠗⠑⠑⠞⠊⠝⠛⠎.sidnlabs.nl/"
_for-sale.example.nl.	1800	IN	TXT	"v=FORSALE1;fval=JPY235566"
~~~

> [!CAUTION]
> It is not always possible to enter Unicode directly in DNS records and zonefiles. Sometimes a presentation format has to be used. See [RFC1035, Section 5.1](https://datatracker.ietf.org/doc/html/rfc1035#section-5.1).

## What is the `fcod=` tag used for?

The `fcod=` tag is a flexible code that only has meaning for parties that have agreed on how to use it. It acts as a private instruction between cooperating systems, such as a domain registry and its registrars, and can be used to control what happens when a domain is for sale - for example, which page is shown or how information is displayed. It has no universal meaning, and outsiders should not attempt to interpret it.

## ⚒ WORK IN PROGRESS - PLEASE CHECK BACK LATER

