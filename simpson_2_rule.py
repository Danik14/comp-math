import math


def simpson_2_rule(f, a, b, n):
    if n % 3 != 0:
        raise ValueError("Number of subintervals must be a multiple of 3")

    h = (b - a) / n
    result = f(a) + f(b)

    for i in range(1, n, 3):
        result += 3 * f(a + i * h) + 3 * f(a + (i + 1) * h) + 2 * f(a + (i + 2) * h)

    return (3 * h / 8) * result


def someF(x):
    # return math.e / (1 + x)
    # return 1 / (1 + math.cos(x))
    return 1 / (1 + x**2)


print(simpson_2_rule(someF, 0, 6, 6))
