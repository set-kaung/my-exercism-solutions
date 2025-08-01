pub const DnaError = error{
    EmptyDnaStrands,
    UnequalDnaStrands,
};
pub fn compute(first: []const u8, second: []const u8) DnaError!usize {
    if (first.len == 0 or second.len == 0) {
        return DnaError.EmptyDnaStrands;
    }
    if (first.len != second.len) {
        return DnaError.UnequalDnaStrands;
    }
    var distance: usize = 0;
    for (first, second) |f, s| {
        if (f != s) {
            distance += 1;
        }
    }
    return distance;
}
