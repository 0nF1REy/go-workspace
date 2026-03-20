[⬅ Voltar para o README](../../README.md)

---

# Guia de Compilação e Distribuição — Space Game (Ebiten)

Este documento detalha o processo completo de compilação e distribuição do jogo Space Game, desenvolvido em Go utilizando a biblioteca Ebiten.

O projeto faz uso de `embed.FS` para gerenciamento de assets, tornando os binários e artefatos gerados totalmente autossuficientes, sem necessidade de pastas externas de assets.

---

## 📋 Índice

- [Requisitos](#requisitos)
- [Compilação Nativa para Linux](#compilacao-linux)
- [Cross-Compilation para Windows](#cross-windows)
- [Distribuição Web (WASM)](#wasm)
- [Observações sobre embed.FS](#embedfs)

---

## ⚙️ Requisitos

- Go 1.20 ou superior

---

## 🐧 Compilação Nativa para Linux

Execute o comando abaixo na raiz do projeto para gerar um binário Linux nativo:

```sh
go build -ldflags="-s -w" -o space-game-linux
```

O binário gerado (`space-game-linux`) estará pronto para distribuição e execução em sistemas Linux x86_64.

---

## 🪟 Cross-Compilation para Windows (a partir do Linux)

Para compilar o jogo para Windows a partir de um ambiente Linux:

```sh
GOOS=windows GOARCH=amd64 go build -o space-game-win.exe
```

Ou para gerar um executável sem abrir terminal:

```sh
GOOS=windows GOARCH=amd64 go build -ldflags="-H=windowsgui" -o space-game-win.exe
```

- O flag `-ldflags="-H=windowsgui"` remove o console e faz o app rodar como GUI nativa no Windows.
- O arquivo `space-game-win.exe` estará pronto para distribuição em sistemas Windows 64 bits.

---

## 🌐 Distribuição Web (WebAssembly — WASM)

O Go permite compilar aplicações para WebAssembly, possibilitando rodar o jogo diretamente no navegador.

> ⚠️ Nota: O suporte a WebAssembly com Ebiten pode ser limitado ou experimental dependendo da versão.

---

### 1. Compilando para WASM

```sh
GOOS=js GOARCH=wasm go build -o space-game.wasm
```

---

### 2. Copiando o runtime do Go

O arquivo `wasm_exec.js` é necessário para executar o WASM no navegador:

```sh
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

---

### 3. Estrutura HTML Boilerplate

Crie um arquivo `index.html` com o seguinte conteúdo:

```html
<!DOCTYPE html>
<html lang="pt-br">
  <head>
    <meta charset="UTF-8" />
    <title>Space Game (Ebiten WASM)</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <style>
      html,
      body {
        height: 100%;
        margin: 0;
        background: #000;
      }
      canvas {
        display: block;
        margin: 0 auto;
        background: #111;
      }
    </style>
  </head>
  <body>
    <script src="wasm_exec.js"></script>
    <script>
      const go = new Go();

      WebAssembly.instantiateStreaming(
        fetch("space-game.wasm"),
        go.importObject,
      )
        .then((result) => {
          go.run(result.instance);
        })
        .catch(async () => {
          const resp = await fetch("space-game.wasm");
          const bytes = await resp.arrayBuffer();
          const result = await WebAssembly.instantiate(bytes, go.importObject);
          go.run(result.instance);
        });
    </script>

    <noscript>Ative o JavaScript para jogar.</noscript>
  </body>
</html>
```

---

### 4. Executando localmente

> ⚠️ Não abra o arquivo HTML diretamente (file://). Use sempre um servidor HTTP.

Devido a restrições do navegador, é necessário servir os arquivos via HTTP.

Crie um arquivo `server.go` com o seguinte conteúdo:

```go
package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	fs := http.FileServer(http.Dir("."))

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".wasm") {
			w.Header().Set("Content-Type", "application/wasm")
		}
		fs.ServeHTTP(w, r)
	})

	http.Handle("/", handler)

	log.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Execute:

```sh
go run server.go
```

Depois, acesse no navegador:

```
http://localhost:8080
```

---

### 5. Distribuição

Distribua os seguintes arquivos juntos:

- `space-game.wasm`
- `wasm_exec.js`
- `index.html`

Não é necessário distribuir pastas de assets, pois tudo está embutido via `embed.FS`.

---

## 📦 Observações sobre embed.FS

- Todos os assets (imagens, sons, fontes, etc.) são incorporados ao binário/WASM via `embed.FS`.
- Não é necessário distribuir pastas externas de assets.
- Os artefatos gerados são autossuficientes e prontos para uso imediato em qualquer ambiente de destino.

---

## 📌 Dicas Extras

### Multi-arch (exemplo ARM64)

```sh
GOOS=linux GOARCH=arm64 go build -o space-game-linux-arm64
```

---

### Compactação para distribuição

```sh
tar -czvf space-game-linux.tar.gz space-game-linux
```
