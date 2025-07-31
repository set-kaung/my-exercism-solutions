// Please implement the `ComputationError.IllegalArgument` error.

pub const ComputationError = error{
    IllegalArgument,
};

pub fn steps(number: usize) anyerror!usize {
    var number2 = number;
    if (number2 <= 0) {
        return ComputationError.IllegalArgument;
    }
    var count: usize = 0;
    while (number2 != 1) {
        if (number2 % 2 == 0) {
            number2 = number2 / 2;
        } else {
            number2 = (number2 * 3) + 1;
        }
        count += 1;
    }
    return count;
}
