w, h, x, y = gets.chomp.split(" ").map(&:to_i)

s = w.to_f * h.to_f / 2.0

if (w % 2 == 0) && (h % 2 == 0) && (x == w / 2) && (y == h / 2)
    puts "#{s} 1"
else
    puts "#{s} 0"
end