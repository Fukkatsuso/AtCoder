n = gets.chomp.to_i
a = []
(1..n).each do
    a.push(gets.chomp.to_i - 1) # 次のボタンのインデックス
end

btn = 0

(1..n).each do |cnt|
    if btn == 1
        puts "#{cnt-1}"
        break
    end
    if cnt == n
        puts -1
        break
    end
    btn = a[btn]
end