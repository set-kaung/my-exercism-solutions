pub const Classification = enum {
    deficient,
    perfect,
    abundant,
};

/// Asserts that `n` is nonzero.
pub fn classify(n: u64) Classification {
    var sum: u64 = 0;
    for (1..n) |i| {
        if (n % i == 0 and i != n) {
            sum += i;
        }
    }
    if (n == sum) {
        return Classification.perfect;
    } else if (n < sum) {
        return Classification.abundant;
    } else {
        return Classification.deficient;
    }
}
