import OpenSSL as ssl
import gensafeprime as sp

e = 65537

def modular_exponentiation(base, exp, mod):
    res = 1
    while (exp > 0):
        if (exp % 2 == 1):
            res = (res * base) % mod
        exp = exp >> 1
        base = (base * base) % mod
    return res


# a is phi, b is e
def multiplicative_inverse(a, b):    
    # The coefficient y of e will be the 
    # multiplicative inverse of e mod phi(n).
    # row layout = [x, y, d, k], we pull out y
    row = extended_euclidean(a, b)
    y = row[1]
    if y < 0:
        y = y + a
    return y


def relatively_prime(a, b):
    row = extended_euclidean(a, b)
    return row[2] == 1 # checking to see if GCD[via extended_euclidean] is 1


# aka GCD
# uses the Table Method of the Extended Euclidean Algorithm
def extended_euclidean(a, b):
    row_minus_one = [1, 0, a, 0]
    row = [0, 1, b, a/b]
    while row[2] > 1:
        k = row[3]
        nr0 = row_minus_one[0] - k*row[0]
        nr1 = row_minus_one[1] - k*row[1]
        nr2 = row_minus_one[2] - k*row[2]
        nk = row[2]/nr2

        row_minus_one = row
        row = [nr0, nr1, nr2, nk]
    return row


def get_primes():
    rel_prime = False
    p, q = 0, 0
    
    while not rel_prime:
        # Ensure high order bit is set
        while (p >> 511) != 1: # or could bit-wise & with 2^512 and check if equal to 2^512
            p = sp.generate(512)
        while (q >> 511) != 1:
            q = sp.generate(512)

        # Verify that phi(n) [(p-1)(q-1) is relatively prime to e]
        phin = (p - 1) * (q - 1)
        if relatively_prime(phin, e) == 1: # or just phin % e != 0, since e is prime
            rel_prime = True
        else:
            p, q = 0, 0

    return p, q


def encrypt_message(m, n):
    return modular_exponentiation(m, e, n)

def decrypt_message(m, d, n):
    return modular_exponentiation(m, d, n)


def test_encryption_decryption(n, d):
    for i in range(0, 10):
        mStr = ssl.rand.bytes(64)
        m = long(mStr.encode('hex'), 16) >> 1 # since n has right bit high, shifting by 1 ensures m is less than n, 512 bits = 64 bytes
        result = modular_exponentiation(modular_exponentiation(m, e, n), d, n)
        assert(m == result)


def rsa():
    p, q = get_primes()
    n = p * q
    phin = (p - 1) * (q - 1)
    d = multiplicative_inverse(phin, e)
    test_encryption_decryption(n, d)


if __name__ == "__main__":
    rsa()