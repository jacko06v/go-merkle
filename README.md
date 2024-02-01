# go-merkle

## create a merkle tree with GO lang to be used with solidity

# this is not useful itself since i used and built it for a specific purpose in a project, but it can be useful as a reference for creating merkle trees in go

### scripts

- convert_int_string.js - convert all the numbers in json to strings
- checkHash.js - check the hash of the merkle tree if is the same in JS(and solidity)
- sampleCOntract.sol - solidity contract to create the nodes and check the hash
- merkle.go - create the merkle tree and save it to a json file

### merkle.go take a json file like

```json
[
  {
    "int_it": "631359070595067400",
    "renting": "3"
  },
  {
    "int_it": "631507570671499300",
    "renting": "5"
  },
  {
    "int_it": "630950076210238000",
    "renting": "5"
  },
  {
    "int_it": "631254397440568800",
    "renting": "3"
  },
  {
    "int_it": "633917065299962900",
    "renting": "1"
  }
]
```

### and create a json file like

```json
{
  "root": "0x3097752e39f8dfa07e1715d4876d4e8439e2a26332f8cd19060b3a7f3029d038",
  "proof": [
    {
      "leaf": "0x8cc628095351985de68ca79a948c500a851880ab495a54dbee7ff84c9c9195ba",
      "proof": [
        "0x9804d704fd9690226772119d69e97b6ac4a4e028a9d1a90d8ea04e91d036ea46",
        "0x0caaaa27f3ddeedd3bfc1c480a0e0a010654e6e3a21e9c1e98089608ed3d8503",
        "0x014a12586340e71b17316863b9e613def46b85632679cddd2bfffc5c026cf51d"
      ],
      "clearData": ["631359070595067391", "3"]
    },
    {
      "leaf": "0xbbc7b3f7222badbaeecce2d69bc704dee2d97181cc6204316465706628606704",
      "proof": [
        "0xc9b94f6c48fe9d7646961cf0c682b04422d503389be67d54b949e2c2dc7e2399",
        "0x344911bd6afd524f0cf2eae75c82499cfc9a7789321fa7744c26218114f2107e",
        "0x3bbb19b48564eae693e1eb5edd95f33974c34a0e5dc70b6f69d152994801fa22"
      ],
      "clearData": ["631507570671499263", "5"]
    },
    {
      "leaf": "0x027fe8442ae2b2d8e5224891a798c2342ffd71985197e7edaa2c9c904c396a6f",
      "proof": [
        "0x78c42edb06e6ffe9af47ca5281eb55d53cab022a2c8cb68a73ca005a53501da1",
        "0xb8e7329263a4a61c352416c71bbf808f3125451be027f95a6fe79e1d5d25aec4",
        "0x014a12586340e71b17316863b9e613def46b85632679cddd2bfffc5c026cf51d"
      ],
      "clearData": ["630950076210237951", "5"]
    },
    {
      "leaf": "0x78c42edb06e6ffe9af47ca5281eb55d53cab022a2c8cb68a73ca005a53501da1",
      "proof": [
        "0x027fe8442ae2b2d8e5224891a798c2342ffd71985197e7edaa2c9c904c396a6f",
        "0xb8e7329263a4a61c352416c71bbf808f3125451be027f95a6fe79e1d5d25aec4",
        "0x014a12586340e71b17316863b9e613def46b85632679cddd2bfffc5c026cf51d"
      ],
      "clearData": ["631254397440568831", "3"]
    },
    {
      "leaf": "0xc9b94f6c48fe9d7646961cf0c682b04422d503389be67d54b949e2c2dc7e2399",
      "proof": [
        "0xbbc7b3f7222badbaeecce2d69bc704dee2d97181cc6204316465706628606704",
        "0x344911bd6afd524f0cf2eae75c82499cfc9a7789321fa7744c26218114f2107e",
        "0x3bbb19b48564eae693e1eb5edd95f33974c34a0e5dc70b6f69d152994801fa22"
      ],
      "clearData": ["633917065299962879", "1"]
    }
  ]
}
```

