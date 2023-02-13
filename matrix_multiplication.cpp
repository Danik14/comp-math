//program for multiplying
//2 matrices


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
const int mod = (int)1e9 + 7;
const ll linf = (ll)1e17;

void solve() {
    cout << "Enter the first matrix: \n";
    int n1, m1; cin >> n1 >> m1;
    vector<vector<int>> a(n1, vector<int>(m1));
    for (int i = 0; i < n1; ++i) {
        for (int j = 0; j < m1; ++j) {
            cin >> a[i][j];
        }
    }
    cout << "Enter the second matrix: \n";
    int n2, m2; cin >> n2 >> m2;
    vector<vector<int>> b(n2, vector<int>(m2));
    for (int i = 0; i < n2; ++i) {
        for (int j = 0; j < m2; ++j) {
            cin >> b[i][j];
        }
    }
    if (m1 != n2) {
        cout << "Dimensions are not appropriate\n";
        return;
    }
    vector<vector<int>> res(n1, vector<int>(m2, 0));
    for (int i = 0; i < n1; ++i) {
        for (int j = 0; j < m2; ++j) {
            for (int k = 0; k < m1; ++k) {
                res[i][j] += a[i][k] * b[k][j];
            }
        }
    }

    cout << "RESULT:\n";
    for (int i = 0; i < n1; ++i) {
        for (int j = 0; j < m2; ++j) {
            cout << res[i][j] << " ";
        }
        cout << "\n";
    }
}


int main() {
    int tt = 1; //cin >> tt;
    while (tt--) solve();
    return 0;
}