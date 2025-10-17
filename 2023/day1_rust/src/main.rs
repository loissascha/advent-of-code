use std::fs;

fn main() {
    let file_path = "input.txt";
    let parts = get_lines(file_path);
    let mut sum = 0;
    for part in parts {
        let numbers = get_numbers(&part);
        sum += numbers;
        println!("{} {}", part, numbers);
    }
    println!("Sum: {}", sum);
}

fn get_str_number(input: &str) -> i32 {
    // println!("get str number input {input}");
    if input.contains("one") {
        return 1;
    }
    if input.contains("two") {
        return 2;
    }
    if input.contains("three") {
        return 3;
    }
    if input.contains("four") {
        return 4;
    }
    if input.contains("five") {
        return 5;
    }
    if input.contains("six") {
        return 6;
    }
    if input.contains("seven") {
        return 7;
    }
    if input.contains("eight") {
        return 8;
    }
    if input.contains("nine") {
        return 9;
    }
    return 0;
}

fn get_numbers(line: &str) -> i32 {
    println!("Get numbers for {line}");
    let mut result = String::from("");

    // find first number
    let mut i = 0;
    for char in line.chars() {
        // if the text up until this point might form a 'text' number (one, two, three, ...)
        let sub: String = line.chars().take(i).collect();
        let strnum = get_str_number(&sub);
        if strnum > 0 {
            result.push_str(&strnum.to_string());
            break;
        }

        // if it is a number
        if char >= '0' && char <= '9' {
            println!("char {char} is number");
            result.push(char);
            break;
        }
        i += 1;
    }

    // find last number
    let mut i = 0;
    let length = line.chars().count();
    for char in line.chars().rev() {
        let idx = length - i - 1;
        let sub: String = line.chars().skip(idx).collect();
        println!("find last num sub: {sub} current char: {char}");
        let strnum = get_str_number(&sub);
        if strnum > 0 {
            result.push_str(&strnum.to_string());
            break;
        }

        // if it is a number
        if char >= '0' && char <= '9' {
            println!("char {char} is number");
            result.push(char);
            break;
        }

        i += 1;
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
