#[allow(unused)]
use std::fs::read_to_string;
fn main() {
    let red = 12;
    let green = 13;
    let blue = 14;

    let mut gameID = 0;
    let mut input = read_to_string("input.txt").unwrap();
    input = input.replace(":", "");
    input = input.replace(",", "");
    input = input.replace(";", "");

    let mut sum = 0;
    for line in input.lines() {
        let words:Vec<&str> = line.split(" ").collect();

        gameID = words[1].parse().unwrap();

        let mut index = 3;

        let mut isValid = true; 

        while index < words.len() {
            match words[index] {
                "red"=>if words[index - 1].parse::<i32>().unwrap() > red {isValid = false; break},
                "blue"=>if words[index - 1].parse::<i32>().unwrap() > blue {isValid = false; break},
                "green"=>if words[index - 1].parse::<i32>().unwrap() > green {isValid = false; break},
                __=>println!("{}", words[index])
            }
            index += 2;
        }
        if isValid {
            sum += gameID;
        }
    }

    println!("\n{}\n", sum)


}