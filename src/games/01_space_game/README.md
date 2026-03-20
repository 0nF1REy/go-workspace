# 🚀 Space Game (Ebiten)

Um jogo 2D estilo arcade desenvolvido em Go utilizando a biblioteca Ebiten.  
O projeto foi estruturado com foco em organização, portabilidade e distribuição multiplataforma.

---

## 🎮 Sobre o projeto

O **Space Game** é um jogo simples de nave espacial onde o jogador deve desviar de meteoros e sobreviver o máximo possível.

O projeto utiliza:

- Go (linguagem principal)
- Ebiten (engine de jogo 2D)
- `embed.FS` (assets embutidos no binário)
- Estrutura modular (`/src`)
- Build multiplataforma (Linux, Windows e WebAssembly)

---

## 📚 Documentação

Abaixo estão os guias disponíveis para compilação, execução e resolução de problemas:

| 📄 Documento                                                    | 📌 Descrição                                                           |
| --------------------------------------------------------------- | ---------------------------------------------------------------------- |
| [Build e Distribuição](docs/guides/BUILD_AND_DISTRIBUTION.md)   | Guia completo para compilar e distribuir o jogo (Linux, Windows e Web) |
| [OpenGL Troubleshooting](docs/guides/OPENGL_TROUBLESHOOTING.md) | Soluções para problemas com GPUs antigas e OpenGL                      |

---

## ▶️ Como rodar o projeto

```bash
go run main.go
```

---

## 🧱 Estrutura interna

- `main.go` → ponto de entrada
- `src/` → lógica do jogo (player, meteors, etc.)
- `docs/` → documentação técnica

---

## 🌍 Plataformas suportadas

- ✅ Linux
- ✅ Windows
- ⚠️ Web (WASM - experimental)

---

## 📦 Distribuição

O projeto suporta:

- Binários nativos
- Cross-compilation
- WebAssembly (WASM)

👉 Consulte:  
[Guia de Build e Distribuição](docs/guides/BUILD_AND_DISTRIBUTION.md)

---

## ⚠️ Problemas conhecidos

Problemas com OpenGL podem ocorrer em GPUs antigas.

👉 Consulte:  
[OpenGL Troubleshooting](docs/guides/OPENGL_TROUBLESHOOTING.md)

---

## 📌 Objetivo do projeto

Este projeto foi desenvolvido com foco em:

- aprendizado de Go
- arquitetura de projetos
- desenvolvimento de jogos 2D
- distribuição multiplataforma
