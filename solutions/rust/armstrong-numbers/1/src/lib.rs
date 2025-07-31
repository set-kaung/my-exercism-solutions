pub fn is_armstrong_number(num: u32) -> bool {
    let str = num.to_string();
    let len = str.len() as u32;
    let new_num = str
        .chars()
        .map(|x| (x.to_digit(10).unwrap() as u128).pow(len as u32))
        .fold(0, |accum, ele| accum + ele);
    num as u128 == new_num
}
