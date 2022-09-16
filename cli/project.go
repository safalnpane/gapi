package cli


import (
    "log"
    "os"
    "encoding/json"
    "fmt"
    "io/ioutil"
)


type Project struct {
    Name string
    Description string
    DefaultServer string
    Servers []Server
}


func (p *Project) AddServer(name, description, baseUrl string) {
    newServer := Server{
        Name: name,
        Description: description,
        BaseURL: baseUrl,
    }
    p.Servers = append(p.Servers, newServer)
}


func ProjectInit(name, description string) {
    // Initialise the gapi project
    // Creates necessary directories and files
    currentDirectory, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    // Create 'gapi_project' directory
    gapiDirectory := currentDirectory + string(os.PathSeparator) + "gapi_project"
    err = os.Mkdir(gapiDirectory, 0755)
    if err != nil {
        log.Fatal(err)
    }
    newProject := Project{
        Name: name,
        Description: description,
        DefaultServer: "localhost",
    }
    newProject.AddServer("localhost", "Local API server", "http://localhost:8000/api")
    newProject.Servers[0].AddHeader("Content-Type", "application/json")
    // Convert into json
    content, err := json.Marshal(newProject)
    if err != nil {
        log.Fatal(err)
    }
    projectConfPath := gapiDirectory + string(os.PathSeparator) + "project.json"
    err = ioutil.WriteFile(projectConfPath, content, 0644)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("[+] Project '%v' Initialised\n", name)
}
