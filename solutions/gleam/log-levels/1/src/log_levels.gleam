import gleam/result
import gleam/string

pub fn message(log_line: String) -> String {
  string.split_once(string.trim(log_line), on: " ")
  |> result.map(fn(x) { string.trim(x.1) })
  |> result.unwrap("")
}

pub fn log_level(log_line: String) -> String {
  string.split_once(log_line, on: ":")
  |> result.map(fn(x) { string.slice(x.0, 1, length: string.length(x.0) - 2) })
  |> result.unwrap("")
  |> string.lowercase
}

pub fn reformat(log_line: String) -> String {
  string.concat([message(log_line), " ", "(", log_level(log_line), ")"])
}
