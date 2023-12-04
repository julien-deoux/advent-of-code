let ( >> ) f g x = g (f x)

let rec lines_of_file f =
  match input_line f with
  | exception End_of_file -> []
  | line -> line :: lines_of_file f

let line_sum =
  String.split_on_char ' ' >> List.map int_of_string >> List.fold_left ( + ) 0

let handle_file =
  lines_of_file >> List.map line_sum
  >> List.iter (string_of_int >> print_endline)

let handle_path path =
  let file = open_in path in
  handle_file file;
  close_in file

let () = handle_path "input.txt"
