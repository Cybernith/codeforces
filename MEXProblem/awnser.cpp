#include <bits/stdc++.h>
using namespace std;

struct SegmentTree {
    int size;
    vector<int> value;
    vector<int> lazy;

    SegmentTree(int n) {
        size = n;
        int base = 1;
        while (base < size) base <<= 1;
        value.assign(base * 4, 0);
        lazy.assign(base * 4, 0);
    }

    void push(int index) {
        int delta = lazy[index];
        if (delta != 0) {
            int left = index * 2;
            int right = left + 1;
            value[left] += delta;
            value[right] += delta;
            lazy[left] += delta;
            lazy[right] += delta;
            lazy[index] = 0;
        }
    }

    void rangeAdd(int index, int left, int right,
                  int queryLeft, int queryRight, int delta) {
        if (queryLeft > right || queryRight < left) return;
        if (queryLeft <= left && right <= queryRight) {
            value[index] += delta;
            lazy[index] += delta;
            return;
        }
        push(index);
        int middle = (left + right) >> 1;
        rangeAdd(index * 2, left, middle, queryLeft, queryRight, delta);
        rangeAdd(index * 2 + 1, middle + 1, right, queryLeft, queryRight, delta);
        value[index] = max(value[index * 2], value[index * 2 + 1]);
    }

    void rangeAdd(int left, int right, int delta) {
        if (left > right) return;
        rangeAdd(1, 0, size - 1, left, right, delta);
    }

    void pointSet(int index, int left, int right,
                  int position, int newValue) {
        if (left == right) {
            value[index] = newValue;
            lazy[index] = 0;
            return;
        }
        push(index);
        int middle = (left + right) >> 1;
        if (position <= middle) {
            pointSet(index * 2, left, middle, position, newValue);
        } else {
            pointSet(index * 2 + 1, middle + 1, right, position, newValue);
        }
        value[index] = max(value[index * 2], value[index * 2 + 1]);
    }

    void pointSet(int position, int newValue) {
        pointSet(1, 0, size - 1, position, newValue);
    }

    int maxValue() const {
        return value[1];
    }
};

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int t;
    if (!(cin >> t)) {
        return 0;
    }
    while (t--) {
        int n;
        cin >> n;
        vector<int> a(n);
        for (int i = 0; i < n; i++) {
            cin >> a[i];
        }
        int coordinateCount = n + 1;
        SegmentTree tree(coordinateCount);
        for (int i = 0; i < n; i++) {
            int v = a[i];
            tree.pointSet(v, 0);
            if (v >= 1) {
                tree.rangeAdd(0, v - 1, 1);
            }
            cout << tree.maxValue();
            if (i + 1 < n) {
                cout << ' ';
            }
        }
        if (t) {
            cout << '\n';
        }
    }
    return 0;
}
