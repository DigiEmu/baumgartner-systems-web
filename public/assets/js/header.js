const header =
  document.querySelector(".site-header");

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


  /* mobile navigation */

  const toggle =
    header.querySelector(
      ".mobile-nav-toggle"
    );

  const nav =
    header.querySelector(
      ".site-nav"
    );

  if (toggle && nav) {

    const closeNavigation = () => {
      header.classList.remove(
        "is-nav-open"
      );

      toggle.setAttribute(
        "aria-expanded",
        "false"
      );

      toggle.setAttribute(
        "aria-label",
        "Open navigation"
      );
    };

    toggle.addEventListener(
      "click",
      () => {
        const open =
          header.classList.toggle(
            "is-nav-open"
          );

        toggle.setAttribute(
          "aria-expanded",
          String(open)
        );

        toggle.setAttribute(
          "aria-label",
          open
            ? "Close navigation"
            : "Open navigation"
        );
      }
    );

    nav.addEventListener(
      "click",
      (event) => {
        const link =
          event.target.closest("a");

        if (
          link &&
          !link.closest(
            ".site-nav__dropdown-menu"
          )
        ) {
          closeNavigation();
        }
      }
    );

    document.addEventListener(
      "keydown",
      (event) => {
        if (event.key === "Escape") {
          closeNavigation();
        }
      }
    );

    window.addEventListener(
      "resize",
      () => {
        if (window.innerWidth > 900) {
          closeNavigation();
        }
      }
    );
  }
}