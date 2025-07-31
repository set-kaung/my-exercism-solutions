const std = @import("std");

// fn toLower(str: []const u8) []u8 {
//     var newStr: []u8 = undefined;
//     std.mem.copyForwards(u8, newStr, str);
//     for (newStr, 0..) |ch, indice| {
//         if (ch >= 'A' and ch <= 'Z') {
//             newStr[indice] = 'a' + (ch - 'A');
//         }
//     }
//     return newStr;
// }
fn checkAllIfOnes(bits: [26]u8) bool {
    for (bits) |i| {
        if (i != 1) {
            return false;
        }
    }
    return true;
}
pub fn isPangram(str: []const u8) bool {
    const allocator = std.heap.page_allocator;
    const newStr = std.ascii.allocLowerString(allocator, str) catch undefined;

    defer allocator.free(newStr);
    var bits: [26]u8 = std.mem.zeroes([26]u8);
    for (newStr) |ch| {
        if (ch >= 'a' and ch <= 'z') {
            bits[ch - 'a'] = 1;
        }
    }
    return checkAllIfOnes(bits);
}
