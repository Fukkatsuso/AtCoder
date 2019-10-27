# WA, TLE

n = gets.chomp.to_i
l = gets.chomp.split(" ").map(&:to_i).sort

ans = 0

(2..n-1).each do |i|
  (1..i-1).each do |j|  # l[i] >= l[i]
    l_min = l[i] - l[j] + 1
    # l[k] >= l_min となるkを探索
    left = 0
    right = j - 1
    k = right / 2
    while l[k] != l_min && (right - left) > 1 do
      if l[k] < l_min
        left = k + 1 
      elsif l[k] > l_min
        right = k
      end
      k = (right + left) / 2
    end

    ans += l[k..j-1].uniq.size if l[k] >= l_min
  end
end

puts ans