import copy

a = [3, [5, 76], 17]
b = a
c = a[:]
d = list(a)
e = a.copy()
f = copy.deepcopy(a)
print(a,b,c,d,e)
a[1][0] = 100
print(a,b,c,d,e)  # 값이 다른 복사한 요소들도 바뀜