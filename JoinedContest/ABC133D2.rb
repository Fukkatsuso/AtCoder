# AC
n = gets.chomp.to_i
a = gets.chomp.split(" ").map(&:to_i)

mt = [0]*n

(0..n-1).each do |i|
    mt[0] += (i%2==1 ? -a[i] : a[i])
end

(0..n-1).each do |i|
    mt[i+1] = 2 * a[i] - mt[i]
end

(0..n-1).each do |i|
    print "#{mt[i]} "
end
puts ""