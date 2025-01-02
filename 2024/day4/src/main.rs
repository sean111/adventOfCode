use std::collections::HashMap;
use std::fs::File;
use std::io::{stdout, BufReader, Read, Result};
use std::io::prelude::*;

fn main() {
    writeln!(stdout(), "Part 1: {}", part1("data/data.txt").unwrap()).unwrap();
}

pub fn part1(file: &str) -> Result<i32> {
    let mut data: HashMap<(i32, i32), char> = HashMap::new();
    let fp = File::open(file)?;
    let mut reader = BufReader::new(fp);
    let mut contents = String::new();
    reader.read_to_string(&mut contents)?;
    let mut line_count: i32 = 0;
    let mut char_count: i32 = 0;
    for line in contents.lines() {
        char_count = 0;
        for char in line.chars() {
            data.insert((line_count, char_count), char);
            char_count += 1;
        }
        line_count += 1;
    }

    for line in 0 .. line_count {
        for char in 0 .. char_count {
            match char {
            }
        }
    }

    Ok(0)
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn part1_test() {
        assert_eq!(part1("data/test.txt").unwrap(), 18);
    }
}
