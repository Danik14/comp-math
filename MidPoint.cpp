//program for calculation
//of middle point (integral)

#include <bits/stdc++.h>

#define pb push_back
#define sz(x) (int)x.size()
#define whole(x) x.begin(), x.end()
#define rwhole(x) x.rbegin(), x.rend()

using namespace std;

typedef long long ll;
typedef long double ld;

const int N = (int)5e5 + 25;
const int maxN = (int)1e5 + 2;
const int INF = (int)1e9 + 7;
const int mod = (int)1e9 + 7;
const ll linf = (ll)1e17;

double f(double x) {
    return x * x * x; //write your f(x) function here
}

void solve() {
    cout << "Enter the precision of error: \n";
    int err; cin >> err;
    cout << "Enter the intervals: \n";
    double a, b; cin >> a >> b; // the intervals of the integral
    cout << "Enter the number of rectangles you wanna divide integral\n";
    int n; cin >> n;
    double delta = (b - a) / n; // range between two rectangles
    double sum = 0;
    for (double point = a; point < b; point += delta) {
        double nxtPoint = point + delta; // nxtPoint is the next point so we can calculate middle point
        double midPoint = (nxtPoint + point) / 2.0;
        sum += f(midPoint); 
    }
    double expectedIntegral = delta * sum;
    cout << fixed << setprecision(err) << expectedIntegral << "\n";
}


int main() {
    int tt = 1;
    while (tt--) solve();
    return 0;
}
