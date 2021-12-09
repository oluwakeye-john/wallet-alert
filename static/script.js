(function () {
  const Ethereum = window.ethereum;

  let address = "";
  let currency = "";
  let email = "";

  const onMetamaskClick = async () => {
    if (Ethereum && Ethereum.isMetaMask) {
      const { error, account } = await getMetamaskAddress();

      if (error) {
        console.log(error);
        updateToast(error);
      } else {
        document.querySelector("#address-input").value = account;
        document.querySelector("#currency-input").value = "ETH";
      }
    } else {
      updateToast("Metamask is not available");
    }
  };

  const getMetamaskAddress = async () => {
    const request = {
      method: "eth_requestAccounts",
      params: [],
    };

    const defaultErrorMessage = "Error occurred";

    try {
      if (Ethereum.isConnected() && Ethereum.selectedAddress) {
        const account = Ethereum.selectedAddress;
        return { error: "", account };
      }

      const accounts = await Ethereum.request(request);

      if (accounts.length >= 1) {
        const account = accounts[0];
        return { error: "", account };
      } else {
        return defaultErrorMessage, "";
      }
    } catch (err) {
      return { error: err?.message || defaultErrorMessage, account: null };
    }
  };

  const createSubscription = async (data) => {
    const { email, address, currency } = data;

    const query = `
      mutation {
          createSubscription(input: { address: "${address}", email: "${email}", currency_code: ${currency} }) {
              address
              is_subscribed
          }
      }
    `;

    return fetch("/query", {
      body: JSON.stringify({ query }),
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((res) => res.json())
      .then((res) => {
        if (res.errors) {
          console.log(res.errors, res.errors[0].message);
          throw res.errors[0].message;
        }
        return;
      })
      .catch((err) => {
        return err || "An error occurred";
      });
  };

  const updateToast = (text, success = false) => {
    const toastEl = document.querySelector("#toast");

    if (success) {
      toastEl.style.color = "var(--primary)";
    } else {
      toastEl.style.color = "var(--error)";
    }
    anime({
      targets: "#toast",
      translateY: [30, 0],
      opacity: [0, 1],
      easing: "spring",
      duration: 100,
    });
    toastEl.innerText = text;

    setTimeout(() => {
      anime({
        targets: "#toast",
        translateY: [0, 30],
        opacity: [1, 0],
        easing: "spring",
        duration: 100,
        complete: () => {
          toastEl.innerText = "";
        },
      });
    }, 3000);
  };

  document
    .querySelector("#metamask-login-btn")
    .addEventListener("click", onMetamaskClick);

  window.addEventListener("load", () => {
    anime({
      targets: ".title, .subtitle, .form-1",
      translateY: [80, 0],
      opacity: [0, 1],
      easing: "spring",
      duration: 200,
      delay: anime.stagger(100),
    });

    anime({
      targets: ".label, .owner",
      translateY: [-20, 0],
      opacity: [0, 1],
      easing: "spring",
      duration: 200,
      delay: anime.stagger(100),
    });
    document.querySelector(".form-1").style.pointerEvents = "all";
  });

  document.querySelector(".form-1").addEventListener("submit", (e) => {
    e.preventDefault();

    address = document.querySelector("#address-input").value;
    currency = document.querySelector("#currency-input").value;

    document.querySelector(".form-1").style.pointerEvents = "none";
    document.querySelector(".form-2").style.pointerEvents = "all";

    anime({
      targets: ".form-1",
      translateY: [0, 100],
      opacity: [1, 0],
      easing: "spring",
      duration: 200,
    });

    anime({
      targets: ".form-2",
      translateY: [100, 0],
      opacity: [0, 1],
      easing: "spring",
      duration: 200,
    });
  });

  document.querySelector(".form-2").addEventListener("submit", async (e) => {
    e.preventDefault();
    email = document.querySelector("#email-input").value;
    try {
      setLoading(true);
      const error = await createSubscription({ address, email, currency });
      if (error) {
        updateToast(error);
      } else {
        updateToast("Success", true);
      }
    } catch (err) {
      updateToast(err);
    } finally {
      setLoading(false);
    }
  });

  document.querySelector(".back-btn").addEventListener("click", () => {
    document.querySelector(".form-1").style.pointerEvents = "all";
    document.querySelector(".form-2").style.pointerEvents = "none";

    anime({
      targets: ".form-1",
      translateY: [100, 0],
      opacity: [0, 1],
      easing: "spring",
      duration: 200,
    });

    anime({
      targets: ".form-2",
      translateY: [0, 100],
      opacity: [1, 0],
      easing: "spring",
      duration: 200,
    });
  });

  const setLoading = (value) => {
    document.querySelector("#form2-btn").disabled = value;
  };
})();
