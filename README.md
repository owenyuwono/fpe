# Golang FPE (Format Preserving Encryption)
## Theory
Format preserving encryption (FPE) refers to a set of techniques for encrypting data such that the ciphertext has the same format as the plaintext. For instance, you can use FPE to encrypt credit card numbers with valid checksums such that the ciphertext is also an credit card number with a valid checksum, or similarly for bank account numbers, US Social Security numbers, or even more general mappings like English words onto other English words.

To encrypt an arbitrary value using FE1, you need to use a ranking method. Basically, the idea is to assign an integer to every value you might encrypt. For instance, a 16 digit credit card number consists of a 15 digit code plus a 1 digit checksum. So to encrypt a credit card number, you first remove the checksum, encrypt the 15 digit value modulo 1015, and then calculate what the checksum is for the new (ciphertext) number. Or, if you were encrypting words in a dictionary, you could rank the words by their lexicographical order, and choose the modulus to be the number of words in the dictionary.

## Implementation
Current implementation uses the FE1 scheme from the paper "Format-Preserving Encryption" by Bellare, Rogaway, et al.

Ported from [node-fe1-fpe](https://github.com/eCollect/node-fe1-fpe) which was ported from [java-fpe](https://github.com/Worldpay/java-fpe) which was ported from [DotFPE](https://dotfpe.codeplex.com/) which was ported from [Botan Library](http://botan.randombit.net/).

## Installation

```
go get github.com/owenyuwono/fpe
```

## Basic usage
```go
package main

import (
    "github.com/owenyuwono/fpe"
)

func main() {
    encrypted, err := fpe.Encrypt(10001, 1, "my-secret-key", "my-non-secret-tweak", 3)
    if err != nil {
        panic(err)
    }
    fmt.Println(encrypted) // 5011
}
```

## Considerations

The implementation is as stable as a rock for a modulus up to 10 000 000. It is designed this way because of speed concerns. For larger range, the matter needs to be discussed with the corresponding developers.

## Todo

- [ ] Decrypt function
- [ ] More verbose documentation

## License

Copyright © 2021 owenyuwono. Licensed under the MIT license.