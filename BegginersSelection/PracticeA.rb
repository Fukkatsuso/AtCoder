a = gets.chomp.to_i
b = gets.chomp.split(" ").map(&:to_i)
s = gets.chomp
puts "#{a + b[0] + b[1]}" + " " + s