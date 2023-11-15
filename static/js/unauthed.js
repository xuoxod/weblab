document.addEventListener("DOMContentLoaded", function () {
  const signinLink = document.querySelector("#sign-in");
  const registerLink = document.querySelector("#register");

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
});
