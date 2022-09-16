package cli


import (
    "log"
    "os"
)


type Project struct {
    Name string
    Description string
    DefaultServer string
    Servers []Server
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
    // Create project.json file
    localhostServer := Server{
        Name: "localhost",
        Description: "Local API server",
        BaseURL: "http://localhost:8000",
    }
    localhostServer.AddHeader("Content-Type", "application/json")
    newProject := Project{
        Name: name,
        Description: description,
        DefaultServer: "localhost",
        Servers: []Server{localhostServer,},
    }
    log.Fatal(newProject)
}
