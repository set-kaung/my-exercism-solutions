const std = @import("std");

pub fn isArmstrongNumber(num: u128) bool {
    if (num < 10) {
        return true;
    }
    var digitCount: u128 = 0;
    var numCopy: u128 = num;
    while (numCopy > 0) : (numCopy /= 10) {
        digitCount += 1;
    }

    var total: u128 = 0;
    numCopy = num;
    while (numCopy > 0) {
        total += std.math.pow(u128, numCopy % 10, digitCount);
        numCopy = numCopy / 10;
    }
    return total == num;
}
