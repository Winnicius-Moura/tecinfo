# 🖥️ TecInfo — Open Source Learning Project

> Repositório oficial do curso **TecInfo 2025–2026**.  Disciplinas Fundamentos de Criação de Sites
> Aprenda na prática como funciona o Open Source.

---

## 📌 Sobre o Projeto

Este repositório foi criado para que os alunos do curso TecInfo vivenciem, na prática, o fluxo de trabalho real de um projeto open source — desde o fork até o pull request aprovado.

Aqui você vai aprender a:

- Fazer **fork** de um repositório
- Trabalhar com **branches**
- Enviar **pull requests**
- Receber e aplicar **code reviews**
- Colaborar em um projeto com múltiplos contribuidores

> Parte das atividades realizadas neste repositório **será avaliada com nota**. Fique atento às instruções de cada exercício.

---

## 📂 Estrutura do Repositório

```
tecinfo/
├── contributions-analysis/   # Back-end de análise de exercícios (Building...)
├── galery/                   # Cards pessoais criados pelos alunos após o fork
├── tutorial/                 # Códigos produzidos em sala de aula
└── documentation/            # Documentação geral do projeto
```

### `contributions-analysis`
Back-end em **Go** responsável por receber, processar e avaliar as contribuições dos alunos.

**Tecnologias:** Go · Gin · PostgreSQL · Docker

### `galery`
Após o fork, cada aluno cria seu próprio **card de apresentação** com suas informações pessoais (nome, linguagens que usa, redes sociais). Os cards aprovados são exibidos nessa galeria.

### `tutorial`
Códigos, exemplos e exercícios produzidos durante as aulas. Use como referência de estudo.

---

## 🚀 Como Participar

### 1. Faça o Fork
Clique em **Fork** no canto superior direito desta página para criar sua cópia do repositório.

### 2. Clone o seu fork
```bash
git clone https://github.com/SEU-USUARIO/tecinfo.git
cd tecinfo
```

### 3. Crie uma branch para sua contribuição
```bash
git checkout -b feat/meu-card
```

### 4. Faça suas alterações e commit
```bash
git add .
git commit -m "feat: add student card - Seu Nome"
```

### 5. Envie para o seu fork
```bash
git push origin feat/meu-card
```

### 6. Abra um Pull Request
Vá até o repositório original e clique em **New Pull Request**. Descreva o que você fez e aguarde o review.

---

## 📋 Regras para Contribuição

- Siga o padrão de **Conventional Commits** nas mensagens (`feat:`, `fix:`, `docs:`, etc.)
- Não altere arquivos fora do escopo da sua tarefa
- Pull Requests sem descrição **não serão aceitos**

---

## 🎓 Avaliação

As notas serão baseadas em critérios como:

| Critério | Descrição |
|---|---|
| ✅ Pull Request enviado | Contribuição submetida corretamente |
| ✅ Qualidade do código | Boas práticas e organização |
| ✅ Descrição do PR | Explicação clara do que foi feito |
| ✅ Exercícios analisados | Resultados pelo sistema de análise |

---

## 🛠️ Rodando o Back-end Localmente

Pré-requisitos: **Docker** e **Docker Compose**

```bash
cd contributions-analysis
docker-compose up -d
make run
```

<p align="center">
  Feito com dedicação para o curso <strong>TecInfo 2025–2026</strong> 🚀
</p>
