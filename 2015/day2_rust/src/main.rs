use std::fs;

fn main() {
    let input = fs::read_to_string("input.txt").expect("Could not read file");
    let split: Vec<String> = input.split("\n").map(|s| s.to_string()).collect();

    for sp in split {
        if sp == String::from("") {
            continue;
        }
        process_line(&sp);
    }
}

fn process_line(input: &String) {
    println!("Process line: {}", input)
}
