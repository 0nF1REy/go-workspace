<p align="center">
    <img src="./docs/images/logotype/go-workspace.png" width="200" alt="Logotipo — Go Workspace" />
</p>

<div align="center">

![Maintenance](https://img.shields.io/maintenance/yes/2026?style=for-the-badge)
![License MIT](https://img.shields.io/badge/license-MIT-blue?style=for-the-badge)

</div>

## 🧭 Guia de Navegação

<a name="guia-navegacao"></a>

- [📖 Descrição](#descricao)
- [🛠️ Build & Execução](#build-execucao)
- [🏢 Projetos](#projetos)
- [🗺️ Diagramas](#diagramas)
- [👤 Sobre o Desenvolvedor](#sobre-o-desenvolvedor)
- [📜 Licença](#licenca)

---

## 📖 Descrição <a name="descricao"></a>

O repositório **go-workspace** é um ambiente completo para estudos, prototipagem e desenvolvimento de APIs, exemplos didáticos e projetos reais em Go, com integração a Bazel, Docker/Podman e PostgreSQL.

---

## 🛠️ Build & Execução <a name="build-execucao"></a>

| Comando       | Descrição rápida                  |
| ------------- | --------------------------------- |
| `make build`  | Compila a aplicação               |
| `make dev`    | Hot reload API (Air + Bazel + DB) |
| `make prod`   | Modo interativo (menu)            |
| `make db-up`  | Inicia o banco PostgreSQL         |
| `make api-up` | Sobe containers da stack          |
| ...           | Outros comandos na documentação   |

🔗 [Documentação completa](docs/guides/build-execution.md)

---

## 🏢 Projetos <a name="projetos"></a>

| Projeto         | Descrição rápida                      | Stack/Porta   |
| --------------- | ------------------------------------- | ------------- |
| 01_products_api | API REST, JWT, PostgreSQL, containers | docker/podman |
| ping-api        | Serviço ping HTTP simples             | 8080          |

🔗 [Documentação completa](docs/guides/projects.md)

## 🗺️ Diagramas <a name="diagramas"></a>

Todos os diagramas visuais e fluxos de arquitetura estão centralizados em:

👉 [docs/guides/architecture-visuals.md](docs/guides/architecture-visuals.md)

Esse documento contém:

- Arquitetura Geral (Mermaid)
- Fluxo de Sequência (Mermaid)
- Diagramas e guias visuais futuros

---

## 👤 Sobre o Desenvolvedor <a name="sobre-o-desenvolvedor"></a>

<table align="center">
  <tr>
    <td align="center">
        <br>
        <a href="https://github.com/0nF1REy" target="_blank">
          <img src="./docs/images/developer/alan-ryan.jpg" height="160" alt="Foto — Alan Ryan">
        </a>
        </p>
        <a href="https://github.com/0nF1REy" target="_blank">
          <strong>Alan Ryan</strong>
        </a>
        </p>
        ☕ Peopleware | Tech Enthusiast | Code Slinger ☕
        <br>
        Apaixonado por código limpo, arquitetura escalável e experiências digitais envolventes
        </p>
          Conecte-se comigo:
        </p>
        <a href="https://www.linkedin.com/in/alan-ryan-b115ba228" target="_blank">
          <img src="https://img.shields.io/badge/LinkedIn-Alan_Ryan-0077B5?style=flat&logo=linkedin" alt="LinkedIn">
        </a>
        <a href="https://gitlab.com/alanryan619" target="_blank">
          <img src="https://img.shields.io/badge/GitLab-@0nF1REy-FCA121?style=flat&logo=gitlab" alt="GitLab">
        </a>
        <a href="mailto:alanryan619@gmail.com" target="_blank">
          <img src="https://img.shields.io/badge/Email-alanryan619@gmail.com-D14836?style=flat&logo=gmail" alt="Email">
        </a>
        </p>
    </td>
  </tr>
</table>

</div>

---

## 📜 Licença <a name="licenca"></a>

Este projeto está sob a **licença MIT**. Consulte o arquivo **[LICENSE](LICENSE)** para obter mais detalhes.

> ℹ️ **Aviso de Licença:** &copy; 2026 Alan Ryan da Silva Domingues. Este projeto está licenciado sob os termos da licença MIT. Isso significa que você pode usá-lo, copiá-lo, modificá-lo e distribuí-lo com liberdade, desde que mantenha os avisos de copyright.

⭐ Se este repositório foi útil para você, considere dar uma estrela!
