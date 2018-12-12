N=21888242871839275222246405745257275088548364400416034343698204186575808495617

def hash(x, y):
  t=x*x % N
  x=(t*t*x+y) % N
  t=x*x % N
  x=(t*t*x+y) % N
  t=x*x % N
  x=(t*t*x+y) % N
  t=x*x % N
  x=(t*t*x+y) % N
  t=x*x % N
  x=(t*t*x+y) % N
  return x

#def main(field proofSliceBegin, field proofSliceEnd, private field[4] rootHash, private field[4] sliceBegin, 
# private field[4] sliceEnd, private field[128] proofHash, private field[128] proofLength, private field[128] proofBit) -> (field):

from random import randint

def genRandProof(depth, x, y):
    index=0
    item = 0
    begin = x-randint(0, 2**12-1)
    end = begin+y+randint(0, 2**12-1)
    proofBit=[]
    proofLength=[]
    proofHash=[]
    curItem = item
    curLength = end-begin
    curLeft = begin
    for i in range(0, depth):
        b = randint(0, 1)
        ci = randint(0, N-1)
        cl = randint(0, 2**12-1)
        if i == (depth - 2):
            b = 1
            cl = curLeft
        if i == (depth - 1):
            b = 0
            cl = 2**24-1-curLength 
        proofBit += [b]
        proofLength += [cl]
        proofHash += [ci]
        if b == 1:
            curItem= hash(hash(hash(cl, curLength), ci), curItem)
            curLeft -= cl
            curLength +=cl
        else:
            curItem = hash(hash(hash(curLength, cl), curItem), ci)
            curLength +=cl
    return curItem, begin, end, proofHash, proofLength, proofBit

x = randint(2**21, 2**22)
y = x + 100

res = []
for i in range(0,4):
  res+=[genRandProof(16, x, y)]


res2 = [x, y]
for j in range(0,3):
  for i in range(0,4):
    res2 += [res[i][j]]

for j in range(3,6):
  for i in range(0,4):
    res2 +=res[i][j]

for t in res2:
  if t>=N:
    print("error1")
  if t<0:
    print("error2", t)

print(" ".join(map(lambda x:str(x), res2)))

print("")
print(len(res2))
print(hash(hash(hash(hash(0, res[0][0]), res[1][0]), res[2][0]), res[3][0]))
