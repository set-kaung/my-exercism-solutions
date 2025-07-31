const std = @import("std");

pub fn isIsogram(str: []const u8) bool {
    var bits = std.bit_set.IntegerBitSet(26).initEmpty();
    for (str) |ch| {
        if (std.ascii.isAlphabetic(ch)) {
            if (bits.isSet(std.ascii.toLower(ch) - 'a')) {
                return false;
            } else {
                bits.set(std.ascii.toLower(ch) - 'a');
            }
        }
    }
    return true;
}
