div = ["maerd", "remaerd", "esare", "resare"]
s = gets.chomp.reverse

flag = true
while flag && s!=nil do
    flag = false
    div.each do |d|
        if s[0..d.length-1] == d
            s.slice!(d)
            flag = true
            break
        end
    end
end

if s == nil || s.length == 0
    puts "YES"
else
    puts "NO"
end