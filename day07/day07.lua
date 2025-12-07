file = io.open("input.txt","r")
io.input(file)

lines = {}

for line in io.lines() do
    t = {}
    for ch in line:gmatch(".") do
        t[#t+1] = ch
    end
    table.insert(lines, t)
end

part1 = 0

timelines = {}

for i = 2, #lines, 1 do
    curr = lines[i]
    prev = lines[i-1]

    timelineLine = {}
    for j = 1, #curr, 1 do
        timelineLine[j] = 0
    end
    --split laser
    for j = 1, #curr, 1 do
        if curr[j] == '.' and prev[j] == 'S' then
            curr[j] = '|'
            timelineLine[j] = 1
        elseif curr[j] == '^' and prev[j] == '|' then
            curr[j-1] = '|'
            curr[j+1] = '|'
            part1 = part1 + 1
        elseif curr[j] == '.' and prev[j] == '|' then
            curr[j] = '|'
        end
    end
    --calc possible timelines
    for j = 1, #curr, 1 do
        if curr[j] == '|' then
            if j > 1 and curr[j-1] == '^' then
                timelineLine[j] = timelineLine[j] + timelines[i-1][j-1]
            end
            if prev[j] == '|' then
                timelineLine[j] = timelineLine[j] + timelines[i-1][j]
            end
            if j < #curr and curr[j+1] == '^' then
                timelineLine[j] = timelineLine[j] + timelines[i-1][j+1]
            end
        end
    end

    timelines[i] = timelineLine 
    lines[i] = curr
end

print(part1)

part2 = 0
for i = 1, #(timelines[#timelines]), 1 do
    part2 = part2 + math.tointeger(timelines[#timelines][i])
end
print(part2)