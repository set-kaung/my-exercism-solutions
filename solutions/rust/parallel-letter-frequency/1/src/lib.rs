use std::{
    collections::HashMap,
    sync::mpsc::{self},
    thread,
};

pub fn frequency(input: &[&str], worker_count: usize) -> HashMap<char, usize> {
    let input_text = input;
    let (sender, receiver) = mpsc::sync_channel::<HashMap<char, usize>>(worker_count);

    let size = match input_text.len() / worker_count {
        0 => 1,
        other => other,
    };

    let mut partitions: Vec<Vec<String>> = Vec::new();
    let mut start = 0;

    while start < input_text.len() {
        let end = (start + size).min(input.len());
        let mut v: Vec<String> = Vec::new();
        for &s in &input[start..end] {
            v.push(s.to_string());
        }
        partitions.push(v);
        start = end;
    }

    for part in partitions {
        let sender_clone = sender.clone();
        thread::spawn(move || {
            let mut map: HashMap<char, usize> = HashMap::new();
            for c in part {
                for r in c.chars() {
                    if r.is_alphabetic() {
                        let r = r.to_ascii_lowercase();
                        *map.entry(r).or_insert(0) += 1;
                    }
                }
            }
            sender_clone.send(map).unwrap();
        });
    }
    drop(sender);
    let mut final_map = HashMap::new();
    while let Ok(partial_map) = receiver.recv() {
        for (key, value) in partial_map {
            *final_map.entry(key).or_insert(0) += value;
        }
    }

    final_map
}
