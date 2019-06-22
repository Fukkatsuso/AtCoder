s = gets.chomp.split("").map(&:to_i)

good = true
(0..2).each do |i|
    good = false if s[i] == s[i+1]
end
puts 'Bad' if !good
puts 'Good' if good