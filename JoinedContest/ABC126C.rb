n, k = gets.chomp.split(" ").map(&:to_i)

sigma = 0.0
(1..n).each do |i|
    ruijo = 1
    p = 1
    while k > i * ruijo
        ruijo *= 2
        p *= 0.5
    end
    sigma += p
    # puts "#{i} : #{k} < #{i}*#{ruijo} #{sigma}"
end
sigma = sigma.to_f / n.to_f

puts sigma