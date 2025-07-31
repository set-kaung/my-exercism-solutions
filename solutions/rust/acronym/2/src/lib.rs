pub fn abbreviate(phrase: &str) -> String {
    let mut result = String::new();
    let mut is_after_separator = false;
    let mut is_scanning = true;
    for i in phrase.chars(){
        if i == '-' || i == ' ' {
            is_after_separator = true;
        }else if is_after_separator && i.is_alphabetic(){
            result.push(i.to_ascii_uppercase());
            is_after_separator = false;
        }else if !is_after_separator && i.is_lowercase(){
            is_scanning = true;
        }else if is_scanning && i.is_uppercase() {
            result.push(i.to_ascii_uppercase());
            is_scanning = false
        }
    }

    result

}
