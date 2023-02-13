import math

globalArr = []


def forward(arr):
    newArr = []
    for i in range(1, len(arr)):
        newArr.append(arr[i]-arr[i-1])
    globalArr.append(newArr[0])
    if len(newArr) > 1:
        forward(newArr)
    else:
        return newArr


def substituion(arr, n):
    substitution = arr[0]
    for i in range(1, len(arr)):
        nVal = 1
        for j in range(0, i):
            nVal *= (n-j)
        substitution += (arr[i]*nVal)/math.factorial(i)
    print("substitution: ", substitution)


# h = x1-x0
h = 0.1

# necessary to find f(2.75), so x = 2.75
# n = (x-x0)/h
n = (0.18-0)/h

# array of falues of y
arrY = [1, 1.052, 1.2214, 1.3499, 1.4918]
globalArr.append(arrY[0])


forward(arrY)
# print(globalArr)
substituion(globalArr, n)
