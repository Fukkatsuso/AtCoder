s = gets.chomp.split("")
k = s.length

win = 0
lose = 0
s.each do |i|
    win += 1 if i == 'o'
    lose += 1 if i == 'x'
end

if k < 8
    puts "YES"
elsif win + (15 - k) > 7
    puts "YES"
else
    puts "NO"
end