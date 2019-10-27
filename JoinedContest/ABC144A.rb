a, b = gets.chomp.split(" ").map(&:to_i)

if a >= 10 || b >= 10
  puts -1
else
  puts a * b
end