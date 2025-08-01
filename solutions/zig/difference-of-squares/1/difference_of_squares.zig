const std = @import("std");

pub fn squareOfSum(number: usize) usize {
    return std.math.pow(usize, (number * (number + 1)) / 2, 2);
}

pub fn sumOfSquares(number: usize) usize {
    var total: usize = 0;
    var i: usize = 0;
    while (i <= number) : (i += 1) {
        total += std.math.pow(usize, i, 2);
    }
    return total;
}

pub fn differenceOfSquares(number: usize) usize {
    return squareOfSum(number) - sumOfSquares(number);
}
