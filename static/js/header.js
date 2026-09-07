const header = document.querySelector(".site-header");

if (header) {
  let scrolled = false;

  const updateHeader = () => {
    const y = window.scrollY;

    if (!scrolled && y > 48) {
      scrolled = true;
      header.classList.add("is-scrolled");
    }

    if (scrolled && y < 16) {
      scrolled = false;
      header.classList.remove("is-scrolled");
    }
  };

  updateHeader();

  window.addEventListener(
    "scroll",
    updateHeader,
    { passive: true }
  );
}