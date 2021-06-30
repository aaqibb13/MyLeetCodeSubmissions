N = 15
powered = 2**N
print(powered)
def power(num):
    ini = 0
    for i in str(num):
        ini = ini + int(i)
    return ini
    
res = power(powered)

if res >= 0 and res <= 9:
    print(res)
else:
    res = power(res)
    print(res)
