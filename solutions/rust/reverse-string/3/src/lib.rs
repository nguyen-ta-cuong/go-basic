pub fn reverse(input: &str) -> String {
    let mut chars: Vec<char> = Vec::new();

    for c in input.chars() {
        chars.push(c);
    }

    if chars.is_empty() {
        return String::new();
    }
    
    let mut start: usize = 0;
    let mut end: usize = chars.len() - 1;
    while start < end {
        chars.swap(start,end);
        start += 1;
        end -= 1;
    }
    let mut result = String::new();
    for c in chars {
        result.push(c);
    }
    result
}
