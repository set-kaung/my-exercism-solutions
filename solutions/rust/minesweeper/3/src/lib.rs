pub fn annotate(minefield: &[&str]) -> Vec<String> {
    let moves = [
        (1, 0),
        (1, 1),
        (0, 1),
        (-1, 1),
        (1, -1),
        (-1, 0),
        (0, -1),
        (-1, -1),
    ];
    let mut result: Vec<String> = Vec::new();
    for i in 0..minefield.len() {
        let mut curr_row = String::new();
        for j in 0..minefield[i].len() {
            if minefield[i].as_bytes()[j] == b' ' {
                let x = check(i, j, &moves, minefield);
                if x > 0 {
                    curr_row.push_str(&x.to_string());
                    continue;
                }
            }
            curr_row.push(minefield[i].as_bytes()[j] as char);
        }
        result.push(curr_row);
    }
    result
}

pub fn is_valid(x: i32, y: i32, row_len: i32, col_len: i32) -> bool {
    x >= 0 && x < row_len && y >= 0 && y < col_len
}

pub fn check(i: usize, j: usize, moves: &[(i32, i32)], minefield: &[&str]) -> usize {
    let mut count = 0;
    let row_len = minefield.len() as i32;
    let col_len = minefield[i].len() as i32;
    for (x, y) in moves {
        let new_x = i as i32 + y;
        let new_y = j as i32 + x;
        if is_valid(new_x, new_y, row_len, col_len)
            && minefield[new_x as usize].as_bytes()[new_y as usize] == b'*'
        {
            count += 1
        }
    }
    count
}
