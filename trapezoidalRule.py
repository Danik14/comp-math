# program for solving
# with trapezoidal rule
import math


def trapezoidal_rule(f, a, b, n):
    h = (b - a) / n
    sum = 0.5 * (f(a) + f(b))
    for i in range(1, n):
        sum += f(a + i * h)
    return sum * h


def someF(x):
    # return math.e**x * x
    return 0.42857142857142894 * x**2 - 0.9142857142857153 * x + 1.8571428571428583


print(trapezoidal_rule(someF, 0, 4, 200))
