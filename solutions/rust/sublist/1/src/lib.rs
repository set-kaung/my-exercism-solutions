#[derive(Debug, PartialEq, Eq)]
pub enum Comparison {
    Equal,
    Sublist,
    Superlist,
    Unequal,
}

fn isSubList<T: PartialEq>(longer: &[T], shorter: &[T]) -> bool {
    let mut isEq = false;
    for i in 0..longer.len() {
        if longer[i] == shorter[0] && i + shorter.len() <= longer.len() {
            isEq = true;
            for j in 0..shorter.len() {
                if longer[i + j] != shorter[j] {
                    isEq = false
                }
            }
            if isEq {
                break;
            }
        }
    }
    return isEq;
}

pub fn sublist<T: PartialEq>(first: &[T], second: &[T]) -> Comparison {
    if first.len() == 0 && second.len() == 0 {
        return Comparison::Equal;
    }
    if first.len() == 0 {
        return Comparison::Sublist;
    }

    if second.len() == 0 {
        return Comparison::Superlist;
    }
    if first.len() > second.len() {
        if isSubList(first, second) {
            return Comparison::Superlist;
        }
    } else if first.len() < second.len() {
        if isSubList(second, first) {
            return Comparison::Sublist;
        }
    } else {
        if isSubList(first, second) {
            return Comparison::Equal;
        }
    }
    return Comparison::Unequal;
}
