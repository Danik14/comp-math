import math
import numpy

lsm = dict()
lsm["x"] = 0
lsm["y"] = 0
lsm["x2"] = 0
lsm["x3"] = 0
lsm["x4"] = 0
lsm["xy"] = 0
lsm["x2y"] = 0


def getSums(xArr, yArr):

    for i in range(len(xArr)):
        lsm["x"] += xArr[i]
        lsm["y"] += yArr[i]
        lsm["x2"] += math.pow(xArr[i], 2)
        lsm["x3"] += math.pow(xArr[i], 3)
        lsm["x4"] += math.pow(xArr[i], 4)
        lsm["xy"] += xArr[i]*yArr[i]
        lsm["x2y"] += math.pow(xArr[i], 2)*yArr[i]


def printSums(lsm):
    print("%f\t%f\t%f\t%f\t%f\t%f\t%f" % (
        lsm["x"], lsm["y"], lsm["x2"], lsm["x3"], lsm["x4"], lsm["xy"], lsm["x2y"]))
# def printTable(table):
#     print("x\ty\tx2\tx3\tx4\txy\tx2y")
#     for i in range(len(table["x"])):
#         print("%f\t%f\t%f\t%f\t%f\t%f\t%f" % (
#             lsm["x"], lsm["y"], lsm["x2"], lsm["x3"], lsm["x4"], lsm["xy"], lsm["x2y"]))


xArr = [1, 2, 3, 4, 5, 6, 7, 8, 9]
yArr = [2, 6, 7, 8, 10, 11, 11, 10, 9]


# xArr = [1, 2, 3, 4, 5, 6, 7]
# yArr = [-5, -2, 5, 16, 31, 50, 73]
getSums(xArr, yArr)

printSums(lsm)

# sum(y) = an+b*sum(x)+c*sum(x2)
# sum(xy) = a*sum(x)+b*sum(x2)+c*sum(x3)
# sum(x2y) = a*sum(x2)+b*sum(x3)+c*sum(x4)
print("%da+%db+%dc = %d" % (len(xArr), lsm["x"], lsm["x2"], lsm["y"]))
print("%da+%db+%dc = %d" % (lsm["x"], lsm["x2"], lsm["x3"], lsm["xy"]))
print("%da+%db+%dc = %d" % (lsm["x2"], lsm["x3"], lsm["x4"], lsm["x2y"]))

a = numpy.array([[len(xArr), lsm["x"], lsm["x2"]],
                 [lsm["x"], lsm["x2"], lsm["x3"]],
                 [lsm["x2"], lsm["x3"], lsm["x4"]]])
b = numpy.array([lsm["y"], lsm["xy"], lsm["x2y"]])

x = numpy.linalg.solve(a, b)
print(x)
