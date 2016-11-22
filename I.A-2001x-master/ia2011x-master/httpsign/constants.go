package httpsign

const MaxSkewSec = 5                      // 5 sec
const CacheTimeout = 100                  // 100 sec
const CacheCapacity = 5000 * CacheTimeout // 5,000 msg/sec * 100 sec = 500,000 elements

const XMailgunSignature = "X-ia2001-Signature"
const XMailgunSignatureVersion = "X-ia2001-Signature-Version"
const XMailgunNonce = "X-ia2001-Nonce"
const XMailgunTimestamp = "X-ia2001-Timestamp"
