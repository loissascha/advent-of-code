use std::fs;

fn main() {
    let input = fs::read_to_string("input.txt").expect("expected");
    let mut dial_pos = 50;
    let mut zeros = 0;

    let split = input.split("\n");
    for line in split {
        if line == "" {
            continue;
        }
        println!("line: {line}");
        if line.starts_with("L") {
            let t = line.strip_prefix("L").expect("ex");
            let tt: i32 = t.parse::<i32>().expect("exp");
            for _ in 0..tt {
                dial_pos = click_left(dial_pos);
                if dial_pos == 0 {
                    zeros += 1;
                }
            }
        }
        if line.starts_with("R") {
            let t = line.strip_prefix("R").expect("ex");
            let tt: i32 = t.parse::<i32>().expect("exp");
            for _ in 0..tt {
                dial_pos = click_right(dial_pos);
                if dial_pos == 0 {
                    zeros += 1;
                }
            }
        }
    }

    println!("zeros: {zeros}");
}

fn click_left(mut dial_pos: i32) -> i32 {
    dial_pos = dial_pos - 1;
    if dial_pos < 0 {
        dial_pos = 99;
    }
    dial_pos
}

fn click_right(mut dial_pos: i32) -> i32 {
    dial_pos = dial_pos + 1;
    if dial_pos > 99 {
        dial_pos = 0;
    }
    dial_pos
}
