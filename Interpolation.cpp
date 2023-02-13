// both Forward and Backward
// Newton's Interpolation
#include <stdio.h>
void forward(float x[4], float y[4][4], int n);
void BackWard(float x[4], float y[4][4], int n);
int main()
{
	int i, j;
	int n = 5; // number of arguments
	float x[5] = { 0, 2, 4, 6, 8};
	float y[5][5] = {
		{ 4, 0, 0, 0, 0 },
		{ 26, 0, 0, 0, 0 },
		{ 58, 0, 0, 0, 0 },
		{ 112, 0, 0, 0, 0 },
		{ 466, 0, 0, 0, 0}
	};
	float a = 5; // interpolation point
	float h, u, sum, p;
	for (j = 1; j < n; j++) {
		for (i = 0; i < n - j; i++) {
			y[i][j] = y[i + 1][j - 1] - y[i][j - 1];
		}
	}
	printf("\n The forward difference table is:\n");
	for (i = 0; i < n; i++) {
		printf("\n");
		for (j = 0; j < n - i; j++) {
			printf("%f\t", y[i][j]);
		}
	}

	p = 1.0;
	sum = y[0][0];
	h = x[1] - x[0];
	u = (a - x[0]) / h;
	for (j = 1; j < n; j++) {
		p = p * (u - j + 1) / j;
		sum = sum + p * y[0][j];
	}
	printf("\nThe value of y at x=%0.1f is %0.6f", a, sum);
	return 0;
}

/*
void BackWard(float x[5], float y[5][5], int n)
{
	int i, j;
	float a = 5; // interpolation point
	float h, u, sum, p;
	for (j = 1; j < n; j++) {
		for (i = j; i < n; i++) {
			y[i][j] = y[i][j - 1] - y[i - 1][j - 1];
		}
	}
	printf("\nThe backward difference table is:\n");
	for (i = 0; i < n; i++) {
		printf("\n");
		for (j = 0; j <= i; j++) {
			printf("%f\t", y[i][j]);
		}
	}

	p = 1.0;
	sum = y[n - 1][0];
	h = x[1] - x[0];
	u = (a - x[n - 1]) / h;
	for (j = 1; j < n; j++) {
		p = p * (u + j - 1) / j;
		sum = sum + p * y[n - 1][j];
	}

	printf("\nThe value of y at x=%0.1f is %0.3f", a, sum);
}
*/