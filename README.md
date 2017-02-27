# urlgo

Basic tool for http requests. Inspired by ssl `gnutls_handshake() failed: Error in the pull function.
` issue with curl and wget on stock AWS Ubuntu 16.04 when requesting `https://sm.megafon.ru/`.

# usage


Could support more curl flags in future but currently takes just one argument, url:

```sh
urlgo http://httpbin.org/get
```
