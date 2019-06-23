package main

import(
  "encoding/csv"
  "log"
  "os"
)

var sample = []string{
  "sample_1",
  "sample_2",
  "sample_3",
}

func main(){
  file, err := os.Create("./madeCsv/testsample.csv")
  if err != nil{
    log.Println(err)
  }
  defer file.Close()

  writer := csv.NewWriter(file) // utf-8
  writer.UseCRLF = true
  writer.Comma = ';'
  writer.Write(sample)
  writer.Flush()
}

