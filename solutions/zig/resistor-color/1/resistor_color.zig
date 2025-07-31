pub const ColorBand = enum { black, brown, red, orange, yellow, green, blue, violet, grey, white };

pub fn colorCode(color: ColorBand) usize {
    const cs = colors();
    for (cs, 0..) |c, i| {
        if (c == color) {
            return i;
        }
    }
    return 0;
}
pub fn colors() []const ColorBand {
    return &[_]ColorBand{
        .black, .brown, .red,    .orange, .yellow,
        .green, .blue,  .violet, .grey,   .white,
    };
}
