mod mult;
use crate::mult::get_mult;

use std::io::{BufRead, stdin};

fn main() {
    let lines = stdin().lock().lines();

    let mut number_count = 0;
    let mut mult = 1i64;
    for line in lines {
        let line_str = line.expect("Failed to read line");
        if line_str.is_empty() {
            continue;
        }

        let parts = line_str.split(' ').filter(|line| !line.is_empty());
        for part in parts {
            mult = get_mult(mult, part.parse::<i64>().expect("Failed to parse i64"));
            number_count += 1;

            if number_count >= 2 {
                break;
            }
        }
        if number_count >= 2 {
            break;
        }
    }

    println!("Mult: {}", mult);
}
