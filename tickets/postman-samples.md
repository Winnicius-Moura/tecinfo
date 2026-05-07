# Postman — Exemplos de submissão HTML/CSS

**Endpoint:** `POST /contributions-analysis/api/v1/test/html-css/submit`  
**Content-Type:** `application/json`

---

## Nota baixa (~20%) — Aluno não leu o ticket

Problemas: usa classe de container errada (`badges-list`, `links-sociais`), `<h4>` com texto livre em português, 3 badges (acima do limite), nenhum `.social-link`, CSS mínimo e sem `border-radius`.

```json
{
  "contributor_id": "00000000-0000-0000-0000-000000000001",
  "html_content": "<article>\n  <h3>Marcos Rocha</h3>\n  <p>Iniciante em HTML</p>\n  <h4>Linguagens que uso</h4>\n  <div class=\"badges-list\">\n    <div class=\"badge\" style=\"background-color: #f7df1e; color: black\">JavaScript</div>\n    <div class=\"badge\" style=\"background-color: #3874a4; color: white\">Python</div>\n    <div class=\"badge\" style=\"background-color: #b07219; color: white\">Java</div>\n  </div>\n  <h4>Redes sociais</h4>\n  <div class=\"links-sociais\">\n    <a href=\"https://github.com/marcosrocha\">GitHub</a>\n    <a href=\"https://linkedin.com/in/marcosrocha\">LinkedIn</a>\n  </div>\n</article>\n<style>\n  body { font-family: sans-serif; }\n  .badge { padding: 0.5rem; }\n  a { color: blue; }\n</style>"
}
```

**Checks esperados — passou:**
- `article-wrapper` (5)
- `h3-username` (5)
- `paragraph-bio` (3)
- `has-style-block` (3)

**Checks esperados — falhou:** todos os demais  
**Score estimado: 16/100**

---

## Nota média (~63%) — Aluno leu o ticket mas pulou detalhes

Problemas: segundo `<h4>` com texto errado, usa `<div class="social-container">` em vez de `<section>`, badges sem `color` inline, nenhum `<img class="social-icon">`, `target="_blank"` ausente nos links, sem CSS do `article`.

```json
{
  "contributor_id": "00000000-0000-0000-0000-000000000002",
  "html_content": "<article>\n  <h3>Camila Torres</h3>\n  <p>Aprendendo desenvolvimento web</p>\n  <h4>Programming languages I use</h4>\n  <section class=\"container\">\n    <div class=\"badge\" style=\"background-color: #f7df1e\">JavaScript</div>\n    <div class=\"badge\" style=\"background-color: #3874a4\">Python</div>\n  </section>\n  <h4>Links das minhas redes</h4>\n  <div class=\"social-container\">\n    <a class=\"social-link\" href=\"https://github.com/camilatorres\">GitHub</a>\n    <a class=\"social-link\" href=\"https://linkedin.com/in/camilatorres\">LinkedIn</a>\n  </div>\n</article>\n<style>\n  body { font-family: sans-serif; }\n  h3 { margin: 0.6rem 0 0.25rem; }\n  p { margin: 0; }\n  h4 { margin: 0.6rem 0 0.25rem; }\n  .container { display: flex; flex-wrap: wrap; gap: 1rem; }\n  .badge { padding: 0.5rem; border-radius: 0.25rem; }\n  .icon { width: 2rem; }\n  .social-container { display: flex; flex-wrap: wrap; gap: 0.5rem; }\n  a { display: flex; align-items: center; text-decoration: none; color: #374151; background: #f3f4f6; border-radius: 999px; padding: 0.25rem 0.65rem; }\n</style>"
}
```

**Checks esperados — passou:**
- `article-wrapper` (5), `h3-username` (5), `paragraph-bio` (3)
- `section-container` (6) — usa `<section>` corretamente
- `structure-order` (5)
- `badge-count` (9), `badge-background-color` (7)
- `social-link-count` (9), `social-link-href` (4)
- `has-style-block` (3), `badge-css` (4), `container-css` (3)

**Checks esperados — falhou:**
- `h4-titles-text` — segundo h4 errado
- `section-social-container` — usa div
- `badge-text-color` — sem `color` inline
- `social-link-target` — sem `target="_blank"`
- `social-link-icon` — sem `<img class="social-icon">`
- `article-css` — sem seletor `article` no CSS
- `social-link-css` — CSS no seletor `a` e não `.social-link`

**Score estimado: 63/100**

---

## Nota total (100%) — Implementação fiel ao ticket

Implementação completa: `<section>` semântico nos dois containers, textos exatos nos `<h4>`, badges com `background-color` e `color` inline, links com `href` real + `target="_blank"` + `<img class="social-icon">`, CSS completo com seletor `article`, `<style>` após `</article>`.

```json
{
  "contributor_id": "00000000-0000-0000-0000-000000000003",
  "html_content": "<article>\n  <h3>Lucas Fernandes</h3>\n  <p>Desenvolvedor apaixonado por tecnologia</p>\n  <h4>Programming languages I use</h4>\n  <section class=\"container\">\n    <div class=\"badge\" style=\"background-color: #3874a4; color: white\">Go</div>\n    <div class=\"badge\" style=\"background-color: #f7df1e; color: black\">JavaScript</div>\n  </section>\n  <h4>Social Links</h4>\n  <section class=\"social-container\">\n    <a class=\"social-link\" href=\"https://github.com/lucasfernandes\" target=\"_blank\">\n      <img class=\"social-icon\" src=\"https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/github/github-original.svg\" />\n      GitHub\n    </a>\n    <a class=\"social-link\" href=\"https://linkedin.com/in/lucasfernandes\" target=\"_blank\">\n      <img class=\"social-icon\" src=\"https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/linkedin/linkedin-original.svg\" />\n      LinkedIn\n    </a>\n  </section>\n</article>\n<style>\n  body { font-family: sans-serif; }\n  article { max-width: 16rem; padding: 1rem; background: #fefefe; border-radius: 0.5rem; box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1); }\n  h3 { margin: 0.6rem 0 0.25rem; }\n  p { margin: 0; }\n  h4 { margin: 0.6rem 0 0.25rem; }\n  .container { display: flex; flex-wrap: wrap; gap: 1rem; }\n  .badge { padding: 0.5rem; border-radius: 0.25rem; }\n  .icon { width: 2rem; }\n  .social-container { display: flex; flex-wrap: wrap; gap: 0.5rem; }\n  .social-link { display: flex; align-items: center; gap: 0.4rem; text-decoration: none; color: #374151; font-size: 0.8rem; font-weight: 500; background: #f3f4f6; border-radius: 999px; padding: 0.25rem 0.65rem 0.25rem 0.35rem; transition: background 0.2s; }\n  .social-link:hover { background: #e5e7eb; }\n  .social-icon { width: 1.1rem; height: 1.1rem; }\n</style>"
}
```

**Todos os checks esperados — passou:**  
`article-wrapper` (5) + `h3-username` (5) + `paragraph-bio` (3) + `h4-titles-text` (8) + `section-container` (6) + `section-social-container` (6) + `structure-order` (5) + `badge-count` (9) + `badge-background-color` (7) + `badge-text-color` (6) + `social-link-count` (9) + `social-link-href` (4) + `social-link-target` (5) + `social-link-icon` (4) + `has-style-block` (3) + `article-css` (3) + `badge-css` (4) + `container-css` (3) + `social-link-css` (5)

**Score estimado: 100/100**
