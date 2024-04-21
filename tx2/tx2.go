package tx2

import (
    "fmt"
    "os"
    "github.com/fsnotify/fsnotify"
)

func MonitoraTx2() {
    // Especifica o diretório a ser monitorado
    directory := "./envia"

    // Inicia o monitoramento do diretório
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        fmt.Println("Erro ao iniciar o monitoramento:", err)
        return
    }
    defer watcher.Close()

    // Adiciona o diretório ao watcher
    err = watcher.Add(directory)
    if err != nil {
        fmt.Println("Erro ao adicionar o diretório ao watcher:", err)
        return
    }

    // Loop infinito para observar eventos
    for {
        select {
        case event, ok := <-watcher.Events:
            if !ok {
                return
            }
            // Verifica se é um evento de criação de arquivo
            if event.Op&fsnotify.Create == fsnotify.Create {
                // Pega o nome do arquivo
                filename := event.Name
                // Lê o conteúdo do arquivo
                content, err := readFile(filename)
                if err != nil {
                    fmt.Println("Erro ao ler o arquivo:", err)
                } else {
                    fmt.Println("Arquivo adicionado:", filename)
                    fmt.Println("Conteúdo do arquivo:")
                    fmt.Println(content)
                }
            }
        case err, ok := <-watcher.Errors:
            if !ok {
                return
            }
            fmt.Println("Erro:", err)
        }
    }
}

func readFile(filename string) (string, error) {
    // Lê o conteúdo do arquivo
    content, err := os.ReadFile(filename)
    if err != nil {
        return "", err
    }
    return string(content), nil
}
