const root = document.documentElement;
const systemTheme = window.matchMedia("(prefers-color-scheme: dark)");

function storedTheme() {
  try {
    return localStorage.getItem("baumgartner-theme");
  } catch (e) {
    return null;
  }
}

function applyTheme() {
  const stored = storedTheme();
  root.dataset.theme =
    stored || (systemTheme.matches ? "dark" : "light");
}

applyTheme();

systemTheme.addEventListener("change", () => {
  if (!storedTheme()) {
    applyTheme();
  }
});

function syncBrowserThemeColor() {
  const meta = document.getElementById("theme-color-meta");
  if (!meta) return;

  const theme = root.dataset.theme || "light";
  meta.setAttribute(
    "content",
    theme === "dark" ? "#151619" : "#b8912b"
  );
}

syncBrowserThemeColor();

const themeColorObserver = new MutationObserver(syncBrowserThemeColor);
themeColorObserver.observe(root, {
  attributes: true,
  attributeFilter: ["data-theme"],
});
