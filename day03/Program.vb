Imports System
Imports System.Linq

Module Program
    Sub Main(args As String())
        Dim file As List(Of String) = IO.File.ReadAllLines("input.txt").ToList
        Console.WriteLine(Calculate(file, 2)) 'part1
        Console.WriteLine(Calculate(file, 12)) 'part2
    End Sub

    Function Calculate(file As List(Of String), len As Integer) As Long
        Dim count As Long = 0
        For Each line In file
            Dim bestStr = ""
            Dim tempStr = line
            For i As Integer = len To 1 Step -1
                Dim largest = tempStr.Remove(tempStr.Length + 1 - i).Max
                Dim index = tempStr.IndexOf(largest)
                bestStr += largest
                tempStr = tempStr.Substring(index + 1)
            Next
            count += Int64.Parse(bestStr)
        Next
        Return count
    End Function
End Module
