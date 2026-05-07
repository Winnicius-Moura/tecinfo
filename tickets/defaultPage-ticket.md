# DefaultPage Layout — Painel Administrativo (17/04/2026)

> **Contexto da atividade:** Você irá construir do zero a estrutura base de um painel administrativo. Essa é a chamada *defaultPage* — a página que serve como template para todas as telas internas do sistema. O que muda de página para página é apenas o conteúdo central; a sidebar, o header e a topbar permanecem iguais. O objetivo é aprender a **dividir uma interface em regiões visuais** com HTML semântico e em seguida aplicar os estilos CSS para reproduzir o layout do design de referência.

---

### [Feature] — DefaultPage — Implementar estrutura base do painel em HTML e CSS

#### Contexto (User Story)

Como desenvolvedor(a) front-end,  
Eu quero implementar a estrutura base do painel administrativo em HTML e CSS,  
Para que essa página sirva como template fixo para todas as telas internas do sistema.

---

#### Design de Referência

O layout é dividido em **duas colunas**: uma barra lateral fixa à esquerda (sidebar) e uma área principal que ocupa o restante da tela à direita. A área principal, por sua vez, é dividida verticalmente em três faixas: o header (banner do sistema), a topbar (breadcrumb e botões) e o conteúdo central.

```
+------------------+------------------------------------------+
|                  |       HEADER — banner do sistema         |
|   SIDEBAR        +------------------------------------------+
|   Logo           +                                          +
|   > Times        |   Bem-vindo(a) ao seu painel, Admin!     |
|   > Importações  |   admin@email.com                        |
|                  |                                          |
|                  |   [ ☁  Importações              ]       |
|                  |   [ 👥 Gerenciamento de Times   ]       |
|                  |                                          |
+------------------+------------------------------------------+
```

---

#### Estrutura de Pastas do Projeto

```
defaultPage/
  pages/
    defaultPage.html     ← página principal (painel)
    importacoes.html     ← página de importações
  css/
    defaultPage.css      ← todos os estilos
  defaultPage-ticket.md
```

O CSS fica em um arquivo separado na pasta `css/`. Cada HTML o referencia com:
```html
<link rel="stylesheet" href="../css/defaultPage.css">
```
O `../` sobe um nível (saindo de `pages/`) para então entrar em `css/`.

---

## Tarefa 1 — HTML: Estrutura da página

> **Objetivo desta etapa:** Montar os elementos HTML **sem nenhum estilo**. Nesta fase o resultado visual será uma lista vertical sem formatação — isso é normal e esperado. O foco é entender a semântica e a hierarquia dos elementos.

O arquivo deve se chamar `pages/defaultPage.html`. O `<head>` deve conter a tag `<link>` apontando para o CSS externo — **não use `<style>` diretamente no HTML**.

**Estrutura HTML — elementos obrigatórios e ordem:**

| Ordem | Elemento | Classe | Função |
|---|---|---|---|
| 1 | `<div>` | `class="app-layout"` | Container raiz — envolve toda a página |
| 2 | `<aside>` | `class="sidebar"` | Coluna da barra lateral |
| 3 | `<div>` | `class="sidebar-logo"` | Área do logo dentro da sidebar |
| 4 | `<nav>` | `class="sidebar-nav"` | Grupo de links de navegação |
| 5 | `<a>` (×2) | `class="nav-item"` | Links do menu (Times e Importações) |
| 6 | `<div>` | `class="main-wrapper"` | Coluna da área principal |
| 7 | `<header>` | `class="page-header"` | Banner do sistema no topo |
| 8 | `<div>` | `class="topbar"` | Faixa com breadcrumb e botões |
| 9 | `<div>` | `class="breadcrumb"` | Caminho de navegação (🏠 > Painel) |
| 10 | `<div>` | `class="topbar-actions"` | Agrupa os botões de ação |
| 11 | `<button>` | `class="btn-settings"` | Botão de configurações (⚙) |
| 12 | `<button>` | `class="btn-logout"` | Botão de logout |
| 13 | `<main>` | `class="page-content"` | Área do conteúdo da página |
| 14 | `<h1>` | — | Título de boas-vindas |
| 15 | `<p>` | `class="user-email"` | Email do usuário logado |
| 16 | `<div>` | `class="cards-grid"` | Container que agrupa os cards |
| 17 | `<div>` (×2) | `class="card"` | Card de acesso rápido a uma funcionalidade |
| 18 | `<div>` | `class="card-icon"` | Ícone do card |
| 19 | `<div>` | `class="card-body"` | Texto do card |
| 20 | `<h3>` | — | Título do card |
| 21 | `<p>` | — | Descrição do card |

