const std = @import("std");

pub const ChessboardError = error{
    IndexOutOfBounds,
};

pub fn square(index: usize) ChessboardError!u64 {
    if (index > 64 or index < 1) {
        return ChessboardError.IndexOutOfBounds;
    }
    std.debug.print("square = {d}\n", .{index - 1});
    return std.math.pow(u64, 2, index - 1);
}

pub fn total() u64 {
    return std.math.pow(u64, 2, 64) - 1;
}
