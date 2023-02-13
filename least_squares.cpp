//program for calculation
//using least squares method

#include <bits/stdc++.h>

#define pb push_back
#define sz(x) (int)x.size()
#define whole(x) x.begin(), x.end()
#define rwhole(x) x.rbegin(), x.rend()

using namespace std;

typedef long long ll;
typedef long double ld;

const int N = (int)3e5 + 25;
const int maxN = (int)1e5 + 2;
const int INF = (int)1e9 + 7;
const double EPS = 1e-7;
const ll linf = (ll)1e17;

double get(int n, double sumX, double sumY, double sumXY, double sumXX) {
    return (n * 1.0 * sumXY - sumX * sumY) / (n * sumXX - (sumX * sumX));
}

double get(int n, double m, double sumY, double sumX) {
    if (sumY - m * sumX < EPS) return 0.0;
    return (sumY - m * sumX) / n;
}

void solve() {
    cout << "Enter the number of points: \n";
    int n; cin >> n;
    cout << "Enter the points:\n";
    vector<double> x(n), y(n);
    for (int i = 0; i < n; ++i) {
        cin >> x[i] >> y[i];
    }

    vector<vector<double>> t(n, vector<double> (4));
    double sumX = 0, sumY = 0, sumXY = 0, sumXX = 0;
    for (int j = 0; j < 4; ++j) {
        if (j == 0) {
            for (int i = 0; i < n; ++i) t[i][j] = x[i], sumX += x[i];
        }
        else if (j == 1) {
            for (int i = 0; i < n; ++i) t[i][j] = y[i], sumY += y[i];
        }
        else if (j == 2) {
            for (int i = 0; i < n; ++i) t[i][j] = x[i] * y[i], sumXY += t[i][j];
        }
        else if (j == 3) {
            for (int i = 0; i < n; ++i) t[i][j] = x[i] * x[i], sumXX += t[i][j];
        }
    }
    cout << "The table is;\n";
    cout << "X  Y  XY  X^2\n";
    for (int i = 0; i < n; ++i) {
        for (int j = 0; j < 4; ++j) {
            cout << t[i][j] << " ";
        }
        cout << "\n";
    }

    printf ("SUMS ARE: SUMX: %lf SUMY: %lf SUMXY: %lf SUMX^2: %lf\n", sumX, sumY, sumXY, sumXX);

    double slope = get(n, sumX, sumY, sumXY, sumXX);
    cout << "The slope is: " << slope << "\n";
    cout << "Y - inter: " << get(n, slope, sumY, sumX);
}

int main() {
    int tt = 1;
    while (tt--) solve();
    return 0;
}