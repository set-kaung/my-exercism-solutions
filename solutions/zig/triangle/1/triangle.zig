pub const TriangleError = error{Invalid};

pub const Triangle = struct {
    a: f64,
    b: f64,
    c: f64,

    pub fn init(a: f64, b: f64, c: f64) TriangleError!Triangle {
        if (a == 0 or b == 0 or c == 0) {
            return TriangleError.Invalid;
        }
        if (a + b < c or a + c < b or b + c < a) {
            return TriangleError.Invalid;
        }
        const tri = Triangle{ .a = a, .b = b, .c = c };
        return tri;
    }

    pub fn isEquilateral(self: Triangle) bool {
        return self.a == self.b and self.b == self.c;
    }

    pub fn isIsosceles(self: Triangle) bool {
        return self.a == self.b or self.b == self.c or self.a == self.c;
    }

    pub fn isScalene(self: Triangle) bool {
        return !self.isIsosceles();
    }
};
