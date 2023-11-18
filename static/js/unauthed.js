document.addEventListener("DOMContentLoaded", function () {
  if (
    document.querySelector("#sign-in") &&
    document.querySelector("#register")
  ) {
    const signinLink = document.querySelector("#sign-in");
    const registerLink = document.querySelector("#register");
    const aboutLink = document.querySelector("#about-link");

    if (signinLink) {
      signinLink.addEventListener("click", () => {
        signin();
      });
    }

    if (registerLink) {
      registerLink.addEventListener("click", () => {
        register();
      });
    }

    if (aboutLink) {
      aboutLink.addEventListener("click", (e) => {
        about();
      });
    }
  }

  // Intersection observer - menu
  if ("IntersectionObserver" in window) {
    let observer = new IntersectionObserver(
      (entries, observer) => {
        entries.forEach((entry) => {
          /* Here's where we deal with every intersection */
          entry.target.src = entry.target.dataset.src;
          observer.unobserve(entry.target);
        });
      },
      { rootMargin: "0px 0px -200px 0px" }
    );
  }
});
