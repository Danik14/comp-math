from numpy import array
from scipy.linalg import funm
m = array([[5, -2],
           [1, 2]])
tmp = funm(m, lambda x: x**2 - 7*x + 12)

print(tmp)
