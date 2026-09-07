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

  window.addEventListener("scroll", updateHeader, {
    passive: true,
  });

  const toggle = header.querySelector(".mobile-nav-toggle");
  const nav = header.querySelector(".site-nav");
  const technologyDropdown = header.querySelector(".site-nav__dropdown");
  const technologyToggle = technologyDropdown?.querySelector(
    ".site-nav__dropdown-toggle"
  );

  const closeTechnologyDropdown = () => {
    if (!technologyDropdown || !technologyToggle) {
      return;
    }

    technologyDropdown.classList.remove("is-open");
    technologyToggle.setAttribute("aria-expanded", "false");
  };

  const closeNavigation = () => {
    header.classList.remove("is-nav-open");

    if (toggle) {
      toggle.setAttribute("aria-expanded", "false");
      toggle.setAttribute("aria-label", "Open navigation");
    }

    closeTechnologyDropdown();
  };

  if (toggle && nav) {
    toggle.addEventListener("click", (event) => {
      event.preventDefault();
      event.stopPropagation();

      const open = header.classList.toggle("is-nav-open");

      toggle.setAttribute("aria-expanded", String(open));
      toggle.setAttribute(
        "aria-label",
        open ? "Close navigation" : "Open navigation"
      );

      if (!open) {
        closeTechnologyDropdown();
      }
    });

    if (technologyDropdown && technologyToggle) {
      technologyToggle.addEventListener("click", (event) => {
        if (window.innerWidth > 900) {
          return;
        }

        event.preventDefault();
        event.stopPropagation();

        const open = technologyDropdown.classList.toggle("is-open");

        technologyToggle.setAttribute("aria-expanded", String(open));
      });
    }

    nav.addEventListener("click", (event) => {
      const link = event.target.closest("a");

      if (!link) {
        return;
      }

      if (link.closest(".site-nav__dropdown-toggle")) {
        return;
      }

      closeNavigation();
    });

    document.addEventListener("click", (event) => {
      if (
        header.classList.contains("is-nav-open") &&
        !header.contains(event.target)
      ) {
        closeNavigation();
      }
    });

    document.addEventListener("keydown", (event) => {
      if (event.key === "Escape") {
        closeNavigation();
      }
    });

    window.addEventListener("resize", () => {
      if (window.innerWidth > 900) {
        closeNavigation();
      }
    });
  }
}