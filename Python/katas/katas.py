from preloaded import NATO # NATO['A'] == 'Alfa', etc

def to_nato_first_cut(words : str) -> str:
    result = []
    for c in words:
        if c.isalpha():
            result.append(NATO[c.upper()])
        elif c != ' ':
            result.append(c)
    return " ".join(result)

def to_nato(words : str) -> str:
    return ' '.join([ NATO[c.upper()] if c.isalpha() else c for c in words if c != ' '])
    

#
# 
def move_zeros(lst):
    newArray = [n for n in lst if n != 0]
    newArray.extend([0] * lst.count(0))
    return newArray

def move_zeros_clever(array):
    return [x for x in array if x] + [0]*array.count(0)

# <5 kyu> Hamster me
# https://www.codewars.com/kata/595ddfe2fc339d8a7d000089/train/python
def hamster_me(code, message):
    offset = 1
    encoder= {}
    toEncode = [c for c in code]
    while len(toEncode) > 0:
        nextList = []
        for c in toEncode:
            if c not in encoder:
                cval = c
                if offset > 1:
                    cval = chr(((( ord(c[0]) - offset + 27 ) - ord('a') ) % 26 ) + ord('a'))
                encoder[c] = f"{cval}{offset}"
                nextList.append( 'a' if c == 'z' else chr(ord(c)+1))       
        toEncode = nextList
        offset += 1

    return ''.join ( [ encoder[c] for c in message] )

# <5 kyu> Endianness Conversion
# https://www.codewars.com/kata/56f2dd31e40b7042ad001026/train/python
def switch_endian(n, bits):
    # easy cases: n must be positive
    # bits must be power of 2, greater than or equal to 8
    if n < 0:
        return None
    if bits < 8 or (bits & (bits-1)):
        return None
    # n must be less than 2**bits
    if n >= 2**bits:
        return None

    # rearrange the bytes:
    nBytes = bits // 8
    tnew = 0
    for i in range(nBytes):
        tnew = (tnew << 8) + (n & 255)
        n = n >> 8
    return tnew

# <4 kyu> Twice linear
# https://www.codewars.com/kata/5672682212c8ecf83e000050/train/python
def dbl_linear(n):
    u = [1]
    i = 0
    j = 0

    while len(u) <= n:
        n1 = u[i]*2 + 1
        n2 = u[j]*3 + 1
        if n1 < n2:
            u.append(n1)
            i += 1
        elif n2 < n1:
            u.append(n2)
            j += 1
        else:
            u.append(n1)
            i += 1
            j += 1

    return u[n]