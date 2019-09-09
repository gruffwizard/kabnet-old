package util

import (
  "log"
  "fmt"
  "os"
  "strings"
  "io/ioutil"
  "encoding/json"
  "gopkg.in/yaml.v3"
  "strconv"

)

type Text struct {
  lines []string
}



func ToForm(mac string,s string) string {

    r := mac[0:2] + s + mac[2:4] + s + mac[4:6] + s + mac[6:8] + s + mac[8:10] + s + mac[10:12]
    return r

}

func GetHomeDir() string {
  home,err := os.UserHomeDir()
  if err != nil { log.Fatal(err) }
  return home
}
func FindIDFile() string {

  home:=GetHomeDir()

  var ssh=home+"/.ssh"
  files, err := ioutil.ReadDir(ssh)
  if err != nil {
      log.Print("Can't find a local ID file")
      log.Fatal(err)
  }

  for _, file := range files {
      name:=file.Name()
      if strings.HasSuffix(name, ".pub") {
        // got a matching private key?
        priv := ssh+"/"+strings.TrimSuffix(name, ".pub")
          if _, err := os.Stat(priv); err == nil {
            return priv
          }
      }

  }
  return ""
}
func WriteFile(dir string,fname string, contents string) {

  WriteFile2(dir+"/"+fname,contents)
}
func WriteFile2(path string, contents string) {

  log.Printf("save file %s",path)

 if path=="/tmp/foo/openshift/bootstrap.ign"  { log.Panic()}

  f, err := os.Create(path)
  if err!=nil { log.Fatal(err)}
  defer f.Close()
  f.WriteString(contents)
}




func ToInt(n string) int {

  v, err := strconv.Atoi(n)

  if err == nil {return v}

  return -1

}

func SaveAsYaml (path string, d  interface{}) {
  data,err := yaml.Marshal(&d)
  if err != nil {
      log.Fatalf("cannot marshal data: %v", err)
  }
  WriteFile2(path,string(data))
}
func SaveAsJson (path string, d  interface{}) {
  data,err := json.Marshal(&d)
  if err != nil {
      log.Fatalf("cannot marshal data: %v", err)
  }
  WriteFile2(path,string(data))
}

func LoadFromJsonFile(ignfile string) map[string]interface{} {

  jsonFile, err := os.Open(ignfile)
  if  err!=nil  { log.Fatal(err)}
  defer jsonFile.Close()

  byteValue, _ := ioutil.ReadAll(jsonFile)

    var result map[string]interface{}
    json.Unmarshal([]byte(byteValue), &result)

    return  result
}

func CreateText() *Text {
  var t Text
  return &t
}

func LoadFile(file string) string {

  data, err := ioutil.ReadFile(file)

  if err!=nil {
    log.Printf("error loading file [%s]",file)
    log.Fatal(err)
  }

  return string(data)
}

func (t *Text) AsString() string {
  return strings.Join(t.lines,"\n")+"\n"
}

func (t *Text) Add(format string, a ...interface{}) {
  t.lines=append(t.lines,fmt.Sprintf(format,a...))
}

func CopyFile(from string,todir string, tofile string) {

  log.Printf("copy %s to %s/%s",from,todir,tofile)

  f :=LoadFile(from)
  WriteFile(todir,tofile,f)

}

func CreateDir(path string,dirname string) string {
  dir := path+"/"+dirname

  if _, err := os.Stat(dir); os.IsNotExist(err) {
    os.Mkdir(dir, 0777)
  }
  return dir
}

func FileMustExist(path string) {

  if _, err := os.Stat(path); os.IsNotExist(err) {
    log.Fatalf("required file %s does not exist",path)
  }

}

func FileExists(path string) bool {

  if _, err := os.Stat(path); os.IsNotExist(err) {
    return false
  }
  return true

}
func CreateFile(path string,fname string) *os.File {

  log.Printf("Create File %s/%s",path,fname)

  f, err := os.Create(path+"/"+fname)
  if err!=nil { log.Fatal(err)}
  return f

}

func Emit(f *os.File,format string, a ...interface{}) {

    _, err:=fmt.Fprintf(f,format+"\n",a...)
    if err != nil {
      log.Fatal(err)
    }
}

func FetchFiles(dir string, url string,files []string) {
    for _,n := range files {
      FetchFile(dir,url,n)
    }
}
