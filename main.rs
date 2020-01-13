use std::{
  fs::File,
  io::{prelude::*, BufReader},
  path::Path,
};

fn lines_from_file(filename: impl AsRef<Path>) -> Vec<String> {
  let file = File::open(filename).expect("no such file");
  let buf = BufReader::new(file);
  buf.lines()
    .map(|l| l.expect("Could not parse line"))
    .collect()
}

fn get_name_index(name: &String, array: &Vec<String>) -> Option<usize> {
  array.iter().position(|x| x == name)
}

fn main() {
  let mut data: Vec<String> = Vec::new();

  let search_query = String::from("servmill.com");

  let lines = lines_from_file("./url-list.txt");
  for line in lines {
    data.push(line);
  }

  let _index = get_name_index(&search_query, &data);

  println!("{:?}", _index);  
}