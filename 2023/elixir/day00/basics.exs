input = File.read!("input.txt")
lines = String.split(input, "\n")

Enum.map(lines, fn l ->
  if l != "" do
    numbers = String.split(l, " ")

    sum =
      List.foldl(numbers, 0, fn n, acc ->
        case Integer.parse(n) do
          {i, _} -> i + acc
          _ -> 0
        end
      end)

    IO.puts(sum)
  end
end)
