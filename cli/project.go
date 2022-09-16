package cli


import (
    "os"
    "encoding/json"
    "fmt"
    "io/ioutil"
)

const (
    projectDirectoryName = "gapi_project"
    projectConfigName = "project.json"
)

var (
    currentProjectDirectory string
    currentProjectConfig string
)


type Project struct {
    Name string
    Description string
    DefaultServer string
    Servers []Server
}


// Check if 'gapi_project' folder exists on the current directory.
// Also make sure, the 'project.json' exists.
func IsProject() bool {
    currentDirectory, err := os.Getwd()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    currentProjectDirectory = currentDirectory + string(os.PathSeparator) + projectDirectoryName
    currentProjectConfig = currentProjectDirectory + string(os.PathSeparator) + projectConfigName
    if _, err = os.Stat(currentProjectConfig); err == nil {
        return true
    }
    return false
}


func (p *Project) AddServer(name, description, baseUrl string) {
    newServer := Server{
        Name: name,
        Description: description,
        BaseURL: baseUrl,
    }
    p.Servers = append(p.Servers, newServer)
}


// Convert the project struct into JSON and write into the project file
func (p *Project) WriteProject() {
    content, err := json.Marshal(p)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    err = ioutil.WriteFile(currentProjectConfig, content, 0644)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}


// Read the project config and populate the project struct
func (p *Project) ReadProject() {
    content, err := ioutil.ReadFile(currentProjectConfig)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    err = json.Unmarshal(content, p)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}


func ProjectInit(name, description string) {
    // Initialise the gapi project
    // Creates necessary directories and files
    if IsProject() {
        fmt.Println("[-] Project Already exists.")
        os.Exit(1)
    }
    err := os.Mkdir(currentProjectDirectory, 0755)
    if err != nil {
        fmt.Println(err)
    }
    newProject := Project{
        Name: name,
        Description: description,
        DefaultServer: "localhost",
    }
    newProject.AddServer("localhost", "Local API server", "http://localhost:8000/api")
    newProject.Servers[0].AddHeader("Content-Type", "application/json")
    newProject.WriteProject()
    fmt.Printf("[+] Project '%v' Initialised\n", name)
}
