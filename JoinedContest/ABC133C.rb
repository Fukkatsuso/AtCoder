l, r = gets.chomp.split(" ").map(&:to_i)

if r - l >= 2019
    puts 0
elsif r%2019 <= l%2019
    puts 0
else
    min = 2019
    (l..r-1).each do |i|
        (i+1..r).each do |j|
            md = ((i%2019) * (j%2019)) % 2019
            min = min <= md ? min : md
        end
    end
    puts min
end