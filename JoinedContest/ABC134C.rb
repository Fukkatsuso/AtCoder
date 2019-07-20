n = gets.chomp.to_i
a = []
n.times do
    a << gets.chomp.to_i
end
b = a.sort.reverse  # 降順
max = b[0]

(0..n-1).each do |i|
    if a[i] != max
        puts max
    else
        puts b[1]
    end
end