**Importante — primeiro link ativo:** O primeiro `<a class="nav-item">` deve ter também a classe `active`: `class="nav-item active"`. Isso será estilizado no CSS para indicar a página atual.

---

## Tarefa 2 — CSS: Estilo do layout

> **Objetivo desta etapa:** Aplicar os estilos CSS para transformar a estrutura HTML no layout do design de referência. Cada bloco abaixo corresponde a um seletor e descreve o comportamento visual esperado — cabe a você escrever o CSS.

---

### Reset global

**Seletor `*`**  
Remova as margens e paddings padrão do navegador (`margin: 0`, `padding: 0`). Defina `box-sizing: border-box` para que bordas e paddings sejam incluídos no cálculo do tamanho dos elementos.

---

### Body

**Seletor `body`**  
Sem margem. Fonte `sans-serif`. Cor de fundo levemente acinzentada (`#f0f2f5`).

---

### Layout principal (`.app-layout`)

**Seletor `.app-layout`**  
Use `display: flex` para colocar a sidebar e o conteúdo principal lado a lado (em linha). A altura mínima deve ocupar a tela inteira (`min-height: 100vh`).

---

### Sidebar

**Seletor `.sidebar`**  
Largura fixa de `180px`. Fundo branco. Padding interno de `1rem`. Borda direita sutil (`1px solid #e0e0e0`).

**Seletor `.sidebar-logo`**  
Texto centralizado. Padding vertical de `1rem` (`1rem 0`). Borda inferior sutil. Texto em negrito com cor roxa (`#5b4fcf`).

**Seletor `.sidebar-nav`**  
Use `display: flex` com `flex-direction: column` para empilhar os links verticalmente. Margem superior de `1rem`.

**Seletor `.nav-item`**  
`display: block`. Padding de `0.75rem 1rem`. Sem sublinhado (`text-decoration: none`). Cor do texto cinza escuro (`#444`). Cantos arredondados de `0.4rem`.

**Seletor `.nav-item:hover`**  
No hover, o fundo muda para azul bem claro (`#eef0ff`) e o texto fica roxo (`#5b4fcf`).

**Seletor `.nav-item.active`**  
Mesmo visual do hover, mas sempre visível: fundo `#eef0ff`, texto roxo, negrito.

---

### Área principal

**Seletor `.main-wrapper`**  
`flex: 1` para ocupar todo o espaço restante ao lado da sidebar. `display: flex` com `flex-direction: column` para empilhar header, topbar e conteúdo.

---

### Header

**Seletor `.page-header`**  
Fundo com gradiente horizontal (`linear-gradient(to right, #1a2a6c, #2e6cb8)`). Padding de `1rem 2rem`. Cor de texto branco. Texto centralizado.

---

### Topbar

**Seletor `.topbar`**  
`display: flex`. Alinhamento vertical centralizado (`align-items: center`). Espaço distribuído entre os filhos (`justify-content: space-between`). Padding de `0.75rem 2rem`. Fundo branco. Borda inferior sutil.

**Seletor `.breadcrumb`**  
Tamanho de fonte pequeno (`0.85rem`). Cor cinza (`#888`).

**Seletor `.topbar-actions`**  
`display: flex`. Gap de `0.75rem` entre os botões.

**Seletor `.btn-settings`**  
Fundo transparente. Sem borda. Cursor `pointer`. Tamanho de fonte `1rem`.

**Seletor `.btn-logout`**  
Fundo transparente. Borda `1px solid #ccc`. Cantos arredondados (`0.4rem`). Padding pequeno (`0.3rem 0.8rem`). Cursor `pointer`. Tamanho de fonte `0.85rem`.

