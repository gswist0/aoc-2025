//> using scala 3.7.4
import scala.io.Source
@main
def main(): Unit = 
  val file: String = "input.txt"
  val lines: String = Source.fromResource(file).mkString
  val linesSplit = lines.split(",").map(x => {
    val split = x.split("-")
    (split(0).trim().toLong,split(1).trim().toLong)
  })

  println(checkIncorrects(linesSplit,false))
  println(checkIncorrects(linesSplit,true))


def checkIncorrects(Ids: Array[(Long,Long)], part2: Boolean) : Long =
  var incorrect: Long = 0
  for (line <- Ids){
    for (num <- line._1 until line._2 + 1){
      val str = num.toString()
      val len = str.length()
      var declaredInvalid = false
      val limit = if part2 then len + 1 else 3
      for (possibleLength <- 2 until limit){
        if (len % possibleLength == 0 && !declaredInvalid) then {
          val chunkSize = len/possibleLength
          val chunks = str.grouped(chunkSize).toArray
          if chunks.distinct.length == 1 then {
            incorrect += num
            declaredInvalid = true
          }
        }
      }
    }
  }
  incorrect

