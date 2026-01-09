# 🧪 Tools to play with and test the ForSale method

> [!NOTE]
> Tools work for draft-davids-forsalereg-16.

Build:

~~~
go build name.go
~~~

## fs-check.go

A basic validator / syntax checker

> [!TIP]
> If you are a developer who wants to write your own tools, make sure to carefully read [the draft](https://datatracker.ietf.org/doc/draft-davids-forsalereg/) or the [implementation checklist](forsale_checklist.md).

## fs-generate.go

A basic record generator - see demo output below.

Also see

* https://forsalereg.sidnlabs.nl/demo
* https://forsale.bitfire.nl
* https://forsalereg.sidnlabs.nl/forsale (in Dutch 🇳🇱)

Some domain names to test:
~~~
example.nl
spongat.nl
bitfire.nl
j78.nl
example.co.nl
nascar.watkins-glen.ny.us
pnawebloket.nl
pna-webloket.nl

testdns.nl
forsale.testdns.nl
dynamic.testdns.nl
परीक्षा.testdns.nl
xn--11b5bs3a9aj6g.testdns.nl 
~~~

## Demo output of fs-generate

~~~
./fs-generate example.nl
Generating _for-sale TXT records for domain: example.nl

Choose content type:
 1) fval (asking price)
 2) furi (contact URI)
 3) ftxt (free text)
 4) fcod (code)
 5) view (show current records)
 6) done (finish)
Enter choice (1-6): 1
Enter asking price (or type 'cancel' to return): EUR150
Record added.

Choose content type:
 1) fval (asking price)
 2) furi (contact URI)
 3) ftxt (free text)
 4) fcod (code)
 5) view (show current records)
 6) done (finish)
Enter choice (1-6): 2
Enter contact URI (or type 'cancel' to return): https://example.nl/for-sale.txt
Record added.

Choose content type:
 1) fval (asking price)
 2) furi (contact URI)
 3) ftxt (free text)
 4) fcod (code)
 5) view (show current records)
 6) done (finish)
Enter choice (1-6): 3
Enter free text (or type 'cancel' to return): Deze domeinnaam is te koop!
Record added.

Choose content type:
 1) fval (asking price)
 2) furi (contact URI)
 3) ftxt (free text)
 4) fcod (code)
 5) view (show current records)
 6) done (finish)
Enter choice (1-6): 6

Final DNS zone file snippet:
_for-sale.example.nl. IN TXT "v=FORSALE1;fval=EUR150"
_for-sale.example.nl. IN TXT "v=FORSALE1;furi=https://example.nl/for-sale.txt"
_for-sale.example.nl. IN TXT "v=FORSALE1;ftxt=Deze domeinnaam is te koop!"
~~~
