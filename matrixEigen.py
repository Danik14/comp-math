#program for finding
#eigenvalue and eigenvector of a matrix

import numpy as np


def normalize(x):
    fac = abs(x).max()
    x_n = x / x.max()
    return fac, x_n


x = np.array([1, 1])
a = np.array([[1, 2], [3, 4]])

for i in range(20):
    x = np.dot(a, x)
    lambda_1, x = normalize(x)

print("Eigenvalue:", lambda_1)
print("Eigenvector:", x)
