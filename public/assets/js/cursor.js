(() => {
  const finePointer =
    window.matchMedia(
      "(hover: hover) and (pointer: fine)"
    );

  if (!finePointer.matches) {
    return;
  }

  const cursor =
    document.createElement("div");

  cursor.className =
    "pearl-cursor";

  cursor.setAttribute(
    "aria-hidden",
    "true"
  );

  document.body.appendChild(cursor);


  let x = -100;
  let y = -100;

  let targetX = -100;
  let targetY = -100;


  const render = () => {
    x +=
      (targetX - x) * 0.22;

    y +=
      (targetY - y) * 0.22;

    cursor.style.transform =
      `translate3d(${x}px, ${y}px, 0)`;

    requestAnimationFrame(render);
  };


  document.addEventListener(
    "pointermove",
    (event) => {
      targetX = event.clientX;
      targetY = event.clientY;

      cursor.classList.add(
        "is-visible"
      );
    },
    { passive: true }
  );


  document.addEventListener(
    "pointerover",
    (event) => {
      const interactive =
        event.target.closest(
          "a, button, input, textarea, select, [role='button']"
        );

      cursor.classList.toggle(
        "is-interactive",
        Boolean(interactive)
      );
    }
  );


  document.addEventListener(
    "pointerleave",
    () => {
      cursor.classList.remove(
        "is-visible"
      );
    }
  );


  document.addEventListener(
    "pointerdown",
    () => {
      cursor.classList.add(
        "is-pressed"
      );
    }
  );


  document.addEventListener(
    "pointerup",
    () => {
      cursor.classList.remove(
        "is-pressed"
      );
    }
  );


  render();
})();