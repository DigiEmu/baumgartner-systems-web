(() => {
  const stages = Array.from(
    document.querySelectorAll(".arca-demo__stage")
  );

  if (!stages.length) return;


  const showStage = (number) => {
    const target = stages.find(
      (stage) =>
        stage.dataset.stage ===
        String(number)
    );

    if (!target) return;


    const mobile =
      window.matchMedia(
        "(max-width: 800px)"
      ).matches;


    /*
     * Switch visible stage first.
     * The previous stage disappears from layout,
     * so we wait for the browser to recalculate height
     * before adjusting scroll position.
     */

    stages.forEach((stage) => {
      stage.hidden =
        stage !== target;
    });


    requestAnimationFrame(() => {
      requestAnimationFrame(() => {

        const header =
          document.querySelector(
            ".site-header"
          );

        const headerHeight =
          header
            ? header.getBoundingClientRect().height
            : 0;


        const targetTop =
          target.getBoundingClientRect().top +
          window.scrollY -
          headerHeight;


        window.scrollTo({
          top: Math.max(0, targetTop),

          behavior:
            mobile
              ? "auto"
              : "smooth",
        });

      });
    });
  };


  document.addEventListener(
    "click",
    (event) => {

      /*
       * Continue through walkthrough
       */

      const next =
        event.target.closest(
          "[data-next]"
        );

      if (next) {
        showStage(
          next.dataset.next
        );

        return;
      }


      /*
       * Verification controls
       */

      const verification =
        document.querySelector(
          "[data-verification]"
        );

      const changeBox =
        document.querySelector(
          "[data-change-box]"
        );

      const evidenceState =
        document.querySelector(
          "[data-evidence-state]"
        );

      const simulate =
        event.target.closest(
          "[data-simulate-change]"
        );

      const restore =
        event.target.closest(
          "[data-restore]"
        );


      /*
       * Simulate evidence change
       */

      if (
        simulate &&
        verification
      ) {

        verification.classList.add(
          "is-changed"
        );


        const icon =
          verification.querySelector(
            ".arca-demo__verification-icon"
          );

        const title =
          verification.querySelector(
            "h3"
          );

        const copy =
          verification.querySelector(
            ".arca-demo__verification-copy"
          );


        if (icon) {
          icon.textContent = "!";
        }

        if (title) {
          title.textContent =
            "EVIDENCE CHANGED";
        }

        if (copy) {
          copy.textContent =
            "The current source no longer matches the evidence state used for the recorded result.";
        }


        if (evidenceState) {
          evidenceState.textContent =
            "Changed";
        }

        if (changeBox) {
          changeBox.hidden = false;
        }


        simulate.hidden = true;


        const restoreButton =
          document.querySelector(
            "[data-restore]"
          );

        if (restoreButton) {
          restoreButton.hidden =
            false;
        }


        return;
      }


      /*
       * Restore original evidence state
       */

      if (
        restore &&
        verification
      ) {

        verification.classList.remove(
          "is-changed"
        );


        const icon =
          verification.querySelector(
            ".arca-demo__verification-icon"
          );

        const title =
          verification.querySelector(
            "h3"
          );

        const copy =
          verification.querySelector(
            ".arca-demo__verification-copy"
          );


        if (icon) {
          icon.textContent = "✓";
        }

        if (title) {
          title.textContent =
            "VERIFIED";
        }

        if (copy) {
          copy.textContent =
            "The current evidence set matches the recorded state.";
        }


        if (evidenceState) {
          evidenceState.textContent =
            "Intact";
        }

        if (changeBox) {
          changeBox.hidden = true;
        }


        restore.hidden = true;


        const simulateButton =
          document.querySelector(
            "[data-simulate-change]"
          );

        if (simulateButton) {
          simulateButton.hidden =
            false;
        }


        return;
      }


      /*
       * Replay walkthrough
       */

      const restart =
        event.target.closest(
          "[data-restart]"
        );

      if (restart) {

        /*
         * Reset verification state before restart
         */

        const verification =
          document.querySelector(
            "[data-verification]"
          );

        if (verification) {

          verification.classList.remove(
            "is-changed"
          );


          const icon =
            verification.querySelector(
              ".arca-demo__verification-icon"
            );

          const title =
            verification.querySelector(
              "h3"
            );

          const copy =
            verification.querySelector(
              ".arca-demo__verification-copy"
            );


          if (icon) {
            icon.textContent = "✓";
          }

          if (title) {
            title.textContent =
              "VERIFIED";
          }

          if (copy) {
            copy.textContent =
              "The current evidence set matches the recorded state.";
          }
        }


        const evidenceState =
          document.querySelector(
            "[data-evidence-state]"
          );

        if (evidenceState) {
          evidenceState.textContent =
            "Intact";
        }


        const changeBox =
          document.querySelector(
            "[data-change-box]"
          );

        if (changeBox) {
          changeBox.hidden = true;
        }


        const restoreButton =
          document.querySelector(
            "[data-restore]"
          );

        if (restoreButton) {
          restoreButton.hidden =
            true;
        }


        const simulateButton =
          document.querySelector(
            "[data-simulate-change]"
          );

        if (simulateButton) {
          simulateButton.hidden =
            false;
        }


        showStage(1);
      }

    }
  );
})();