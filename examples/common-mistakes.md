# 🫣 ForSale - Common mistakes observed in the wild

Here are some more mistakes we encountered in the wild:

## 🤦🏻‍♀️ Ambiguous Constructs

### ❌ Wrong
~~~
_for-sale.example.com. 1800 IN TXT "v=FORSALE1;furi=https://example.net/forsale;fval=EUR1000"
~~~

### ✅ Right

~~~
_for-sale.example.com. 1800 IN TXT "v=FORSALE1;furi=https://example.net/forsale"
_for-sale.example.com. 1800 IN TXT "v=FORSALE1;fval=EUR1000"
~~~

## ⚒ WORK IN PROGRESS - PLEASE CHECK BACK LATER

