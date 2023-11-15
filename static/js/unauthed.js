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
});
