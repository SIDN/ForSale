# 🪧 ForSale - a digital for sale sign on domain names

Various resources with regard to [draft-davids-forsalereg](https://datatracker.ietf.org/doc/draft-davids-forsalereg/)[^1] - a lightweight method to add digital for sale signs to domain names.

🎉 Ready to sell your domain name? Why not put it up for sale with a [digital For Sale sign](https://www.sidnlabs.nl/en/news-and-blogs/a-digital-for-sale-sign-for-nl-domain-names) ?

## It's that simple!

~~~
_for-sale.example.nl.	1800	IN	TXT	"v=FORSALE1;ftxt=Some human-readable info here."
_for-sale.example.nl.	1800	IN	TXT	"v=FORSALE1;furi=https://domain-sale.example.com/test"
_for-sale.example.nl.	1800	IN	TXT	"v=FORSALE1;fval=USD1500"
~~~

> [!IMPORTANT]
> 🚨 Make sure to add the version tag `v=FORSALE1;` first, followed by a valid content tag (such as `ftxt=` or `furi=`) before entering the actual content value. Only then does it count as an official "for sale" sign that specialized tools will recognize.

> [!TIP]
> Check out the [examples](/examples) directory for more cool, advanced stuff!

ℹ️ More info here: https://forsalereg.sidnlabs.nl/

## About ForSale

> **ForSale** is a standardized "_for-sale" DNS name for signaling that a domain name is available for purchase (or lease). The proposal addresses a common operational need by providing a machine-readable, globally consistent mechanism to indicate that domains are for sale, complementing inconsistent web-based parking pages and proprietary signaling methods.



## ⚒ WORK IN PROGRESS - PLEASE CHECK BACK LATER

![rpp logo](media/ForSale-banner.png)

[^1]: https://datatracker.ietf.org/doc/draft-davids-forsalereg/
