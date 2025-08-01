const std = @import("std");

pub const QueenError = error{
    InitializationFailure,
};

pub fn abs(n: i8) i8 {
    if (n < 0) {
        return -n;
    }
    return n;
}

pub const Queen = struct {
    row: i8,
    col: i8,

    pub fn init(row: i8, col: i8) QueenError!Queen {
        if (row < 0 or col < 0 or row > 7 or col > 7) {
            return QueenError.InitializationFailure;
        }
        return Queen{ .row = row, .col = col };
    }

    pub fn canAttack(self: Queen, other: Queen) QueenError!bool {
        if (self.col == other.col or self.row == other.row) {
            return true;
        }
        const div = std.math.divCeil(i8, ((self.row + 1) - (other.row + 1)), ((self.col + 1) - (other.col + 1))) catch {
            return QueenError.InitializationFailure;
        };
        const slope = abs(div);
        std.debug.print("slope = {d}\n", .{slope});
        if (slope == 1) {
            return true;
        }

        return false;
    }
};
