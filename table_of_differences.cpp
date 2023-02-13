//program for making
//a table of differences

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
const ll linf = (ll)1e17;

int f(int x) {
    return (x*x*x) + 5*x - 7;
}

void solve() {
    cout << "Enter the range: \n";
    int l, r; cin >> l >> r;
    vector<int> cur;
    for (int i = l; i <= r; ++i) {
        cur.pb(f(i));
    }

    while (!cur.empty()) {
        for (int x : cur) cout << x << " ";
        if (cur.size() == 1) { cur.clear(); }
        else {
            vector<int> nxt;
            for (int i = 0; i < sz(cur) - 1; ++i) { nxt.pb(cur[i + 1] - cur[i]); }
            cur = nxt;
        }
        cout << "\n";
    }
}



int main() {
    int tt = 1; //cin >> tt;
    while (tt--) solve();
    return 0;
}

/*

*/
