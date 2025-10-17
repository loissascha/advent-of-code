use std::fs;

fn main() {
    let file_path = "input_test2.txt";
    let parts = get_lines(file_path);
    let mut sum = 0;
    for part in parts {
        let numbers = get_numbers(&part);
        sum += numbers;
        println!("{} {}", part, numbers);
    }
    println!("Sum: {}", sum);
}

fn get_numbers(line: &str) -> i32 {
    let chars = line.chars();
    let mut result = String::from("");
    for char in chars {
        if char >= '0' && char <= '9' {
            result.push(char);
        }
    }
    let length = result.chars().count();
    if length == 0 {
        return 0;
    }
    if length == 1 {
        result.push_str(&result.clone());
    }
    if length > 2 {
        let chars: Vec<char> = result.chars().collect();
        let first = chars[0];
        let last = chars[chars.iter().count() - 1];
        result = String::from("");
        result.push(first);
        result.push(last);
    }
    result.parse().unwrap()
}

fn get_lines(file_path: &str) -> Vec<String> {
    let contents = fs::read_to_string(file_path).expect("Should have been able to read the file");

    contents.lines().map(|s| s.to_string()).collect()
}
