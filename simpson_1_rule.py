import math


def simpson_1_rule(f, a, b, n):
    h = (b - a) / n
    sum = f(a) + f(b)
    for i in range(1, n):
        if i % 2 == 1:
            sum += 4 * f(a + i * h)
        else:
            sum += 2 * f(a + i * h)
    return sum * h / 3


def someF(x):
    return math.e / (1 + x)


print(simpson_1_rule(someF, 0, 6, 20))
