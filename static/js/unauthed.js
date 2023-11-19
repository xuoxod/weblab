document.addEventListener("DOMContentLoaded", function () {
  const signinLink = document.querySelector("#mobile-sign-in");
  const registerLink = document.querySelector("#mobile-register");
  const aboutLink = document.querySelector("#mobile-about-link");

  if (aboutLink && registerLink && aboutLink) {
    addClickHandler(aboutLink, about);

    addClickHandler(registerLink, register);

    addClickHandler(signinLink, signin);
  }
});

function siteMenuEffect() {
  var x = document.getElementById("site-menu");
  if (x.className.indexOf("w3-show") == -1) {
    x.className += " w3-show";
  } else {
    x.className = x.className.replace(" w3-show", "");
  }
}
