# TLE
n = gets.chomp.to_i
a = gets.chomp.split(" ").map(&:to_i)

max0 = a[0] <= a[n-1] ? 2*a[0] : 2*a[n-1]

mt = [0]*n
(0..max0).each do |mx|
    mt[0] = mx
    rp = true
    (1..n-1).each do |i|
        mt[i] = a[i-1] - mt[i-1]
        if mt[i] < 0
            rp = false
            break
        end
    end
    if mt[0] + mt[n-1] == a[n-1] && rp
        (0..n-1).each do |i|
            print "#{2*mt[i]} "
        end
        puts ""
        break
    end
end