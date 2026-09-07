const root = document.documentElement;

const systemTheme = window.matchMedia(
  "(prefers-color-scheme: dark)"
);

function applySystemTheme() {
  root.dataset.theme =
    systemTheme.matches
      ? "dark"
      : "light";
}

applySystemTheme();

systemTheme.addEventListener(
  "change",
  applySystemTheme
);