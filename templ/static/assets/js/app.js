import "../css/style.css"


document.addEventListener('DOMContentLoaded', () => {
    const myDrawerTrigger = document.querySelector('.drawer #my-drawer')
    document.querySelectorAll('.drawer .drawer-side ul li').forEach((el) => {
        el.addEventListener('click', () => {
            if (myDrawerTrigger) {
                myDrawerTrigger.checked = false;
            }
        })
    })
})
