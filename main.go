// main.go
package main

import (
	"llmbridge/config"
	"llmbridge/server"
	"log"
)

func main() {
    // Carregar configuració
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("failed to load config: %v", err)
    }


    // Inicialitzar el servidor
    srv := server.NewServer(cfg)
    
    // Configurar middlewares i rutes
    if err := srv.Setup(); err != nil {
        log.Fatalf("failed to set up middlewares: %v", err)
    }
    
    // Iniciar el servidor
    if err := srv.Run(); err != nil {
        log.Fatalf("failed to start server: %v", err)
    }
        
}