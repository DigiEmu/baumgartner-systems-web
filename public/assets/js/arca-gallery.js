(() => {
  const lightbox = document.querySelector(".arca-lightbox");
  if (!lightbox) return;

  const image = lightbox.querySelector(".arca-lightbox__image");
  const triggers = document.querySelectorAll("[data-zoom-src]");
  const closeTargets = lightbox.querySelectorAll("[data-lightbox-close]");

  const openLightbox = (src, alt) => {
    image.src = src;
    image.alt = alt || "";
    lightbox.hidden = false;
    document.body.style.overflow = "hidden";
  };

  const closeLightbox = () => {
    lightbox.hidden = true;
    image.src = "";
    image.alt = "";
    document.body.style.overflow = "";
  };

  triggers.forEach((trigger) => {
    trigger.addEventListener("click", () => {
      openLightbox(
        trigger.dataset.zoomSrc,
        trigger.dataset.zoomAlt
      );
    });
  });

  closeTargets.forEach((target) => {
    target.addEventListener("click", closeLightbox);
  });

  document.addEventListener("keydown", (event) => {
    if (event.key === "Escape" && !lightbox.hidden) {
      closeLightbox();
    }
  });
})();