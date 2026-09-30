pub fn reverse(input: &str) -> String {
    let mut chars: Vec<char> = Vec::new();

    for c in input.chars() {
        chars.push(c);
    }

    if chars.len() == 0 {
        return String::new();
    }
    
    let mut start: usize = 0;
    let mut end: usize = chars.len() - 1;
    while start < end {
        let temp = chars[start];
        chars[start] = chars[end];
        chars[end] = temp;
        start += 1;
        end -= 1;
    }
    let mut result = String::new();
    for c in chars {
        result.push(c);
    }
    result
}
