const sections = [
  document.querySelector("#hero"),
  document.querySelector("#problem"),
  document.querySelector("#evidence"),
  document.querySelector("#work"),
  document.querySelector("#proof"),
  document.querySelector("#boundary"),
  document.querySelector("#contact")
].filter(Boolean);

const header = document.querySelector(".site-header");

function getHeaderOffset() {
  return header ? header.getBoundingClientRect().height : 0;
}

function getCurrentSectionIndex() {
  const offset = getHeaderOffset() + 24;

  let currentIndex = 0;

  for (let i = 0; i < sections.length; i += 1) {
    const rect = sections[i].getBoundingClientRect();

    if (rect.top <= offset) {
      currentIndex = i;
    }
  }

  return currentIndex;
}

function goToSection(index) {
  const section = sections[index];

  if (!section) {
    return;
  }

  const top =
    window.scrollY +
    section.getBoundingClientRect().top -
    getHeaderOffset();

  window.scrollTo({
    top,
    behavior: "smooth"
  });
}

window.addEventListener("keydown", (event) => {
  if (
    event.key !== "ArrowDown" &&
    event.key !== "PageDown"
  ) {
    return;
  }

  const target = event.target;

  if (
    target instanceof HTMLInputElement ||
    target instanceof HTMLTextAreaElement ||
    target instanceof HTMLSelectElement
  ) {
    return;
  }

  event.preventDefault();

  const currentIndex = getCurrentSectionIndex();

  goToSection(
    Math.min(
      currentIndex + 1,
      sections.length - 1
    )
  );
});