n = gets
a = gets.chomp.split(" ").map(&:to_i)
i = 0
flag = 1
while flag == 1 do
    a.each do |b|
        if b % 2 == 1 && flag == 1
            puts i
            flag = 0
        end
    end
    a.map!{|b| b /= 2}
    i += 1
end