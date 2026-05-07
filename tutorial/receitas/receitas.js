// Abre o modal com o título e texto da receita clicada
function abrirModal(titulo, texto) {
  document.getElementById('modal-titulo').textContent = titulo;
  document.getElementById('modal-texto').textContent = texto;
  document.getElementById('fundo-modal').classList.add('ativo');
}

// Fecha o modal
function fecharModal() {
  document.getElementById('fundo-modal').classList.remove('ativo');
}
