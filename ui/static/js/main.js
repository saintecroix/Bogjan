const openModalButtons = document.querySelectorAll('.open-modal-button');
const modal = document.getElementById('modal');

openModalButtons.forEach((openModalButtons) => {
    openModalButtons.addEventListener('click', () => {
        modal.style.display = 'flex';
    });
});

modal.addEventListener('click', (e) => {
    if (e.target === modal) {
        modal.style.display = 'none';
    }
});
