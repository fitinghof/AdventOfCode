#[allow(unused)]
use std::fs::read_to_string;
fn main() {
    let mut input = read_to_string("input.txt").unwrap();
    input = input.replace(":", "");
    input = input.replace(",", "");
    input = input.replace(";", "");

    let mut sum = 0;
    for line in input.lines() {
        let words:Vec<&str> = line.split(" ").collect();

        let mut index = 3;

        let mut red = 0;
        let mut green = 0;
        let mut blue = 0;
    
        while index < words.len() {
            match words[index] {
                "red"=>if words[index - 1].parse::<i32>().unwrap() > red {red = words[index - 1].parse::<i32>().unwrap()},
                "blue"=>if words[index - 1].parse::<i32>().unwrap() > blue {blue = words[index - 1].parse::<i32>().unwrap()},
                "green"=>if words[index - 1].parse::<i32>().unwrap() > green {green = words[index - 1].parse::<i32>().unwrap()},
                __=>println!("{}", words[index])
            }
            index += 2;
        }
        sum += red*blue*green;
    }
    println!("\n{}\n", sum)
}