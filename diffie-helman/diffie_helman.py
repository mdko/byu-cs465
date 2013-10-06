import OpenSSL as ssl
import gensafeprime as sp

def modular_exponentiation(base, exp, mod):
    res = 1
    while (exp > 0):
        if (exp % 2 == 1):
            res = (res * base) % mod
        exp = exp >> 1
        base = (base * base) % mod
    return res

def diffie_helman():
    g = 5
    rbytes = ssl.rand.bytes(63)
    s = long(rbytes.encode('hex'), 16)
    p = sp.generate(504)
    print "g=%d, s=%d, p=%d" % (g, s, p)
    print "g^s mod p =%d" % (modular_exponentiation(g, s, p))

if __name__ == "__main__":
    diffie_helman()