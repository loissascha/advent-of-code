use std::fs;

fn main() {
    let input = fs::read_to_string("input.txt").expect("Should have been able to read file");
    let chars: Vec<char> = input.chars().collect();
    let mut floor = 0;
    for char in chars {
        if char == '(' {
            floor += 1;
        }
        if char == ')' {
            floor -= 1;
        }
    }
    println!("final floor: {}", floor);
}
