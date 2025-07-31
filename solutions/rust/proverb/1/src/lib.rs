use itertools::Itertools;
pub fn build_proverb(list: &[&str]) -> String {
    if list.is_empty() {
        return String::new();
    }
    list.into_iter()
        .tuple_windows()
        .map(|(a, b)| format!("For want of a {} the {} was lost.", a, b))
        .chain(
            list.get(0)
                .map(|first| format!("And all for the want of a {}.", first))
                .into_iter(),
        )
        .collect::<Vec<String>>()
        .join("\n")
}
