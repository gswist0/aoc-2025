#include <iostream>
#include <fstream>
#include <vector>

using namespace std;

bool check(vector<string> lines, int x, int y){
    int adjacents = 0;
    int startRangeX = (x < 1) ? 0 : x-1;
    int endRangeX  = (x >= lines[y].length() - 1) ? lines[y].length() - 1 : x + 1;
    int startRangeY = (y < 1) ? 0 : y - 1;
    int endRangeY = (y >= lines.size() - 1) ? lines[y].length() - 1 : y + 1;
    for(int i = startRangeX; i <= endRangeX; i++){
        for(int j = startRangeY; j <= endRangeY; j++){
            if(lines[i][j] == '@'){
                adjacents++;
            }
            if (adjacents > 4){
                return false;
            }
        }
    }
    return true;
}

int main(){
    vector<string> lines;
    ifstream file("input.txt");
    string line;
    while(getline(file,line)){
        lines.push_back(line);
    }
    int part1 = 0;
    for (int i = 0; i < lines.size(); i++){
        for(int j = 0; j<lines[i].length(); j++){
            if(lines[i][j] == '@' && check(lines,i,j)){
                part1++;
            }
        }
    }
    cout<<part1<<endl;
    int part2 = 0;
    bool allNonAccessible = false;
    while(!allNonAccessible){
        vector<string> newLines = lines;
        int temp = 0;
        for (int i = 0; i < lines.size(); i++){
            for(int j = 0; j<lines[i].length(); j++){
                if(lines[i][j] == '@' && check(lines,i,j)){
                    temp++;
                    newLines[i][j] = '.';
                }
            }
        }
        part2 += temp;
        lines = newLines;
        if (temp == 0){
            allNonAccessible = true;
        }
    }
    cout<<part2<<endl;
}

