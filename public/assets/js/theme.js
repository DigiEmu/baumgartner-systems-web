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
function syncBrowserThemeColor() {
  const meta = document.getElementById("theme-color-meta");

  if (!meta) {
    return;
  }

  const theme =
    document.documentElement.dataset.theme || "light";

  meta.setAttribute(
    "content",
    theme === "dark"
      ? "#151619"
      : "#b8912b"
  );
}

syncBrowserThemeColor();

const themeColorObserver = new MutationObserver(
  syncBrowserThemeColor
);

themeColorObserver.observe(
  document.documentElement,
  {
    attributes: true,
    attributeFilter: ["data-theme"]
  }
);
