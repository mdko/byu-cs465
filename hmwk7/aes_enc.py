from Crypto import Random   # https://www.dlitz.net/software/pycrypto/
from Crypto.Cipher import AES


def main():
    infile = open('plaintext.txt')
    inmsg = infile.read()
    # This AES is AES-128 and thus a 16-byte key

    ############################################################
    # ECB, CBC plaintext length must be a multiple of block_size
    #
    # pad the message
    extraend = len(inmsg) % AES.block_size
    padmsg = inmsg + '0'*(AES.block_size - extraend)

    # ECB
    key = Random.get_random_bytes(16)
    cipherecb = AES.new(key, AES.MODE_ECB)
    bloc = 0
    outmsgecb = ''
    # no initilization vector needed
    # input 16 byte chunks into ECB encryption
    while bloc < len(padmsg):
        outmsgecb += cipherecb.encrypt(padmsg[bloc : bloc + AES.block_size])
        bloc += AES.block_size
    output = ''.join(x.encode('hex') for x in outmsgecb)
    print 'Output for AES using ECB:\n %s' % output

    # CBC
    key = Random.get_random_bytes(16)
    iv = Random.new().read(AES.block_size)
    ciphercbc = AES.new(key, AES.MODE_CBC, iv)
    bloc = 0
    outmsgcbc = ''
    # no initilization vector needed
    # input 16 byte chunks into ECB encryption
    while bloc < len(padmsg):
        outmsgcbc += ciphercbc.encrypt(padmsg[bloc : bloc + AES.block_size])
        bloc += AES.block_size
    output = ''.join(x.encode('hex') for x in outmsgcbc)
    print 'Output for AES using CBC:\n %s' % output

    
    ############################################################
    # CFB plaintext length must be a multiple of segment_size/8
    #
    key = Random.get_random_bytes(16)
    iv = Random.new().read(AES.block_size)
    ciphercfb = AES.new(key, AES.MODE_CFB, iv, segment_size=16)
    bloc = 0
    outmsgcfb = ''
    while bloc < len(inmsg) - AES.block_size: # or <=
        outmsgcfb += ciphercfb.encrypt(inmsg[bloc : bloc + AES.block_size])
        bloc += AES.block_size
    outmsgcfb += ciphercfb.encrypt(inmsg[bloc :])
    output = ''.join(x.encode('hex') for x in outmsgcfb)
    print 'Output for AES using CFB:\n %s' % output     # should be same size as original message, since no padding occured


if __name__ == "__main__":
    main()