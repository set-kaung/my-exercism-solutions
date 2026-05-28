use std::collections::{HashSet, VecDeque};

#[derive(PartialEq, Eq, Debug)]
pub enum Bucket {
    One,
    Two,
}

/// A struct to hold your results in.
#[derive(PartialEq, Eq, Debug)]
pub struct BucketStats {
    /// The total number of "moves" it should take to reach the desired number of liters, including
    /// the first fill.
    pub moves: u8,
    /// Which bucket should end up with the desired number of liters? (Either "one" or "two")
    pub goal_bucket: Bucket,
    /// How many liters are left in the other bucket?
    pub other_bucket: u8,
}

#[derive(Debug, Clone, Copy)]
struct State {
    first: u8,
    second: u8,
    step: u8,
}

fn state_key(x: u8, y: u8) -> String {
    format!("({},{})", x, y)
}

fn next_states(s: State, c1: u8, c2: u8, intial_bucket: &Bucket) -> Vec<State> {
    let mut next_states = Vec::new();
    let first_pour = s.first.min(c2 - s.second);
    let second_pour = s.second.min(c1 - s.first);

    let next_pairs = vec![
        (c1, s.second),
        (s.first, c2),
        (0, s.second),
        (s.first, 0),
        (s.first - first_pour, s.second + first_pour),
        (s.first + second_pour, s.second - second_pour),
    ];
    for (x, y) in next_pairs {
        if (x == c1 && y == 0) || (x == 0 && y == c2) {
            match intial_bucket {
                Bucket::One => {
                    if x == 0 {
                        continue;
                    }
                }
                Bucket::Two => {
                    if y == 0 {
                        continue;
                    }
                }
            }
        }
        next_states.push(State {
            first: x,
            second: y,
            step: s.step + 1,
        });
    }
    next_states
}

fn find_answer(
    states: &mut VecDeque<State>,
    seen: &mut HashSet<String>,
    goal: u8,
    capacity_1: u8,
    capacity_2: u8,
    intial_bucket: &Bucket,
) -> Option<State> {
    loop {
        match states.pop_front() {
            Some(s) => {
                seen.insert(state_key(s.first, s.second));
                for next_state in next_states(s, capacity_1, capacity_2, intial_bucket) {
                    if next_state.first == goal || next_state.second == goal {
                        return Some(next_state);
                    }
                    if !seen.contains(&state_key(next_state.first, next_state.second)) {
                        states.push_back(next_state);
                    }
                }
            }
            None => break None,
        }
    }
}

/// Solve the bucket problem
pub fn solve(
    capacity_1: u8,
    capacity_2: u8,
    goal: u8,
    start_bucket: &Bucket,
) -> Option<BucketStats> {
    if goal > capacity_1.max(capacity_2) {
        return None;
    }

    let initial_state = match start_bucket {
        Bucket::One => State {
            first: capacity_1,
            second: 0,
            step: 1,
        },
        Bucket::Two => State {
            first: 0,
            second: capacity_2,
            step: 1,
        },
    };

    let mut seen: HashSet<String> = HashSet::new();
    seen.insert(state_key(initial_state.first, initial_state.second));

    let mut vec_deque: VecDeque<State> = VecDeque::new();
    vec_deque.push_back(initial_state);

    let answer: Option<State> = if initial_state.first == goal || initial_state.second == goal {
        Some(initial_state)
    } else {
        find_answer(
            &mut vec_deque,
            &mut seen,
            goal,
            capacity_1,
            capacity_2,
            start_bucket,
        )
    };

    match answer {
        Some(answer) => {
            let (result_bucket, other_val) = if answer.first == goal {
                (Bucket::One, answer.second)
            } else {
                (Bucket::Two, answer.first)
            };

            Some(BucketStats {
                moves: answer.step,
                goal_bucket: result_bucket,
                other_bucket: other_val,
            })
        }
        None => None,
    }
}
