# import numpy as np
# from scipy import optimize
# import matplotlib.pyplot as plt

# plt.style.use("seaborn-poster")

# # generate x and y
# x = np.linspace(0, 1, 101)
# y = 1 + x + x * np.random.random(len(x))

# # assemble matrix A
# A = np.vstack([x, np.ones(len(x))]).T

# # turn y into a column vector
# y = y[:, np.newaxis]

# # Direct least square regression
# alpha = np.dot((np.dot(np.linalg.inv(np.dot(A.T, A)), A.T)), y)
# print(alpha)


# # plot the results
# plt.figure(figsize=(10, 8))
# plt.plot(x, y, "b.")
# plt.plot(x, alpha[0] * x + alpha[1], "r")
# plt.xlabel("x")
# plt.ylabel("y")
# plt.show()


def calc_parabola_vertex(x1, y1, x2, y2, x3, y3):
    """
    Adapted and modifed to get the unknowns for defining a parabola:
    http://stackoverflow.com/questions/717762/how-to-calculate-the-vertex-of-a-parabola-given-three-points
    """

    denom = (x1 - x2) * (x1 - x3) * (x2 - x3)
    A = (x3 * (y2 - y1) + x2 * (y1 - y3) + x1 * (y3 - y2)) / denom
    B = (x3 * x3 * (y1 - y2) + x2 * x2 * (y3 - y1) + x1 * x1 * (y2 - y3)) / denom
    C = (
        x2 * x3 * (x2 - x3) * y1 + x3 * x1 * (x3 - x1) * y2 + x1 * x2 * (x1 - x2) * y3
    ) / denom

    return A, B, C


print(calc_parabola_vertex(2, 6.07, 4, 12.85, 6, 31.47))
