use std::fs;

fn main() {
    let file_path = "input_test.txt";
    let parts = get_lines(file_path);
    for part in parts {
        println!("{}", part);
    }
}

fn get_lines(file_path: &str) -> Vec<String> {
    let contents = fs::read_to_string(file_path).expect("Should have been able to read the file");

    contents.lines().map(|s| s.to_string()).collect()
}
