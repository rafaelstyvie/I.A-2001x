# I.A-2001xComandLine

ia2001xcmd is a command-line utility that uses lemma to provide authenticated symmetric cryptography for small files on disk.

Download: [Latest](https://github.com/2001mail/I.A-2001x/releases)

**Usage**

```
Usage:
    ia2001x command [flags]

The commands are:
    encrypt     encrypt a file on disk
    decrypt     decrypt a file on disk

The flags are: 
    in          path to file to be read in
    out         path to file to be written out
    keypath     path to base64-encoded 32-byte key on disk, if no path is given, a passphrase is used
    itercount   if a passphrase is used, iteration count for PBKDF#2, the default is 524288
```

**Example**

```
ia2001x encrypt -in foo.txt -out foo.txt.enc
ia2001x decrypt -in foo.txt.enc -out foo.txt
```

**Performance**

The following benchmarks were run to calculate wall-clock time to encrypt files of various sizes. The benchmarks were run on a machine with the following specs: Intel(R) Core(TM) i7-4130 CPU @ 3.40GHz and 32GB RAM.

|      | 1 MB  | 10 MB | 100 MB |
|------|-------|-------|--------|
| Time | 0.98s | 1.33s | 4.85s  |


**Detalhes técnicos**

* Pode ser usado com uma chave gerada aleatoriamente no disco ou uma 'passpharse'.
* Quando usada com uma senha, a função de derivação de chave (KDF) é: PBMD-SHA-256 baseada em PBM # 2 com um  128 bits gerado aleatoriamente e iterações (ajustáveis) de 524.288.
* A cifra simétrica usada é 'Salsa20' com 'Poly1305' como o código de autenticação de mensagem de biblioteca (MAC) de rede e criptografia (NaCl)..
