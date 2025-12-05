const fs = require('node:fs');
const lineReader = require('line-reader');
const { kMaxLength } = require('node:buffer');


const fresh_ids = [];
const ids = [];

function mergeRanges(range1, range2){
    const x = range1[0]
    const y = range1[1]
    const a = range2[0]
    const b = range2[1]
    switch(true){
        case x > b || a > y :
            return [];
        case x <= b && x >= a :
            return  [Math.min(x,a), Math.max(y,b)];
        case a <= y && a >= x :
            return [Math.min(a,x), Math.max(y,b)];
        case y <= b && y >= a :
            return [Math.min(a,x), Math.max(y,b)];
        case b <= y && b >= x :
            return [Math.min(a,x), Math.max(b,y)];
        default:
            return [];
    }
}

lineReader.eachLine('input.txt', (line,last) => {
    switch(true){
        case line.includes("-"):
            const split = line.split("-");
            const start = parseInt(split[0]);
            const end = parseInt(split[1]);
            fresh_ids.push([start,end]);
            break;
        case line.trim() === "":
            break;
        default:
            ids.push(parseInt(line));
            break;
    }
    if(last){
        let part1 = 0;
        ids.forEach(id => {
            fresh_ids.some(range => {
                if(range[0] <= id && range[1] >= id){
                    part1++;
                    return true;
                }
            })
        });
        console.log(part1)
        let part2 = 0;
        let mergedIds = fresh_ids;
        let allDone = false
        while(!allDone){
            let merged = 0;
            for (let i = 0; i < mergedIds.length; i++){
                let br = false
                for (let j = i + 1; j < mergedIds.length; j++){
                    const result = mergeRanges(mergedIds[i],mergedIds[j]);
                    if (result.length != 0){
                        mergedIds.splice(j,1)
                        mergedIds.splice(i,1)
                        mergedIds.push(result)
                        br = true
                        merged++;
                        break;
                    }
                }
                if (br)
                    break;
            }
            if (merged == 0){
                allDone = true;
            }
        }
        mergedIds.forEach(range => {
            part2+=range[1] - range[0] + 1
        })
        console.log(part2)
    }
});


