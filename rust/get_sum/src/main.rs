use std::io::{stdin, BufRead};

fn main() {
    let mut lines = stdin().lock().lines();

    let mut number_count = 0;
    let mut sum = 0i64;
    while let Some(line) = lines.next() {
        let line_str = line.expect("Failed to read line");
        if line_str.is_empty() {
            continue;
        }

        let parts = line_str.split(" ");
        for part in parts {
            sum += part.parse::<i64>().expect("Failed to parse i64");
            number_count += 1;

            if number_count >= 2 {
                break;
            }
        }
        if number_count >= 2 {
            break;
        }
    }

    println!("Sum: {}", sum);
}
