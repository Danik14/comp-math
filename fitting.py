import numpy as np
import matplotlib.pyplot as plt


def polyfit(x, y, degree):
    results = {}
    coeffs = np.polyfit(x, y, degree)
    results["polynomial"] = coeffs.tolist()
    # r-squared
    p = np.poly1d(coeffs)
    yhat = p(x)
    ybar = np.sum(y) / len(y)
    ssreg = np.sum((yhat - ybar) ** 2)
    sstot = np.sum((y - ybar) ** 2)
    results["determination"] = ssreg / sstot
    return results


x = np.array([75, 80, 93, 65, 87, 71, 98, 68, 84, 77])
y = np.array([82, 78, 86, 72, 91, 80, 95, 72, 89, 74])

print(polyfit(x, y, 2))
