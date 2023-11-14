const log = (msg) => console.log(msg);

const stringify = (obj) => JSON.stringify(obj);

const parser = (strObj) => JSON.parse(strObj);

const addHandler = (theElement, whichEvent, method) => {
  if (null != theElement && null != whichEvent && typeof method == "function") {
    theElement.addEventListener(whichEvent, method);
  }
};

const addClickHandler = (theElement, handler) => {
  if (null != theElement && typeof handler == "function") {
    addHandler(theElement, "click", handler);
  }
};

const addKeyupHandler = (theElement, handler) => {
  if (null != theElement && typeof handler == "function") {
    addHandler(theElement, "keyup", handler);
  }
};

const addKeydownHandler = (theElement, handler) => {
  if (null != theElement && typeof handler == "function") {
    addHandler(theElement, "keydown", handler);
  }
};

const addOnFocusHandler = (theElement, handler) => {
  if (null != theElement && typeof handler == "function") {
    addHandler(theElement, "focus", handler);
  }
};

const addOnChangeHandler = (theElement, handler) => {
  if (null != theElement && typeof handler == "function") {
    addHandler(theElement, "change", handler);
  }
};

const addOffFocusHandler = (theElement, handler) => {
  if (null != theElement && typeof handler == "function") {
    addHandler(theElement, "focusout", handler);
  }
};

const appendChild = (parent, child) => {
  if (null != parent && null != child) {
    parent.appendChild(child);
  }
};

const appendBeforeLastChild = (parent, child) => {
  if (null != parent && null != child) {
    const lastChildIndex = parent.children.length - 1;
    const lastChild = parent.children[lastChildIndex];
    parent.insertBefore(child, lastChild);
  }
};

const append = (parent, child) => {
  parent.append(child);
};

const removeChildren = (parent) => {
  parent.querySelectorAll("*").forEach((dialog) => {
    dialog.remove();
  });
};

const countChildren = (parent) => {
  if (null != parent) {
    return parent.children.length;
  }
  return null;
};

const addAttribute = (theElement, whichAttribute, attributeValue) => {
  if (null != theElement) {
    theElement.setAttribute(whichAttribute, attributeValue);
  }
};

const setAttribute = (theElement, whichAttribute, attributeValue) => {
  if (null != theElement) {
    theElement.setAttribute(whichAttribute, attributeValue);
  }
};

const getAttribute = (theElement, whichAttribute) => {
  if (null != theElement && null != whichAttribute) {
    return theElement.getAttribute(`${whichAttribute}`) || null;
  }
  return "Element is null";
};

const removeAttribute = (theElement, whichAttribute) => {
  if (null != theElement) {
    if (theElement.hasAttribute(whichAttribute)) {
      theElement.removeAttribute(whichAttribute);
    }
  }
};

const getElement = (nameIdClass) => {
  let element = null;
  if (null != (element = document.querySelector(`${nameIdClass}`))) {
    return element;
  }
  if (null != (element = document.querySelector(`#${nameIdClass}`))) {
    return element;
  }
  if (null != (element = document.querySelector(`.${nameIdClass}`))) {
    return element;
  }
  return null;
};

const cap = (arg) => {
  let word_split = null,
    line = "",
    word = arg.toString();
  if (null !== word && undefined !== word) {
    if (
      word.trim().toLowerCase() === "id" ||
      word.trim().toLowerCase() === "ssn" ||
      word.trim().toLowerCase() === "sku" ||
      word.trim().toLowerCase() === "vm" ||
      word.trim().toLowerCase() === "mac" ||
      word.trim().toLowerCase() === "imei" ||
      word.trim().toLowerCase() === "os" ||
      word.trim().toLowerCase() === "atm" ||
      word.trim().toLowerCase() === "pa" ||
      word.trim().toLowerCase() === "rjw"
    ) {
      return word.toUpperCase();
    } else if (word.match(/[-]/)) {
      if (null !== (word_split = word.split(["-"])).length > 0) {
        for (let i = 0; i < word_split.length; i++) {
          if (i < word_split.length - 1) {
            line +=
              word_split[i].substring(0, 1).toUpperCase() +
              word_split[i].substring(1) +
              "-";
          } else {
            line +=
              word_split[i].substring(0, 1).toUpperCase() +
              word_split[i].substring(1);
          }
        }
        return line;
      }
    } else if (word.match(/[ ]/)) {
      if (null !== (word_split = word.split([" "])).length > 0) {
        for (let i = 0; i < word_split.length; i++) {
          if (i < word_split.length - 1) {
            line +=
              word_split[i].substring(0, 1).toUpperCase() +
              word_split[i].substring(1) +
              " ";
          } else {
            line +=
              word_split[i].substring(0, 1).toUpperCase() +
              word_split[i].substring(1);
          }
        }
        return line;
      }
    } else {
      return word.substring(0, 1).toUpperCase() + word.substring(1);
    }
  }
};

