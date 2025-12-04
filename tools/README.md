# 🧪 Tools to play with and test the ForSale method

> [!NOTE]
> Tools work for draft-davids-forsalereg-16

Build:

~~~
go build name.go
~~~

## fs-check.go

A validator / syntax checker

## fs-generate.go

A record generator - see demo output below

Also see

* https://forsalereg.sidnlabs.nl/demo
* https://forsale.bitfire.nl
* https://forsalereg.sidnlabs.nl/forsale (in Dutch 🇳🇱)

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
