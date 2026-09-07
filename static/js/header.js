const header =
  document.querySelector(".site-header");

if (header) {
  let scrolled = false;


  /* =========================================================
     HEADER SCROLL STATE
     ========================================================= */

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


  /* =========================================================
     ACTIVE SECTION
     ========================================================= */

  const sectionLinks =
    [
      ...header.querySelectorAll(
        '.site-nav a[href^="#"]'
      )
    ];

  const sections =
    sectionLinks
      .map((link) => {
        const selector =
          link.getAttribute("href");

        if (
          !selector ||
          selector === "#"
        ) {
          return null;
        }

        const section =
          document.querySelector(
            selector
          );

        if (!section) {
          return null;
        }

        return {
          link,
          section
        };
      })
      .filter(Boolean);


  const updateActiveSection = () => {
    if (!sections.length) {
      return;
    }

    const viewportMarker =
      window.scrollY + 160;

    let active =
      sections[0];

    for (const item of sections) {
      if (
        item.section.offsetTop <=
        viewportMarker
      ) {
        active = item;
      }
    }


    /*
     * Force final contact section active
     * when page bottom is reached.
     */

    const atBottom =
      window.innerHeight +
      window.scrollY >=
      document.documentElement.scrollHeight -
      8;

    if (atBottom) {
      active =
        sections[
          sections.length - 1
        ];
    }


    sectionLinks.forEach(
      (link) => {
        link.classList.remove(
          "is-active"
        );
      }
    );

    if (active?.link) {
      active.link.classList.add(
        "is-active"
      );
    }
  };

  updateActiveSection();

  window.addEventListener(
    "scroll",
    updateActiveSection,
    { passive: true }
  );


  /* =========================================================
     MOBILE NAVIGATION
     ========================================================= */

  const toggle =
    header.querySelector(
      ".mobile-nav-toggle"
    );

  const nav =
    header.querySelector(
      ".site-nav"
    );

  const technologyDropdown =
    header.querySelector(
      ".site-nav__dropdown"
    );

  const technologyToggle =
    technologyDropdown?.querySelector(
      ".site-nav__dropdown-toggle"
    );


  const closeTechnologyDropdown = () => {
    if (
      !technologyDropdown ||
      !technologyToggle
    ) {
      return;
    }

    technologyDropdown.classList.remove(
      "is-open"
    );

    technologyToggle.setAttribute(
      "aria-expanded",
      "false"
    );
  };


  const closeNavigation = () => {
    header.classList.remove(
      "is-nav-open"
    );

    if (toggle) {
      toggle.setAttribute(
        "aria-expanded",
        "false"
      );

      toggle.setAttribute(
        "aria-label",
        "Open navigation"
      );
    }

    closeTechnologyDropdown();
  };


  if (toggle && nav) {

    /*
     * Main hamburger
     */

    toggle.addEventListener(
      "click",
      (event) => {
        event.stopPropagation();

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

        if (!open) {
          closeTechnologyDropdown();
        }
      }
    );


    /*
     * Technology submenu
     */

    if (
      technologyDropdown &&
      technologyToggle
    ) {
      technologyToggle.addEventListener(
        "click",
        (event) => {

          /*
           * Desktop dropdown behaviour
           * remains CSS / existing behaviour.
           */

          if (
            window.innerWidth > 900
          ) {
            return;
          }

          event.preventDefault();
          event.stopPropagation();

          const open =
            technologyDropdown.classList.toggle(
              "is-open"
            );

          technologyToggle.setAttribute(
            "aria-expanded",
            String(open)
          );
        }
      );
    }


    /*
     * Close after selecting a link
     */

    nav.addEventListener(
      "click",
      (event) => {
        const link =
          event.target.closest("a");

        if (!link) {
          return;
        }

        closeNavigation();

        /*
         * Update marker immediately
         * after anchor navigation.
         */

        window.requestAnimationFrame(
          updateActiveSection
        );
      }
    );


    /*
     * Click outside
     */

    document.addEventListener(
      "click",
      (event) => {
        if (
          header.classList.contains(
            "is-nav-open"
          ) &&
          !header.contains(
            event.target
          )
        ) {
          closeNavigation();
        }
      }
    );


    /*
     * Escape
     */

    document.addEventListener(
      "keydown",
      (event) => {
        if (
          event.key === "Escape"
        ) {
          closeNavigation();

          if (toggle) {
            toggle.focus();
          }
        }
      }
    );


    /*
     * Reset when returning
     * to desktop layout
     */

    window.addEventListener(
      "resize",
      () => {
        if (
          window.innerWidth > 900
        ) {
          closeNavigation();
        }

        updateActiveSection();
      }
    );
  }
}