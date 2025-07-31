pub const ColorBand = enum(usize) {
    black = 0,
    brown = 1,
    red = 2,
    orange = 3,
    yellow = 4,
    green = 5,
    blue = 6,
    violet = 7,
    grey = 8,
    white = 9,
};

const std = @import("std");

pub fn colorCode(colors: [2]ColorBand) usize {
    const first = colors[0];
    const second = colors[1];

    return (@intFromEnum(first) * 10) + @intFromEnum(second);
}
