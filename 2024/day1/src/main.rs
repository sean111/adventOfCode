use std::fs::File;
use std::io::{stdout, BufReader};
use std::io::prelude::*;
use regex::Regex;

fn main() {
    let p1res = part1("data/data.txt").unwrap();
    writeln!(stdout(), "Part1: {}", p1res).unwrap();
}

pub fn part1(file: &str) -> std::io::Result<i32> {
    let fp = File::open(file)?;
    let mut reader = BufReader::new(fp);
    let mut contents = String::new();
    let mut list1:Vec<i32> = Vec::new();
    let mut list2:Vec<i32> = Vec::new();
    let mut total: i32 = 0;
    let regex = Regex::new(r"(\d+)\s\s\s(\d+)").unwrap();
    reader.read_to_string(&mut contents)?;
    writeln!(stdout(), "{}", contents)?;
    let values: Vec<(i32, i32)> = regex.captures_iter(contents.as_str()).map(|numbers| {
        let l1 = numbers.get(1).unwrap().as_str().parse::<i32>().unwrap();
        let l2 = numbers.get(2).unwrap().as_str().parse::<i32>().unwrap();
        (l1, l2)
    }).collect();
    writeln!(stdout(), "values {:?}", values)?;

    values.into_iter().for_each(|(l1, l2)| {
        list1.push(l1);
        list2.push(l2);
    });
    list1.sort();
    list2.sort();

    let count = list1.len();

    for x in 0..count {
        total += (list1[x] - list2[x]).abs();
    }

    writeln!(stdout(), "list1 {:?}", list1)?;
    writeln!(stdout(), "list2 {:?}", list2)?;

    Ok(total)
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_part1() {
        let result = part1("data/test.txt").unwrap();
        assert_eq!(result, 11);
    }
}