const newElement = (type) => {
  if (null != type && typeof type == "string") {
    return document.createElement(type);
  }
  return null;
};

// notify display a custom alert to user
// type string: success, error or warning
// msg string: message to user
const notify = (type, msg, time = 3) => {
  notie.alert({
    type: type,
    text: msg,
    time: time,
  });
};

// modal display custom modal
// title string: modal's title
// text string: the message to user
// icon built-in: success, warning, error, info or question
// btnText string: button's text
// showStatus: true or false
const modal = (
  title,
  text,
  icon = "info",
  btnText = "Okay",
  btnClose = true,
  showStatus = true
) => {
  Swal.fire({
    title: title,
    text: text,
    icon: icon,
    confirmButtonText: btnText,
    showCloseButton: btnClose,
    showLoading: showStatus,
  });
};

const signin = async () => {
  const form = await Swal.fire({
    title: "Log In",
    icon: "info",
    showConfirmButton: true,
    confirmButtonText: "Confirm",
    showCancelButton: true,
    cancelButtonText: "Cancel",
    allowEscapeKey: true,
    allowEnterKey: true,

    html: `
    <form id="signin-form">
      <div class="input-group">
          <label class="input-group-text">
              <strong>
                  <i class="bi bi-envelope-at-fill fs-3"></i>
              </strong>
          </label>

          <input id="email" type="email" name="email" placeholder="Enter email address" autocomplete="false" class="form-control">
      </div>

      <div class="input-group mt-3">
          <label class="input-group-text">
              <strong>
                  <i class="bi bi-lock-fill fs-3"></i>
              </strong>
          </label>

          <input id="password" type="password" name="password" placeholder="Enter password" autocomplete="true" class="form-control">
      </div>
    </form>
  `,
    focusConfirm: true,
    preConfirm: () => {
      return [
        document.querySelector("#email").value,
        document.querySelector("#password").value,
      ];
    },
  })
    .then((results) => {
      const { isConfirmed } = results;
      if (isConfirmed) {
        log(`Confirmed`);
        const signinForm = document.querySelector("#signin-form");
        const email = document.querySelector("#email").value;
        const password = document.querySelector("#password").value;
        const token = document.querySelector("#csrf").value;

        if (email && password && token) {
          log(`Token:\t${token}\n`);
          const formData = new FormData(signinForm);
          formData.append("csrf_token", token);
          try {
            fetch("/login", {
              method: "post",
              body: formData,
            })
              .then((response) => response.json())
              .then((data) => {
                log(data);
              });
          } catch (err) {
            log(err);
          }

          log(`Submitted Signin Form\n`);
        } else {
          Swal.closeModal();
        }
      } else {
        log(`Dismissed`);
        Swal.closeModal();
      }
    })
    .catch((err) => {
      log(err);
    });

  // if (formValues) {
  //   const confirmButton = Swal.getConfirmButton();
  //   const cancelButton = Swal.getCancelButton();

  //   Swal.fire({
  //     title: "Confirm",
  //     text: "This action cannot be undone.",
  //     icon: "success",
  //     showCancelButton: true,
  //     cancelButtonText: "Cancel",
  //     confirmButtonText: "Yes, submit",
  //   }).then((results) => {
  //     const { isConfirmed, isDenied, isDismissed } = results;
  //     if (isConfirmed) {
  //       log(`Confirmed`);
  //     } else if (isDenied) {
  //       log(`Denied`);
  //     } else {
  //       log(`Dismissed`);
  //     }
  //   });
  // } else {
  //   Swal.fire({
  //     title: "Done",
  //     icon: "success",
  //     animation: true,
  //   });
  // }
};
