# def detOf2x(arr):
#     return arr[0][0] * arr[1][1] - arr[0][1] * arr[1][0]


# def detOf3x(arr):
#     return (
#         arr[0][0] * arr[1][1] * arr[2][2]
#         + arr[0][1] * arr[1][2] * arr[2][0]
#         + arr[0][2] * arr[1][0] * arr[2][1]
#     ) - (
#         arr[2][0] * arr[1][1] * arr[0][2]
#         + arr[2][1] * arr[1][2] * arr[0][0]
#         + arr[2][2] * arr[1][0] * arr[0][1]
#     )


# arr = [[3, -2], [1, 0]]
# iArr = [[1, 0], [0, 1]]


# def multiplyMatrices(arr1, arr2):
#     # result = [len(arr1)][len(arr2[0])]
#     rows = len(arr1)
#     cols = len(arr2[0])
#     result = [[0 for rows in range(rows)] for cols in range(cols)]
#     for i in range(len(arr1)):
#         for j in range(len(arr2[0])):
#             for k in range(len(arr2)):
#                 result[i][j] += arr1[i][k] * arr2[k][j]
#     return result


# # print(detOf2x([[50, 29], [30, 44]]), "\n")
# # print(detOf3x([[55, 25, 15],
# #                [30, 44, 2],
# #                [11, 45, 77]]))
# newArr = multiplyMatrices([[2, 6, 4], [1, 9, 7]], [[3, 2], [4, 10], [11, 8]])
# for i in newArr:
#     print(i)


import numpy as np


def eigenvalues_and_vectors(matrix):
    eigenvalues, eigenvectors = np.linalg.eig(matrix)
    return eigenvalues, eigenvectors


if __name__ == "__main__":
    matrix = np.array([[8, -4, 2], [4, 0, 2], [0, -2, -4]])
    eigenvalues, eigenvectors = eigenvalues_and_vectors(matrix)
    print("Eigenvalues:", eigenvalues)
    print("Eigenvectors:", eigenvectors)
