c500 = gets.chomp.to_i
c100 = gets.chomp.to_i
c50 = gets.chomp.to_i
x = gets.chomp.to_i
n = 0

(0..c500).each do |a|
    (0..c100).each do |b|
        (0..c50).each do |c|
            n += 1 if a*500 + b*100 + c*50 == x
        end
    end
end

puts n