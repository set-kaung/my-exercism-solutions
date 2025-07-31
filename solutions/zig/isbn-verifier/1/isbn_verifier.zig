const std = @import("std");

pub fn isValidIsbn10(s: []const u8) bool {
    var total: u16 = 0;
    var count: u8 = 10;
    for (s, 0..) |ch, l| {
        if (std.ascii.isAlphabetic(ch) and std.ascii.toLower(ch) != 'x') {
            return false;
        }
        if (count == 0 and l != s.len) {
            return false;
        }
        if (count == 0) {
            break;
        }
        if (std.ascii.toLower(ch) == 'x' and count == 1) {
            total += count * 10;
            count -= 1;
        } else if (std.ascii.isDigit(ch)) {
            total += count * (ch - '0');
            count -= 1;
        }
    }
    return total % 11 == 0 and count == 0;
}
