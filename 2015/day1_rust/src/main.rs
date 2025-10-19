use std::fs;

fn main() {
    let input = fs::read_to_string("input.txt").expect("Should have been able to read file");
    let chars: Vec<char> = input.chars().collect();
    let mut floor = 0;
    let mut i = 0;
    for char in chars {
        if char == '(' {
            floor += 1;
        }
        if char == ')' {
            floor -= 1;
        }
        i += 1;
        if floor == -1 {
            println!("floor -1 at {}", i);
        }
    }
    println!("final floor: {}", floor);
}
