use std::collections::HashSet;

fn same_word(target: String, source: &String) -> bool {
    let mut source_clone = source.clone().chars().collect::<Vec<char>>();
    source_clone.sort();
    let mut target_clone = target.clone().chars().collect::<Vec<char>>();
    target_clone.sort();

    source_clone.iter().collect::<String>() == target_clone.iter().collect::<String>()
}

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &[&'a str]) -> HashSet<&'a str> {
    let mut set: HashSet<&str> = HashSet::new();
    let word = word.to_lowercase();
    for i in possible_anagrams {
        let target = (*i).to_lowercase();
        if word == target || (*i).len() != word.len() {
            continue;
        }
        if same_word(target, &word) {
            set.insert(i);
        }
    }
    set
}