**Seletor `.btn-logout:hover`**  
Fundo levemente cinza (`#f5f5f5`).

---

### Conteúdo

**Seletor `.page-content`**  
Padding de `2rem`. `flex: 1` para ocupar o espaço restante na coluna.

**Seletor `.user-email`**  
Cor cinza (`#666`). Margem superior pequena (`0.25rem`).

---

### Cards

**Seletor `.cards-grid`**  
`display: flex` com `flex-direction: column`. Gap de `1rem` entre os cards. Margem superior de `1.5rem`.

**Seletor `.card`**  
`display: flex` em linha. `align-items: center`. Gap de `1rem` entre ícone e texto. Fundo branco. Cantos arredondados (`0.5rem`). Padding de `1.25rem`. Sombra suave (`0 2px 6px rgba(0, 0, 0, 0.08)`).

**Seletor `.card-icon`**  
Tamanho de fonte `1.5rem`. Cor roxa (`#5b4fcf`). `flex-shrink: 0` para que o ícone não encolha.

**Seletor `.card-body h3`**  
Margem zero. Cor escura (`#222`).

**Seletor `.card-body p`**  
Margem superior de `0.25rem`. Cor cinza (`#666`). Tamanho de fonte `0.9rem`.

---

#### Critério de Conclusão — Tarefas 1 e 2

- [ ] A página possui sidebar fixa à esquerda e área principal à direita
- [ ] O header exibe um banner com gradiente azul
- [ ] O link ativo na sidebar tem destaque visual
- [ ] Os cards exibem ícone e texto lado a lado com sombra suave
- [ ] A estrutura se mantém coesa visualmente ao trocar apenas o conteúdo dentro de `.page-content`
- [ ] O CSS está em arquivo externo (`css/defaultPage.css`), sem `<style>` no HTML

---

## Tarefa 3 — Segunda página: Importações

> **Objetivo desta etapa:** Criar uma segunda página (`pages/importacoes.html`) que reutiliza o mesmo layout base (sidebar + header) e demonstra que apenas o conteúdo interno muda. A página terá um botão que, ao ser clicado, exibe uma mensagem de alerta na tela.

**O que deve ser implementado:**

| Elemento | Classe / Atributo | Detalhe |
|---|---|---|
| Sidebar e header | — | Idênticos ao `defaultPage.html`; apenas o `active` muda para "Importações" |
| `<h1>` | — | Título: `Importações` |
| `<p>` | `class="user-email"` | Subtítulo descritivo |
| `<div>` | `id="alerta"` + `class="alert"` | Mensagem de sucesso; começa oculto (`style="display: none"`) |
| `<button>` | `class="btn-import"` | Texto: `Importar dados`; ao clicar, torna o `#alerta` visível |

**O clique do botão** deve usar o atributo `onclick` diretamente no HTML:
```html
onclick="document.getElementById('alerta').style.display = 'block'"
```

**CSS necessário em `defaultPage.css`:**

**`.btn-import`**  
Fundo roxo (`#5b4fcf`). Texto branco. Sem borda. Cantos arredondados (`0.4rem`). Padding de `0.6rem 1.4rem`. Cursor `pointer`. Margem superior de `1.5rem`.

**`.btn-import:hover`**  
Fundo ligeiramente mais escuro (`#4a3fbf`).

**`.alert`**  
Fundo verde claro (`#e6f4ea`). Borda `1px solid #a8d5b5`. Cantos arredondados (`0.4rem`). Padding de `1rem 1.25rem`. Cor do texto verde escuro (`#276c3f`). Margem superior de `1.5rem`.

#### Critério de Conclusão — Tarefa 3

- [ ] O arquivo `pages/importacoes.html` existe e carrega o mesmo CSS externo
- [ ] O link "Importações" na sidebar está com a classe `active`
- [ ] O botão está visível na página
- [ ] Ao clicar no botão, a mensagem de alerta aparece na tela
- [ ] O alerta começa oculto (invisível antes do clique)
