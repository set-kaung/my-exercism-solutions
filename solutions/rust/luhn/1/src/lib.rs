/// Check a Luhn checksum.
///
///
fn two_times(v: (usize, char)) -> Option<char> {
    let (i, c) = v;
    if i % 2 != 0 && c != ' ' {
        match c.to_digit(10) {
            Some(mut d) => {
                d = d * 2;
                if d > 9 {
                    d = d - 9;
                    char::from_digit(d, 10)
                } else {
                    char::from_digit(d, 10)
                }
            }
            None => None,
        }
    } else {
        if c.is_digit(10) {
            Some(c)
        } else {
            None
        }
    }
}
pub fn is_valid(code: &str) -> bool {
    let code = code.replace(" ", "");
    if code.len() <= 1 {
        return false;
    }
    match code
        .chars()
        .rev()
        .enumerate()
        .map(two_times)
        .collect::<Option<Vec<char>>>()
    {
        Some(v) => {
            let v = v.iter().filter(|c| c.is_digit(10));
            let sum = v.map(|c| c.to_digit(10).unwrap()).sum::<u32>();
            sum % 10 == 0
        }
        None => false,
    }
}
